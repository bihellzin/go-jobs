package main

import (
	"github.com/bihellzin/go-jobs/config"
	"github.com/bihellzin/go-jobs/router"
)

var (
	logger *config.Logger
)

func main()  {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
