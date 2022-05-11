package main

import "github.com/CheckPointSW/infinity-next-terraform-cli/cmd"

var version = "development"

func main() {
	cmd.Execute(version)
}
