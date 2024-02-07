package gprc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/status"

	"github.com/forbole/juno/v5/jsonrpc2"
)

var (
	_ grpc.ClientConnInterface = &Connection{}
)

// Connection represents a custom gRPC connection implementation that relies on a RPC connection instead of
// a gRPC connection
type Connection struct {
	jsonrpcClient *jsonrpc2.Client
	gprcCdc       encoding.Codec
}

// NewConnection a new Connection instance
func NewConnection(rpcAddress string, cdc codec.Codec) (*Connection, error) {
	jsonRPCClient, err := jsonrpc2.NewClient(rpcAddress, &http.Client{Timeout: time.Minute})
	if err != nil {
		return nil, err
	}

	protoCodec, ok := cdc.(*codec.ProtoCodec)
	if !ok {
		return nil, fmt.Errorf("invalid codec type: %T", cdc)
	}

	return &Connection{
		jsonrpcClient: jsonRPCClient,
		gprcCdc:       protoCodec.GRPCCodec(),
	}, nil
}

// MustCreateConnection returns a new Connection instance, or panics if any error arises
func MustCreateConnection(rpcAddress string, cdc codec.Codec) *Connection {
	conn, err := NewConnection(rpcAddress, cdc)
	if err != nil {
		panic(err)
	}
	return conn
}

// Invoke implements the grpc.ClientConnInterface interface
func (c *Connection) Invoke(ctx context.Context, method string, args, reply any, _ ...grpc.CallOption) error {
	req, err := c.gprcCdc.Marshal(args)
	if err != nil {
		return err
	}

	res, err := c.RunABCIQuery(ctx, method, req)
	if err != nil {
		return fmt.Errorf("abci query: %w", err)
	}

	if !res.Response.IsOK() {
		return status.Error(codes.Unknown, res.Response.Log) // TODO: better status code?
	}

	err = c.gprcCdc.Unmarshal(res.Response.Value, reply)
	if err != nil {
		return err
	}

	return nil
}

// RunABCIQuery runs a new query through the ABCI protocol
func (c *Connection) RunABCIQuery(ctx context.Context, path string, data []byte) (*ABCIQueryResult, error) {
	var res ABCIQueryResult
	err := c.jsonrpcClient.Call(ctx, "abci_query", ABCIQueryRequest{
		Path:   path,
		Data:   data,
		Height: 0,
		Prove:  false,
	}, &res)

	if err != nil {
		return nil, fmt.Errorf("call abci_query: %w", err)
	}

	return &res, nil
}

// NewStream implements the grpc.ClientConnInterface interface
func (c *Connection) NewStream(_ context.Context, _ *grpc.StreamDesc, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not implemented")
}
