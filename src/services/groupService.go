package services

import (
	"context"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
)

// GroupServiceInterface 组服务接口
type GroupServiceInterface interface {
	GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error)
	GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error)
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
		"cpu":    models.ResourceCpu,
		"gpu":    models.ResourceGpu,
		"memory": models.ResourceMemory,
		"disk":   models.ResourceDisk,
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

// GroupQueryByConditionSvc 根据条件查询组信息和其配额信息
func (g *GroupService) GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error) {

	condition := &models.GroupQueryByCondition{
		ID:       data.Id,
		Name:     data.Name,
		ParentID: data.ParentId,
	}

	querySlice, err := g.groupRepo.GroupQueryWithQuotaByConditionRepo(condition, nil)
	if err != nil {
		return nil, err
	}

	var groupData = make(map[int64]*pb_user_v1.GroupWithQuota)
	var groupQuotaTemp = make(map[int64]map[int]map[string]*pb_user_v1.Quota)

	l := len(querySlice)
	for i := 0; i < l; i++ {
		r := querySlice[i]
		if _, ok := groupData[r.ID]; !ok {
			groupData[r.ID] = new(pb_user_v1.GroupWithQuota)
			groupData[r.ID].ParentId = r.ParentID
			groupData[r.ID].Id = r.ID
			groupData[r.ID].Name = r.Name
			groupData[r.ID].Quotas = make([]*pb_user_v1.Quota, 0)
		}
		if models.ResourceType(r.Type) == models.ResourceDisk {
			groupData[r.ID].DiskQuotaSize = int64(r.Total)
		}

		// 对零散的配额数据进行分组
		// 判断 group
		if _, ok := groupQuotaTemp[r.ID]; !ok {
			groupQuotaTemp[r.ID] = make(map[int]map[string]*pb_user_v1.Quota)
		}
		// 判断 is_share 配额类型,使用 is_share + resources_id 进行分组
		if _, ok := groupQuotaTemp[r.ID][r.IsShare]; !ok {
			groupQuotaTemp[r.ID][r.IsShare] = make(map[string]*pb_user_v1.Quota)
		}
		if _, ok := groupQuotaTemp[r.ID][r.IsShare][r.ResourceID]; !ok {
			// 创建 quota 数据,并设置 is_share 和 resources_id 字段信息
			groupQuotaTemp[r.ID][r.IsShare][r.ResourceID] = &pb_user_v1.Quota{
				IsShare:              int64(r.IsShare),
				ResourcesGroupId:     r.ResourceID,
			}
		}

		// 获取 quota 设置对应的值
		quotaData := groupQuotaTemp[r.ID][r.IsShare][r.ResourceID]
		switch models.ResourceType(r.Type) {
		case models.ResourceCpu:
			quotaData.Cpu = int64(r.Total)
		case models.ResourceGpu:
			quotaData.Gpu = int64(r.Total)
		case models.ResourceMemory:
			quotaData.Memory = int64(r.Total)
		}
	}

	// 组装 quota 数据到 group 数据中
	for groupID, v1 := range groupQuotaTemp {
		quotas := make([]*pb_user_v1.Quota, 0)
		for _, v2 := range v1 {
			for _, q := range v2 {
				quotas = append(quotas, q)
			}
		}
		groupData[groupID].Quotas = quotas
	}

	// 结果 map -> slice
	var result = make([]*pb_user_v1.GroupWithQuota, 0)
	for _, g := range groupData {
		result = append(result, g)
	}

	return &pb_user_v1.GroupQueryWithQuotaByConditionResponse{
		Groups: result,
	}, nil
}
