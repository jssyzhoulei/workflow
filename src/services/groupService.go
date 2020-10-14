package services

import (
	"fmt"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"golang.org/x/net/context"
)

// GroupServiceInterface 组服务接口
type GroupServiceInterface interface {
	GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error)
	GroupQueryByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryByConditionRequest) (*pb_user_v1.GroupQueryByConditionResponse, error)
}

// GroupService 组服务,实现了 GroupServiceInterface
type GroupService struct {
	groupRepo repositories.GroupRepoInterface
}

// NewGroupService GroupService 构造函数
func NewGroupService(repos repositories.RepoI) GroupServiceInterface {
	return &GroupService{
		groupRepo: repos.GetGroupRepo(),
	}
}

// GroupAddSvc 添加组
func (g *GroupService) GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
	var err error
	tx := g.groupRepo.GetTx()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newGroup := &models.Group{
		Name:     data.Name,
		ParentID: int(data.ParentId),
	}

	err = g.groupRepo.GroupAddRepo(newGroup, tx)
	if err != nil {
		return &pb_user_v1.GroupResponse{Code: 1}, err
	}

	group, err := g.groupRepo.GroupQueryByNameRepo(data.Name, tx)
	if err != nil {
		return nil, err
	}
	quotaTypeMap := map[string]models.ResourceType{
		"cpu":    models.RESOURCE_CPU,
		"gpu":    models.RESOURCE_GPU,
		"memory": models.RESOURCE_MEMORY,
		"disk":   models.RESOURCE_DISK,
	}

	l := len(data.Quotas)
	var result []*models.Quota
	for i := 0; i < l; i++ {
		q := data.Quotas[i]

		valMap := map[string]int64{
			"cpu":    q.Cpu,
			"gpu":    q.Gpu,
			"memory": q.Memory,
			"disk":   data.DiskQuotaSize,
		}

		for kind, val := range valMap {
			tmp := &models.Quota{
				IsShare:    int(q.IsShare),
				ResourceID: q.ResourcesGroupId,
				Type:       quotaTypeMap[kind],
				GroupID:    group.ID,
				Total:      int(val),
				Used:       0,
			}
			result = append(result, tmp)
		}
	}

	err = g.groupRepo.QuotaAddRepo(result, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &pb_user_v1.GroupResponse{Code: 0}, nil
}

// GroupQueryByConditionSvc ...
func (g *GroupService) GroupQueryByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryByConditionRequest) (*pb_user_v1.GroupQueryByConditionResponse, error) {

	var condition string
	var value []interface{}

	// TODO: 等待实现数据处理逻辑,生成条件

	params := map[string]interface{}{
		"id":         data.Id,
		"name":       data.Name,
		"parent_id":  data.ParentId,
		"created_at": data.CreateTime,
	}

	//for column, val := range params {
	//
	//	switch val.(type) {
	//	case []int64:
	//		fmt.Println("[]int64")
	//	case []string:
	//		fmt.Println("[]string")
	//	case string:
	//		fmt.Println("string")
	//	}
	//}

	fmt.Println(params)


	resp, err := g.groupRepo.GroupQueryByConditionRepo(condition, nil, value...)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	// TODO: 等实现查询结果组装

	return nil, nil
}
