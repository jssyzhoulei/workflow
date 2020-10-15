package main

import (
	"gitee.com/grandeep/org-svc/src/routers"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
)

func main()  {
	routers.Routers(routers.Gin())
	log.Logger().Info("Gin Running...")
	_ = routers.Gin().Run(":88")
}
