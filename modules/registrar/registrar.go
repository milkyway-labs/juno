package registrar

import (
	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/logging"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/modules/pruning"
	"github.com/forbole/juno/v5/modules/telemetry"
	"github.com/forbole/juno/v5/node"
	"github.com/forbole/juno/v5/types"
	"github.com/forbole/juno/v5/types/config"
)

// Context represents the context of the modules registrar
type Context struct {
	ConfigPath     string
	JunoConfig     config.Config
	EncodingConfig types.EncodingConfig
	Database       database.Database
	Proxy          node.Node
	Logger         logging.Logger
}

// NewContext allows to build a new Context instance
func NewContext(
	configPath string,
	parsingConfig config.Config,
	encodingConfig types.EncodingConfig,
	database database.Database,
	proxy node.Node,
	logger logging.Logger,
) Context {
	return Context{
		ConfigPath:     configPath,
		JunoConfig:     parsingConfig,
		EncodingConfig: encodingConfig,
		Database:       database,
		Proxy:          proxy,
		Logger:         logger,
	}
}

// Registrar represents a modules registrar. This allows to build a list of modules that can later be used by
// specifying their names inside the TOML configuration file.
type Registrar interface {
	BuildModules(context Context) modules.Modules
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &EmptyRegistrar{}
)

// EmptyRegistrar represents a Registrar which does not register any custom module
type EmptyRegistrar struct{}

// BuildModules implements Registrar
func (*EmptyRegistrar) BuildModules(_ Context) modules.Modules {
	return nil
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &DefaultRegistrar{}
)

// DefaultRegistrar represents a registrar that allows to handle the default Juno modules
type DefaultRegistrar struct {
}

// NewDefaultRegistrar builds a new DefaultRegistrar
func NewDefaultRegistrar() *DefaultRegistrar {
	return &DefaultRegistrar{}
}

// BuildModules implements Registrar
func (r *DefaultRegistrar) BuildModules(ctx Context) modules.Modules {
	return modules.Modules{
		pruning.NewModule(ctx.JunoConfig, ctx.Database, ctx.Logger),
		telemetry.NewModule(ctx.JunoConfig),
	}
}

// ------------------------------------------------------------------------------------------------------------------

// GetModules returns the list of module implementations based on the given module names.
// For each module name that is specified but not found, a warning log is printed.
func GetModules(mods modules.Modules, names []string, logger logging.Logger) []modules.Module {
	var modulesImpls []modules.Module
	for _, name := range names {
		module, found := mods.FindByName(name)
		if found {
			modulesImpls = append(modulesImpls, module)
		} else {
			logger.Error("Module is required but not registered. Be sure to register it using registrar.RegisterModule", "module", name)
		}
	}
	return modulesImpls
}
