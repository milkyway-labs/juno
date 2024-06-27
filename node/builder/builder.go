package builder

import (
	"fmt"

	"github.com/forbole/juno/v5/node"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	"github.com/forbole/juno/v5/node/remote"
	"github.com/forbole/juno/v5/types"
)

type Context struct {
	EncodingConfig       types.EncodingConfig
	AccountAddressParser types.AccountAddressParser
}

func NewContext(encodingConfig types.EncodingConfig, accountAddressParser types.AccountAddressParser) Context {
	return Context{
		EncodingConfig:       encodingConfig,
		AccountAddressParser: accountAddressParser,
	}
}

func BuildNode(cfg nodeconfig.Config, ctx Context) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details), ctx.AccountAddressParser, ctx.EncodingConfig.Codec, ctx.EncodingConfig.GRPCodec)
	// case nodeconfig.TypeLocal:
	// 	return local.NewNode(cfg.Details.(*local.Details), txConfig, ctx.AccountAddressParser, cdc)
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
