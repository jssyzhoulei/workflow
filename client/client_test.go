package client

import (
	"context"
	"fmt"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"testing"
	"time"
)

func TestOrgServiceClient_GetUserService(t *testing.T) {
	o := NewOrgServiceClient([]string{"127.0.0.1:2379"}, 1, time.Second)
	fmt.Println(o.GetGroupService().GetAllGroup(context.Background(), &pb_user_v1.GroupID{

	}))
}
