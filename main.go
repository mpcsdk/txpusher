package main

import (
	_ "txpusher/internal/packed"

	_ "txpusher/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"txpusher/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
