package component

import "GOIF/types"

var RegisteredComponents []*types.Component

type (
	IComponent interface {
		run(ctx types.ContextRequest) types.HttpTransport
	}

	Builder types.Component
)


func CreateComponent(schema string) *Builder {
	return &Builder{
		Schema: schema,
	}
}

func (builder *Builder) Processor(fn func(ctx types.ContextRequest) types.HttpTransport) *Builder {
	builder.Function = fn
	return builder
}

func (builder *Builder) End() *types.Component {
	component := (*types.Component)(builder)
	RegisteredComponents = append(RegisteredComponents, component)
	return component
}
