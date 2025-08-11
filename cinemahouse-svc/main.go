package main

import (
	_ "cinemahouse-svc/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cinemahouse-svc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
