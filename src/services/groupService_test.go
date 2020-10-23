package services

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/grandeep/org-svc/cmd/org-svc/engine"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"testing"
)

var (
	groupTestRepo repositories.RepoI
	groupTestService GroupServiceInterface
	groupCtx context.Context

)

func initTest() {
	groupCtx = context.Background()
	configPath := "/Users/demo/Code/go/org-svc/resources/config/config.yaml"
	e := engine.NewEngine(configPath)
	groupTestRepo = repositories.NewRepoI(e.DB)
	groupTestService = NewGroupService(groupTestRepo, e.Config)
}

func TestMain(m *testing.M) {
	initTest()
	m.Run()

}

func TestStart(t *testing.T) {
	var err error
	//err = testGroupAddSvc()
	//if err != nil {
	//	t.Error(err)
	//}

	//err = testGroupTreeQuerySvc()
	//if err != nil {
	//	t.Error(err)
	//}

	err = testGroupUpdateSvc()
	if err != nil {
		t.Error(err)
	}
}

// testGroupAddSvc 测试添加组
func testGroupAddSvc() error {

	quotas := []*pb_user_v1.Quota {
		{
			IsShare:              1,
			ResourcesGroupId:     "10,20",
			Gpu:                  123,
			Cpu:                  234,
			Memory:               345,
		},
		{
			IsShare:              2,
			ResourcesGroupId:     "10,20",
			Gpu:                  456,
			Cpu:                  567,
			Memory:               678,
		},
	}

	data := &pb_user_v1.GroupAddRequest{
		Name:                 "张三",
		ParentId:             12,
		DiskQuotaSize:        100,
		Quotas:               quotas,
	}

	resp, err := groupTestService.GroupAddSvc(groupCtx, data)
	if err != nil {
		return err

	}

	if resp.Code != 0 {
		return errors.New("请求失败")
	}
	return nil
}

// testGroupTreeQuerySvc 测试组树查询
func testGroupTreeQuerySvc() error {
	
	data := &pb_user_v1.GroupID{
		Id:                   0,
	}

	resp, err := groupTestService.GroupTreeQuerySvc(groupCtx, data)
	if err != nil {
		return err

	}

	fmt.Println("Group Tree: ", resp)

	return nil
}

// testGroupUpdateSvc 测试组更新
func testGroupUpdateSvc() error {

	//data := &pb_user_v1.GroupUpdateRequest{
	//	Id:                   28,
	//	Name:                 "",
	//	ParentId:             0,
	//	UseParentId:          false,
	//	Description:          "28描述",
	//}
	//
	//resp, err := groupTestService.GroupUpdateSvc(groupCtx, data)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println("update resp: ", resp)

	data2 := &pb_user_v1.GroupUpdateRequest{
		Id:                   14,
		Name:                 "14的新名字",
		ParentId:             12,
		UseParentId:          true,
		Description:          "14的新描述",
	}

	_, err := groupTestService.GroupUpdateSvc(groupCtx, data2)
	if err != nil {
		return err
	}

	return nil
}






