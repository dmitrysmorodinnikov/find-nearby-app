package main

import "find-nearby-backend/cmd"

func main() {
	cli := cmd.NewRootCmd()
	cli.Execute()
}
