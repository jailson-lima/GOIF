package component

import (
	"GOIF/types"
	"testing"
)

var component = CreateComponent("test")

func TestCreateComponent(t *testing.T) {
	checkComponent(component, t)
}

func TestComponent_Processor(t *testing.T) {
	component = component.Processor(func(context ContextRequest) types.HttpTransport {
		return context.HttpTransport
	})

	checkComponent(component, t)
}

func TestComponent_End(t *testing.T) {
	component.End()

	found := false
	for _, components := range RegisteredComponents {
		if components == component {
			found = true
			break
		}
	}

	if found {
		t.Log("Success, component found in RegisteredComponents.")
	} else {
		t.Error("Failed, component not found in RegisteredComponents.")
	}
}

func checkComponent(component *Component, t *testing.T) {
	if component == nil {
		t.Error("Failed, expected return Component{}.")
	} else {
		t.Logf("Success, expected Component{}, got %v.", component)
	}
}