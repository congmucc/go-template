package main

import "gotemplate/cmd"

func main() {
	defer cmd.Clear()
	cmd.Start()
}
