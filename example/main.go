package main

import "github.com/RayfenWindspear/flipper"

func main() {
	err := flipper.DoEverything()
	if err != nil {
		panic(err)
	}
}
