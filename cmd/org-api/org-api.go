package main

import "gitee.com/grandeep/org-svc/src/routers"

func main()  {
	routers.Routers(routers.Gin())
	routers.Gin().Run(":88")
}