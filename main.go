package main

import (
	"sync"

	"./radata"
)

//go:generate goversioninfo -icon=ra-go.ico
//PackageInfo :- https://github.com/josephspurrier/goversioninfo

//global vars
var (
	Version    string = "1.10.0"
	uiPortFlag string = "50001"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	radata.HandleFiles()
	go startUIService(uiPortFlag)
	wg.Wait()
}
