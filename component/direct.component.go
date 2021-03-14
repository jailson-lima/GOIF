package component

import (
	"GOIF/internal"
	"GOIF/types"
	"fmt"
)

func init() {
	CreateComponent("direct").Processor(directProcessor).End()
}

func directProcessor(context ContextRequest) types.HttpTransport {
	for _, functions := range internal.ContextFunctions {
		if functions.Id == fmt.Sprintf("direct:%s", context.Component.Uri) {
			for _, processor := range functions.Processors {
				context.HttpTransport = processor(context.HttpTransport)
			}
			break
		}
	}

	return context.HttpTransport
}