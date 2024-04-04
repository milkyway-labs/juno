package apis

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/forbole/juno/v5/modules/apis/endpoints"
	"github.com/forbole/juno/v5/modules/registrar"
)

// Context contains all the useful data that might be used when registering an API handler
type Context struct {
	registrar.Context
	GRPCConnection *grpc.ClientConn
}

func NewContext(ctx registrar.Context, grpcConnection *grpc.ClientConn) Context {
	return Context{
		Context:        ctx,
		GRPCConnection: grpcConnection,
	}
}

// Registrar represents a function that allows registering API endpoints
type Registrar func(ctx Context, router *gin.Engine) error

// CombinedRegistrar returns a new Registrar combining the given API registrars together
func CombinedRegistrar(registrars ...Registrar) Registrar {
	return func(ctx Context, router *gin.Engine) error {
		for _, register := range registrars {
			err := register(ctx, router)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// DefaultRegistrar returns the default API registrar
func DefaultRegistrar(_ Context, router *gin.Engine) error {
	endpoints.RegisterRoutesList(router)
	return nil
}
