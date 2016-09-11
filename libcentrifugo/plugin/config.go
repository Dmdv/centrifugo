package plugin

import (
	"github.com/centrifugal/centrifugo/libcentrifugo/engine"
	"github.com/centrifugal/centrifugo/libcentrifugo/server"
)

type ConfigSetter interface {
	SetDefault(key string, value interface{})
	BindEnv(key string)
	BindFlag(key string, flagName string)
	StringFlag(name, shorthand string, value string, usage string)
	BoolFlag(name, shorthand string, value bool, usage string)
	IntFlag(name, shorthand string, value int, usage string)
}

type ConfigGetter interface {
	Get(string) interface{}
	GetString(string) string
	GetBool(string) bool
	GetInt(string) int
	IsSet(string) bool
}

type EngineFactory func(server.Node, ConfigGetter) engine.Engine

type Configurator func(ConfigSetter) error

var EngineFactories map[string]EngineFactory

var Configurators map[string]Configurator

func RegisterEngine(name string, fn EngineFactory) {
	EngineFactories[name] = fn
}

func RegisterConfigurator(name string, fn Configurator) {
	Configurators[name] = fn
}

func init() {
	EngineFactories = map[string]EngineFactory{}
	Configurators = map[string]Configurator{}
}