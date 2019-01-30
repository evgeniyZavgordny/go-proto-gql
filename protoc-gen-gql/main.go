package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/danielvladco/go-proto-gql/plugin"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

func main() {
	gen := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		gen.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, gen.Request); err != nil {
		gen.Error(err, "parsing input proto")
	}

	p := &plugin{NewPlugin()}

	gen.CommandLineParameters(gen.Request.GetParameter())
	gen.WrapTypes()
	gen.SetPackageNames()
	gen.BuildTypeNameMap()
	gen.GeneratePlugin(p)

	fileLen := len(gen.Response.GetFile())
	for i := 0; i < fileLen; i++ {
		gen.Response.File[i].Name = proto.String(strings.Replace(gen.Response.File[i].GetName(), ".pb.go", ".pb.graphqls", -1))
		gen.Response.File[i].Content = proto.String(p.GetSchemaByIndex(i).String())
	}

	data, err = proto.Marshal(gen.Response)
	if err != nil {
		gen.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		gen.Error(err, "failed to write output proto")
	}
}

type plugin struct {
	*Plugin
}

func (p *plugin) GenerateImports(file *generator.FileDescriptor) {}
func (p *plugin) Name() string                                   { return "gql" }
func (p *plugin) Init(g *generator.Generator)                    { p.Generator = g }

func (p *plugin) Generate(file *generator.FileDescriptor) {
	for _, fileName := range p.Request.FileToGenerate {
		if fileName == file.GetName() {
			p.InitFile(file)
			p.generate(file)
		}
	}
}

func (p *plugin) generate(file *generator.FileDescriptor) {
	p.P("# Code generated by protoc-gen-gogql. DO NOT EDIT")
	p.P("# source: ", file.GetName(), "\n")
	p.PrintComments(2)
	// TODO: add comments from proto to gql for documentation purposes to all elements. Currently are supported only a few

	if len(p.Oneofs()) > 0 {
		p.P("directive @oneof(name: String) on FIELD_DEFINITION\n")
	}

	// scalars
	for _, entry := range TypeMap2List(p.Scalars()) {
		scalar := entry.Value
		if !scalar.BuiltIn {
			p.P("scalar ", p.GqlModelNames()[scalar], "\n")
		}
	}

	// maps
	for _, entry := range TypeMap2List(p.Maps()) {
		m := entry.Value
		key, val := m.GetMapFields()

		messagesIn := make(map[string]*Type)
		if m.UsedAsInput {
			messagesIn = p.Inputs()
		} else {
			messagesIn = p.Types()
		}
		p.P("# map with key: '", p.GraphQLType(key, nil), "' and value: '", p.GraphQLType(val, messagesIn), "'")
		p.P("scalar ", p.GqlModelNames()[m], "\n")
	}

	// render gql inputs
	for _, entry := range TypeMap2List(p.Inputs()) {
		protoTypeName, m := entry.Key, entry.Value
		if p.IsEmpty(m) || p.IsAny(protoTypeName) {
			continue
		}
		p.P("input ", p.GqlModelNames()[m], " {")
		p.In()
		for _, f := range m.GetField() {
			opts := GetGqlFieldOptions(f)
			if opts != nil {
				if opts.Dirs != nil {
					*opts.Dirs = strings.Trim(*opts.Dirs, " ")
				}
			}

			if typ, ok := p.Inputs()[f.GetTypeName()]; ok {
				if p.IsEmpty(typ) {
					continue
				}
			}
			oneof := ""
			if f.OneofIndex != nil {
				oneof = fmt.Sprintf("@oneof(name: %q)", p.GqlModelNames()[p.GetOneof(OneofRef{Parent: m, Index: f.GetOneofIndex()})])
			}
			p.P(ToLowerFirst(generator.CamelCase(f.GetName())), ": ", p.GraphQLType(f, p.Inputs()), " ", oneof, " ", opts.GetDirs())
		}
		p.Out()
		p.P("}\n")
	}

	// render gql types
	for _, entry := range TypeMap2List(p.Types()) {
		protoTypeName, m := entry.Key, entry.Value
		if p.IsEmpty(m) || p.IsAny(protoTypeName) {
			continue
		}
		p.P("type ", p.GqlModelNames()[m], " {")
		p.In()
		oneofs := make(map[OneofRef]struct{})
		for _, f := range m.GetField() {
			if f.OneofIndex != nil {
				oneofRef := OneofRef{Parent: m, Index: *f.OneofIndex}
				if _, ok := oneofs[oneofRef]; !ok {
					p.P(ToLowerFirst(generator.CamelCase(m.OneofDecl[*f.OneofIndex].GetName())), ": ", p.GqlModelNames()[p.GetOneof(oneofRef)])
				}
				oneofs[oneofRef] = struct{}{}
				continue
			}
			opts := GetGqlFieldOptions(f)
			if opts != nil {
				if opts.Params != nil {
					*opts.Params = "(" + strings.Trim(*opts.Params, " ") + ")"
				}
				if opts.Dirs != nil {
					*opts.Dirs = strings.Trim(*opts.Dirs, " ")
				}
			}

			if typ, ok := p.Types()[f.GetTypeName()]; ok {
				if p.IsEmpty(typ) {
					continue
				}
			}

			p.P(ToLowerFirst(generator.CamelCase(f.GetName())), opts.GetParams(), ": ", p.GraphQLType(f, p.Types()), " ", opts.GetDirs())
		}
		p.Out()
		p.P("}\n")
	}

	for _, entry := range TypeMap2List(p.Oneofs()) {
		oneof := entry.Value
		p.P("union ", p.GqlModelNames()[oneof], " = ", strings.Join(func() (ss []string) {
			for typeName := range oneof.OneofTypes {
				ss = append(ss, p.GqlModelNames()[p.Types()[typeName]])
			}
			return
		}(), " | "), "\n")
	}

	// render enums
	for _, e := range p.Enums() {
		p.P("enum ", p.GqlModelNames()[e], " {")
		p.In()
		for _, v := range e.GetValue() {
			p.P(v.GetName())
		}
		p.Out()
		p.P("}\n")
	}

	if len(p.Mutations()) > 0 {
		p.P("type Mutation {")
		p.In()
		p.renderMethod(p.Mutations())
		p.Out()
		p.P("}\n")
	}
	if len(p.Queries()) > 0 {
		p.P("type Query {")
		p.In()
		p.renderMethod(p.Queries())
		p.Out()
		p.P("}\n")
	}
	if len(p.Subscriptions()) > 0 {
		p.P("type Subscription {")
		p.In()
		p.renderMethod(p.Subscriptions())
		p.Out()
		p.P("}\n")
	}
}

func (p *plugin) renderMethod(methods []*Method) {
	for _, m := range methods {
		in, out := "", "Boolean!"

		// some scalars can be used as inputs such as 'Any'
		if scalar, ok := p.Scalars()[m.InputType]; ok {
			if !p.IsEmpty(scalar) {
				in = "(in: " + p.GqlModelNames()[scalar] + ")"
			}
		} else if input, ok := p.Inputs()[m.InputType]; ok {
			if !p.IsEmpty(input) {
				in = "(in: " + p.GqlModelNames()[input] + ")"
			}
		}
		if scalar, ok := p.Scalars()[m.OutputType]; ok {
			if !p.IsEmpty(scalar) {
				in = "(in: " + p.GqlModelNames()[scalar] + ")"
			}
		} else if typ, ok := p.Types()[m.OutputType]; ok {
			if !p.IsEmpty(typ) {
				out = p.GqlModelNames()[typ]
			}
		}

		p.PrintComments(6, m.ServiceIndex, 2, m.Index)
		if m.Phony {
			p.P("# See the Subscription with the same name")
		}
		p.P(m.Name, in, ": ", out)
	}
}
