package component

import (
	"github.com/JPereirax/GOIF/types"
	"log"
)

type LogComponent struct{}

func init() {
	CreateComponent("log").Processor(LogComponent{}.run).End()
}

func (LogComponent) run(ctx types.ContextRequest) (*types.HttpTransport, error) {
	log.Print(ctx.Component.Uri)
	return ctx.HttpTransport, nil
}