// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/options.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/danielvladco/go-proto-gql/pb"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServiceGQLServer struct{ Service ServiceServer }

func (s *ServiceGQLServer) ServiceMutate1(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Mutate1(ctx, in)
}
func (s *ServiceGQLServer) ServiceMutate2(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Mutate2(ctx, in)
}
func (s *ServiceGQLServer) ServiceQuery1(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Query1(ctx, in)
}

type QueryGQLServer struct{ Service QueryServer }

func (s *QueryGQLServer) QueryQuery1(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Query1(ctx, in)
}
func (s *QueryGQLServer) QueryQuery2(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Query2(ctx, in)
}
func (s *QueryGQLServer) QueryMutate1(ctx context.Context, in *Data) (*Data, error) {
	return s.Service.Mutate1(ctx, in)
}
