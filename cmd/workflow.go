package main

import (
	"github.com/jssyzhoulei/workflow/cmd/engine"
	"github.com/jssyzhoulei/workflow/src/apis"
	"github.com/jssyzhoulei/workflow/src/repositories"
	"github.com/jssyzhoulei/workflow/src/routers"
	"github.com/jssyzhoulei/workflow/src/services"
)

func main() {

	db := engine.InitDB()
	repo := repositories.NewRepo(db)
	svc := services.NewWorkSvc(repo)
	api := apis.NewApis(svc)
	gin := routers.Gin()
	routers.Routers(gin, api)
	_ = gin.Run(":88")
}
