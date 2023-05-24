package main

import (
	"fmt"
	"github.com/mebiusashan/gcms/internal/config"
)

func main() {
	cfx, err := config.Read("E:\\git\\gcms\\configs\\gcms.toml")
	fmt.Println(cfx.Server.Port)
	fmt.Println(err)
}
