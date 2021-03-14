package component

import "GOIF/types"

var RegisteredComponents []*Component

type Component struct {
	Schema    string
	Uri       string
	Function func(context ContextRequest) types.HttpTransport
}

type ContextRequest struct {
	Component Component

	HttpTransport types.HttpTransport
}

func CreateComponent(schema string) *Component {
	return &Component{
		Schema: schema,
	}
}

func (component *Component) Processor(function func(context ContextRequest) types.HttpTransport) *Component {
	component.Function = function
	return component
}

func (component *Component) End() *Component {
	RegisteredComponents = append(RegisteredComponents, component)
	return component
}
