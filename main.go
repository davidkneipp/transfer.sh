package main

import "github.com/davidkneipp/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
