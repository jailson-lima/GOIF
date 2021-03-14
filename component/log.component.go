package component

import (
	"GOIF/types"
	"log"
)

func init() {
	CreateComponent("log").Processor(logProcessor).End()
}

func logProcessor(context ContextRequest) types.HttpTransport {
	log.Print(context.Component.Uri)
	return context.HttpTransport
}