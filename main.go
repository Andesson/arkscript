package main

import (
	"fmt"

	"github.com/Andesson/arkscript/config"
	"github.com/Andesson/arkscript/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Errorf("Config initialization error: %v", err)
		return
	}
	defer config.GetSQLite().Close()
	router.Initialize()
}
