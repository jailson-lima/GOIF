package component

import "github.com/JPereirax/GOIF/types"

var RegisteredComponents []*types.Component

type (
	IComponent interface {
		run(ctx types.ContextRequest) (*types.HttpTransport, error)
	}

	Builder types.Component
)


func CreateComponent(schema string) *Builder {
	return &Builder{
		Schema: schema,
	}
}

func (builder *Builder) Processor(fn func(ctx types.ContextRequest) (*types.HttpTransport, error)) *Builder {
	builder.Function = fn
	return builder
}

func (builder *Builder) End() *types.Component {
	component := (*types.Component)(builder)
	RegisteredComponents = append(RegisteredComponents, component)
	return component
}
