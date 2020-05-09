package main

import (
	"fmt"

	"./config"
)

func main() {
	fetchConfig()
}

func fetchConfig() {
	cnf := new(config.Config)
	c1 := cnf.GetConfig()
	fmt.Println(c1)
	fmt.Println("Current environment Proposed as", config.CurrentEnv)
}
