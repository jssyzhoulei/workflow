package main

import (
	"flag"
	"github.com/jssyzhoulei/workflow/logger"
	"github.com/jssyzhoulei/workflow/src/routers"
)

func main() {
	flag.Parse()
	routers.Routers(routers.Gin())
	log.Logger.Info("Gin Running...")
	_ = routers.Gin().Run(":88")
}
