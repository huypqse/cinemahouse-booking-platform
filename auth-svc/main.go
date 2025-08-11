package main

import (
	_ "auth-svc/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"auth-svc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
