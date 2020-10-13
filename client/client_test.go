package client

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	"testing"
	"time"
)

func TestOrgServiceClient_GetUserService(t *testing.T) {
	o := NewOrgServiceClient([]string{"127.0.0.1:2379"}, 3, time.Second)
	o.GetUserService().AddUserSvc(context.Background(), models.User2{
		UserId: 1,
		UserName: "f",
	})
}
