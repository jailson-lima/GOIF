package component

import (
	"GOIF/types"
	"log"
)

type LogComponent struct{}

func init() {
	CreateComponent("log").Processor(LogComponent{}.run).End()
}

func (LogComponent) run(ctx types.ContextRequest) types.HttpTransport {
	log.Print(ctx.Component.Uri)
	return ctx.HttpTransport
}