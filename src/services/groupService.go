package services

import (
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"golang.org/x/net/context"
)

// GroupServiceInterface 组服务接口
type GroupServiceI interface {
	GroupAdd(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error)
}

// GroupService 组服务,实现了 GroupServiceInterface
type GroupService struct {
	groupRepo repositories.GroupRepoI
}

// NewGroupService GroupService 构造函数
func NewGroupService(repos repositories.RepoI) GroupServiceI {
	return &GroupService{
		groupRepo: repos.GetGroupRepo(),
	}
}

// GroupAdd 添加组
func (g GroupService) GroupAdd(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
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

	err = g.groupRepo.GroupAdd(newGroup, tx)
	if err != nil {
		return &pb_user_v1.GroupResponse{Code: 1}, err
	}

	group, err := g.groupRepo.GroupQueryByName(data.Name, tx)
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

	err = g.groupRepo.QuotaAdd(result, tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return &pb_user_v1.GroupResponse{Code: 0}, nil
}
