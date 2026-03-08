package app

import (
	"context"

	"golang.design/x/clipboard"
)

// global state

var Ctx context.Context

func Startup(ctx context.Context) {
	clipboard.Init()
	Ctx = ctx
}
