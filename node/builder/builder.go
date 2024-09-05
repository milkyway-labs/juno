package builder

import (
	"fmt"

	"github.com/forbole/juno/v5/node"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	"github.com/forbole/juno/v5/node/remote"
	"github.com/forbole/juno/v5/types"
)

type Context struct {
	EncodingConfig   types.EncodingConfig
	TxHashCalculator types.TxHashCalculator
}

func NewContext(
	encodingConfig types.EncodingConfig,
	txHashCalculator types.TxHashCalculator,
) Context {
	return Context{
		EncodingConfig:   encodingConfig,
		TxHashCalculator: txHashCalculator,
	}
}

func BuildNode(cfg nodeconfig.Config, ctx Context) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(
			cfg.Details.(*remote.Details),
			ctx.TxHashCalculator,
			ctx.EncodingConfig.GRPCodec,
		)
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
