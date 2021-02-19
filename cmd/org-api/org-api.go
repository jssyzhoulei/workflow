package main

import (
	"flag"
	"gitee.com/grandeep/org-svc/src/routers"
	"gitee.com/grandeep/org-svc/logger"
)

func main()  {
	flag.Parse()
	routers.Routers(routers.Gin())
	log.Logger.Info("Gin Running...")
	_ = routers.Gin().Run(":88")
}
