package main

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/api/cli/flags"
	"github.com/boltdbwebeditor/boltdbwebeditor/api/webServer"
)

func main() {
	flags := flags.ParseFlags()
	webServer.Start(flags)
}
