package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/grandeep/device-plugin/client"
	"gitee.com/grandeep/device-plugin/src/services"
	"gitee.com/grandeep/org-svc/src/models"
	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
	"gitee.com/grandeep/org-svc/src/repositories"
	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"gitee.com/grandeep/org-svc/utils/src/pkg/md5"
	"strconv"
	"strings"
	"time"
)

// GroupServiceInterface 组服务接口
type GroupServiceInterface interface {
	GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error)
	GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error)
	GroupUpdateSvc(ctx context.Context, data *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error)
	QuotaUpdateSvc(ctx context.Context, data *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error)
	GroupTreeQuerySvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error)
	GroupDeleteSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error)
	QueryGroupAndSubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error)
}

// GroupService 组服务,实现了 GroupServiceInterface
type GroupService struct {
	groupRepo         repositories.GroupRepoInterface
	userRepo          repositories.UserRepoInterface
	kubernetesService services.KubernetesServiceI
	cfg               *config.Config
}

// NewGroupService GroupService 构造函数
func NewGroupService(repos repositories.RepoI, cfg *config.Config) GroupServiceInterface {
	etcdHosts, _ := cfg.GetString("etcdHost")

	deviceClient := client.NewDeviceClient(strings.Split(etcdHosts, ";"), 2, time.Second*5)
	return &GroupService{
		groupRepo:         repos.GetGroupRepo(),
		userRepo:          repos.GetUserRepo(),
		kubernetesService: deviceClient.GetKubernetesService(),
		cfg:               cfg,
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
	md5Str := md5.EncodeMD5(data.Name)

	newGroup := &models.Group{
		Name:      data.Name,
		ParentID:  int(data.ParentId),
		NameSpace: md5Str,
	}

	err = g.groupRepo.GroupAddRepo(newGroup, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	group, err := g.groupRepo.GroupQueryByNameRepo(data.Name, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 相同配额类型,资源组校验不允许重复
	var _share int64
	var _nonShare int64
	var _shareResourcesID string
	var _nonShareResourcesID string
	for i := 0; i < len(data.Quotas); i++ {
		q := data.Quotas[i]

		if q.IsShare == 1 {
			if _share == q.IsShare && _shareResourcesID == q.ResourcesGroupId {
				return &pb_user_v1.GroupResponse{Code: 1}, errors.New("共享配额类型,重复划分相同资源组")
			}
			_share = q.IsShare
			_shareResourcesID = q.ResourcesGroupId
		} else if q.IsShare == 2 {
			if _nonShare == q.IsShare && _nonShareResourcesID == q.ResourcesGroupId {
				return &pb_user_v1.GroupResponse{Code: 1}, errors.New("独享配额类型,重复划分相同资源组")
			}
			_nonShare = q.IsShare
			_nonShareResourcesID = q.ResourcesGroupId
		}
	}

	quotaTypeMap := map[string]models.ResourceType{
		"cpu":    models.ResourceCpu,
		"gpu":    models.ResourceGpu,
		"memory": models.ResourceMemory,
		"disk":   models.ResourceDisk,
	}

	l := len(data.Quotas)
	var result = make([]*models.Quota, 0)
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
		tx.Rollback()
		return nil, err
	}

	_, err = g.kubernetesService.CreateNamespaceSvc(ctx, md5Str)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &pb_user_v1.GroupResponse{Code: 0}, nil
}

// GroupQueryWithQuotaByConditionSvc 根据条件查询组信息和其配额信息
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
			levelPath := strings.Split(r.LevelPath, "-")
			var topParent string
			if levelPath[1] == "" {
				topParent = "0"
			} else {
				topParent = levelPath[1]
			}
			topParentID, err := strconv.ParseInt(topParent, 10, 64)
			if err != nil {
				return nil, errors.New("转换顶级组ID失败:" + err.Error())
			}
			groupData[r.ID].TopParentId = topParentID
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
				IsShare:          int64(r.IsShare),
				ResourcesGroupId: r.ResourceID,
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

// GroupUpdateSvc 组信息更新
func (g *GroupService) GroupUpdateSvc(ctx context.Context, data *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error) {

	d := &models.GroupUpdateRequest{
		ID:   data.Id,
		Name: data.Name,
	}

	if data.UseParentId {
		d.ParentID = &data.ParentId
	}

	err := g.groupRepo.GroupUpdateRepo(d, nil)
	if err != nil {
		return &pb_user_v1.GroupResponse{Code: 1}, err
	}

	return &pb_user_v1.GroupResponse{Code: 0}, nil
}

// QuotaUpdateSvc 配额信息更新
func (g *GroupService) QuotaUpdateSvc(ctx context.Context, data *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error) {

	var err error
	d := &models.QuotaUpdateRequest{
		GroupID:     data.GroupId,
		IsShare:     data.IsShare,
		ResourcesID: data.ResourcesId,
		QuotaType:   data.QuotaType,
		Total:       data.Total,
		Used:        data.Used,
	}

	err = g.groupRepo.QuotaUpdateRepo(d, nil)
	if err != nil {
		return &pb_user_v1.GroupResponse{Code: 1}, nil
	}
	return &pb_user_v1.GroupResponse{Code: 0}, nil
}

// GroupDeleteSvc 组删除(软删除)
func (g *GroupService) GroupDeleteSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error) {
	if data.Id == 0 {
		return nil, errors.New("组 id 不允许为空")
	}

	// 验证组下是否存在用户
	users, err := g.userRepo.GetUserListRepo(models.User{GroupID: int(data.Id)}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(users) > 1 {
		return nil, fmt.Errorf("无法删除组,组内存在用户: %d 个", len(users))
	}

	// 验证是否存在下级组
	groups, err := g.groupRepo.GroupListWithChangedLevelPathRepo(data.Id, nil)
	if err != nil {
		return nil, fmt.Errorf("无法删除组,查询是否包含下级组时错误: %s", err.Error())
	}

	if len(groups) > 1 {
		return nil, fmt.Errorf("无法删除组,包含下级组: %d 个", len(groups))
	}

	// 删除组
	err = g.groupRepo.GroupDeleteRepo(data.Id, nil)
	if err != nil {
		return nil, err
	}

	return &pb_user_v1.GroupResponse{Code: 0}, nil
}

// GroupTreeQuerySvc 组树查询
func (g *GroupService) GroupTreeQuerySvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error) {

	groupList, err := g.groupRepo.GroupListWithChangedLevelPathRepo(data.Id, nil)
	if err != nil {
		return nil, err
	}

	tree := generateGroupTree(groupList, int(data.Id))
	if tree == nil {
		return nil, errors.New("生成失败,结果为空")
	}

	jsonByte, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}

	res := &pb_user_v1.GroupTreeResponse{
		TreeJson: jsonByte,
	}

	return res, nil
}

// insertToParentChildren generateGroupTree 中查找父节点的递归方法
func insertToParentChildren(currentNode *models.GroupTreeNode, node *models.GroupTreeNode, targetParentID string, levelPath []string) error {
	// 如果currentNode的ID是targetParentID,那么就直接加入该节点的children
	if currentNode.ID == targetParentID {
		currentNode.Children = append(currentNode.Children, node)
		return nil
	}
	// 根据levelPath获取下一个节点的ID
	nextNodeID := levelPath[0]

	// 查找currentNode的children中是否包含nextNodeID,不包含则执行最后的return返回错误信息,包含则继续递归查找
	for i := 0; i < len(currentNode.Children); i++ {
		if currentNode.Children[i].ID == nextNodeID {
			return insertToParentChildren(currentNode.Children[i], node, targetParentID, levelPath[1:])
		}
	}
	return fmt.Errorf("未找到父ID: nodeName:%s nodeID:%s parentID:%s\n", node.Name, node.ID, targetParentID)
}

// generateGroupTree 根据数
func generateGroupTree(data []*models.Group, rootParentID int) []*models.GroupTreeNode {

	var result []*models.GroupTreeNode
	var children []*models.Group

	// 分离根节点和子节点
	for i := 0; i < len(data); i++ {
		if data[i].ParentID == rootParentID {
			result = append(result, &models.GroupTreeNode{
				Name:     data[i].Name,
				ID:       strconv.Itoa(data[i].ID),
				Children: make([]*models.GroupTreeNode, 0),
			})
		} else {
			children = append(children, data[i])
		}
	}

	// 如果没有子节点,那么直接返回结果
	if len(children) == 0 {
		return result
	}
	// 遍历子节点,依次插入
	for i := 0; i < len(children); i++ {
		_t := strings.Split(children[i].LevelPath, "-")
		topParentID := _t[1]
		cNode := &models.GroupTreeNode{
			Name:     children[i].Name,
			ID:       strconv.Itoa(children[i].ID),
			Children: make([]*models.GroupTreeNode, 0),
		}
		child := children[i]
		for i := 0; i < len(result); i++ {
			// 找到子节点的顶级父
			if result[i].ID == topParentID {
				// 进行递归查找
				err := insertToParentChildren(result[i], cNode, strconv.Itoa(child.ParentID), _t[2:])
				if err != nil {
					log.Logger().Info(err.Error())
				}
				break // 递归完成直接 break 内层循环
			}
		}
	}
	return result
}

// QueryGroupAndSubGroupsUsersSvc 查询组及其子组下的所有用户
func (g *GroupService) QueryGroupAndSubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error) {

	groupIDs, err := g.groupRepo.QueryGroupIDAndSubGroupsID(data.Id, nil)
	if err != nil {
		return nil, err
	}

	users, err := g.userRepo.GetUserListRepo(models.User{}, nil, nil, groupIDs...)
	if err != nil {
		return nil, err
	}

	var result []*pb_user_v1.UserProto
	l := len(users)
	for i := 0; i < l; i++ {
		user := users[i]

		_user := &pb_user_v1.UserProto{
			Id:        &pb_user_v1.Index{Id: int64(user.ID)},
			UserName:  user.UserName,
			LoginName: user.LoginName,
			Mobile:    int64(user.Mobile),
			GroupId:   int64(user.GroupID),
			UserType:  int64(user.UserType),
			RoleIds:   nil,
		}
		// TODO: 查询组及其子组下的所有用户 -> 添加roleIDs
		result = append(result, _user)
	}

	return &pb_user_v1.Users{
		Users: result,
	}, nil
}
