package main

import (
	"flag"
	"gitee.com/grandeep/org-svc/src/routers"
)

func main()  {
	flag.Parse()
	routers.Routers(routers.Gin())
	_ = routers.Gin().Run(":88")
}
