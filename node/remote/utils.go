package remote

import (
	"context"
	"crypto/tls"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding"

	grpctypes "github.com/forbole/juno/v5/cosmos-sdk/types/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/forbole/juno/v5/gprc"
)

var (
	HTTPProtocols = regexp.MustCompile("https?://")
)

// GetHeightRequestContext adds the height to the context for querying the state at a given height
func GetHeightRequestContext(context context.Context, height int64) context.Context {
	return metadata.AppendToOutgoingContext(
		context,
		grpctypes.GRPCBlockHeightHeader,
		strconv.FormatInt(height, 10),
	)
}

// MustCreateGrpcConnection creates a new gRPC connection using the provided configuration and panics on error
func MustCreateGrpcConnection(cfg *Details, cdc encoding.Codec) grpc.ClientConnInterface {
	grpConnection, err := CreateGrpcConnection(cfg, cdc)
	if err != nil {
		panic(err)
	}
	return grpConnection
}

// CreateGrpcConnection creates a new gRPC client connection from the given configuration
func CreateGrpcConnection(cfg *Details, cdc encoding.Codec) (grpc.ClientConnInterface, error) {
	// If the gRPC config is not specified, we can create a gRPC-over-RPC connection
	if cfg.GRPC == nil {
		grpcConnection, err := gprc.NewConnection(cfg.RPC.Address, cdc)
		if err != nil {
			return nil, err
		}

		return grpcConnection, nil
	}

	var grpcOpts []grpc.DialOption
	if !strings.Contains("cfg.GRPC.Address", "https://") {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})))
	}

	address := HTTPProtocols.ReplaceAllString(cfg.GRPC.Address, "")
	return grpc.Dial(address, grpcOpts...)
}
