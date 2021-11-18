package main

import (
	"fmt"
	"github.com/whaoa/dot-todo-api/package/config"
)

func main() {
	fmt.Println(config.Get())
}
