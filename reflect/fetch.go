// Package grpcreflection provides gRPC reflection client.
// Currently, gRPC reflection depends on Protocol Buffers, so we split this package from grpc package.
package main

import (
	"context"
	"errors"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	gr "github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
)

const reflectionServiceName = "grpc.reflection.v1alpha.ServerReflection"

var ErrTLSHandshakeFailed = errors.New("TLS handshake failed")

// Client defines gRPC reflection client.
type Client interface {
	// ListPackages lists file descriptors from the gRPC reflection server.
	// ListPackages returns these errors:
	//   - ErrTLSHandshakeFailed: TLS misconfig.
	ListPackages() ([]*descriptor.FileDescriptorProto, error)
	// Reset clears internal states of Client.
	Reset()
}

type client struct {
	client *gr.Client
}

// NewClient returns an instance of gRPC reflection client for gRPC protocol.
func NewClient(conn *grpc.ClientConn) Client {
	return &client{
		client: gr.NewClient(context.Background(), grpc_reflection_v1alpha.NewServerReflectionClient(conn)),
	}
}

func (c *client) ListPackages() ([]*descriptor.FileDescriptorProto, error) {
	ssvcs, err := c.client.ListServices()
	if err != nil {
		msg := status.Convert(err).Message()
		// Check whether the error message contains TLS related error.
		// If the server didn't enable TLS, the error message contains the first string.
		// If Evans didn't enable TLS against to the TLS enabled server, the error message contains
		// the second string.
		if strings.Contains(msg, "tls: first record does not look like a TLS handshake") ||
			strings.Contains(msg, "latest connection error: <nil>") {
			return nil, ErrTLSHandshakeFailed
		}
		return nil, errors.Wrap(err, "failed to list services from reflecton enabled gRPC server")
	}

	fds := make([]*descriptor.FileDescriptorProto, 0, len(ssvcs))
	for _, s := range ssvcs {
		if s == reflectionServiceName {
			continue
		}
		svc, err := c.client.ResolveService(s)
		if err != nil {
			return nil, err
		}
		fds = append(fds, svc.GetFile().AsFileDescriptorProto())
	}

	return fds, nil
}

func (c *client) Reset() {
	c.client.Reset()
}
