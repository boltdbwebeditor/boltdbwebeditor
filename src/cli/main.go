package main

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/cli/flags"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/webServer"
)

func main() {
	flags := flags.ParseFlags()
	webServer.Start(flags)
}
