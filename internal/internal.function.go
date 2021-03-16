package internal

import "GOIF/types"

var ContextFunctions []*types.InternalFunction

type Function struct {
	id         string
	processors []func(transport *types.HttpTransport) *types.HttpTransport
}

func From(id string) *Function {
	return &Function{
		id: id,
	}
}

func (function *Function) Processor(processor func(transport *types.HttpTransport) *types.HttpTransport) *Function {
	function.processors = append(function.processors, processor)
	return function
}

func (function *Function) End() {
	ContextFunctions = append(ContextFunctions, &types.InternalFunction{
		Id: function.id,
		Processors: function.processors,
	})
}
