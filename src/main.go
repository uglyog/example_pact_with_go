package main

import (
	"pact"
	"github.com/paulbellamy/mango"
)

func main() {
	stack := new(mango.Stack)
	stack.Address = ":3000"
	stack.Middleware(mango.ShowErrors(""))
	stack.Run(pact.Producer)
}
