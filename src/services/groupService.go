package services
//
//import (
//	"context"
//	"crypto/md5"
//	"encoding/hex"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"gitee.com/grandeep/device-plugin/client"
//	"gitee.com/grandeep/device-plugin/src/services"
//	"gitee.com/grandeep/org-svc/src/models"
//	pb_user_v1 "gitee.com/grandeep/org-svc/src/proto/user/v1"
//	"gitee.com/grandeep/org-svc/src/repositories"
//	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
//	"gitee.com/grandeep/org-svc/logger"
//	"strconv"
//	"strings"
//	"time"
//)
//
//// GroupServiceInterface 组服务接口
//type GroupServiceInterface interface {
//	GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error)
//	GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error)
//	GroupUpdateSvc(ctx context.Context, data *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error)
//	QuotaUpdateSvc(ctx context.Context, data *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error)
//	GroupTreeQuerySvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error)
//	GroupDeleteSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error)
//	QueryGroupAndSubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error)
//	SetGroupQuotaUsedSvc(_ context.Context, data *pb_user_v1.SetGroupQuotaUsedRequest) (*pb_user_v1.GroupResponse, error)
//	QueryGroupIDAndSubGroupsIDSvc(_ context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupIDsResponse, error)
//	QueryQuotaByConditionSvc(_ context.Context, data *pb_user_v1.QueryQuotaByCondition) (*pb_user_v1.QueryQuotaByConditionResponse, error)
//	QuerySubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error)
//	GetAllGroup(_ context.Context, groupId *pb_user_v1.GroupID) (*pb_user_v1.Groups, error)
//	QueryQuotaSvc(_ context.Context, groupID *pb_user_v1.GroupID) (*pb_user_v1.QueryQuotaResponse, error)
//	QueryTopGroupExcludeSelfUsersSvc(_ context.Context, data *pb_user_v1.GroupIDWithPage) (*pb_user_v1.GroupUsersWithPage, error)
//}
//
//// GroupService 组服务,实现了 GroupServiceInterface
//type GroupService struct {
//	groupRepo         repositories.GroupRepoInterface
//	userRepo          repositories.UserRepoInterface
//	kubernetesService services.KubernetesServiceI
//	cfg               *config.Config
//}
//
//// NewGroupService GroupService 构造函数
//func NewGroupService(repos repositories.RepoI, cfg *config.Config) GroupServiceInterface {
//	etcdHosts, _ := cfg.GetString("etcdHost")
//
//	deviceClient := client.NewDeviceClient(strings.Split(etcdHosts, ";"), 2, time.Second*5)
//	return &GroupService{
//		groupRepo:         repos.GetGroupRepo(),
//		userRepo:          repos.GetUserRepo(),
//		kubernetesService: deviceClient.GetKubernetesService(),
//		cfg:               cfg,
//	}
//}
//
//// GroupAddSvc 添加组
//func (g *GroupService) GroupAddSvc(ctx context.Context, data *pb_user_v1.GroupAddRequest) (*pb_user_v1.GroupResponse, error) {
//	var err error
//	tx := g.groupRepo.GetTx()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//
//	// 非顶级父ID的组,查询其顶级父ID的namespace
//	// 非顶级组查询其顶级父ID的资源组和配额信息, 执行参数校验
//	var k8sNameSpace string
//	var topGroupQuota *models.QueryQuota
//
//	if data.ParentId != 0 {
//		// 通过父级查询顶级组信息
//		parentGroup, err := g.groupRepo.GroupQueryByIDRepo(data.ParentId, tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//		// 获取顶级父ID
//		var topGroupID int64
//		if parentGroup.ParentID == 0 {
//			topGroupID = int64(parentGroup.ID)
//		} else {
//			topGroupIDStr := strings.Split(parentGroup.LevelPath, "-")[1]
//			topGroupID, err = strconv.ParseInt(topGroupIDStr, 10, 64)
//			if err != nil {
//				tx.Rollback()
//				return nil, err
//			}
//		}
//
//		topGroup, err := g.groupRepo.GroupQueryByIDRepo(topGroupID, tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//		k8sNameSpace = topGroup.NameSpace
//
//		// 查询顶级组的资源组和配额
//		topGroupQuota , err = g.groupRepo.QueryQuota(topGroupID, tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//
//	} else {
//		m := md5.New()
//		m.Write([]byte(data.Name))
//		m.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
//		md5Str := hex.EncodeToString(m.Sum(nil))
//		k8sNameSpace = "org-svc-" + md5Str
//	}
//
//	newGroup := &models.Group{
//		Name:        data.Name,
//		ParentID:    int(data.ParentId),
//		NameSpace:   k8sNameSpace,
//		Description: data.Description,
//	}
//
//	err = g.groupRepo.GroupAddRepo(newGroup, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	group, err := g.groupRepo.GroupQueryByNameRepo(data.Name, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 相同配额类型不允许重复
//	var _share = false
//	var _nonShare = false
//
//	for i := 0; i < len(data.Quotas); i++ {
//		q := data.Quotas[i]
//
//		if q.IsShare == 1 {
//			if _share {
//				return nil, errors.New("不允许重复添加相同的共享类型(share=1)")
//			}
//			_share = true
//		} else if q.IsShare == 2 {
//			if _nonShare {
//				return nil, errors.New("不允许重复添加相同的共享类型(share=2)")
//			}
//			_nonShare = true
//		}
//	}
//
//	quotaTypeMap := map[string]models.ResourceType{
//		"cpu":    models.ResourceCpu,
//		"gpu":    models.ResourceGpu,
//		"memory": models.ResourceMemory,
//		"disk":   models.ResourceDisk,
//	}
//
//	l := len(data.Quotas)
//	var result = make([]*models.Quota, 0)
//	for i := 0; i < l; i++ {
//		q := data.Quotas[i]
//
//		valMap := map[string]int64{
//			"cpu":    q.Cpu,
//			"gpu":    q.Gpu,
//			"memory": q.Memory,
//		}
//
//		for kind, val := range valMap {
//			// 可以精简下列判断
//			var resourcesGroupID string
//			if topGroupQuota != nil {
//				if q.IsShare == 1 {
//					quotaInfo := topGroupQuota.ShareQuota
//					resourcesGroupID = quotaInfo.ResourcesGroupId
//
//					// 执行校验
//					switch quotaTypeMap[kind] {
//						case models.ResourceCpu:
//							if int(val) > quotaInfo.CpuTotal {
//								return nil, errors.New("资源类型:共享 非顶级组,CPU配额超出父级大小")
//							}
//						case models.ResourceGpu:
//							if int(val) > quotaInfo.GpuTotal {
//								return nil, errors.New("资源类型:共享 非顶级组,GPU配额超出父级大小")
//							}
//						case models.ResourceMemory:
//							if int(val) > quotaInfo.MemoryTotal {
//								return nil, errors.New("资源类型:共享 非顶级组,Memory配额超出父级大小")
//							}
//					}
//				} else if q.IsShare == 2 {
//					quotaInfo := topGroupQuota.NonShareQuota
//					resourcesGroupID = quotaInfo.ResourcesGroupId
//
//					// 执行校验
//					switch quotaTypeMap[kind] {
//					case models.ResourceCpu:
//						if int(val) > quotaInfo.CpuTotal {
//							return nil, errors.New("资源类型:独享 非顶级组,CPU配额超出父级大小")
//						}
//					case models.ResourceGpu:
//						if int(val) > quotaInfo.GpuTotal {
//							return nil, errors.New("资源类型:独享 非顶级组,GPU配额超出父级大小")
//						}
//					case models.ResourceMemory:
//						if int(val) > quotaInfo.MemoryTotal {
//							return nil, errors.New("资源类型:独享 非顶级组,Memory配额超出父级大小")
//						}
//					}
//				}
//			} else {
//				resourcesGroupID = q.ResourcesGroupId
//			}
//
//			tmp := &models.Quota{
//				IsShare:    int(q.IsShare),
//				ResourceID: resourcesGroupID,
//				Type:       quotaTypeMap[kind],
//				GroupID:    group.ID,
//				Total:      int(val),
//				Used:       0,
//			}
//			result = append(result, tmp)
//		}
//	}
//	// 磁盘配额单独添加
//	if topGroupQuota != nil {
//		if int(data.DiskQuotaSize) > topGroupQuota.DiskQuotaTotal {
//			return nil, errors.New("资源类型:磁盘配额 非顶级组,磁盘配额超出父级大小")
//		}
//	}
//
//	result = append(result, &models.Quota{
//		IsShare:    0,
//		ResourceID: "",
//		Type:       models.ResourceDisk,
//		GroupID:    group.ID,
//		Total:      int(data.DiskQuotaSize),
//		Used:       0,
//	})
//
//	err = g.groupRepo.QuotaAddRepo(result, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 只有顶级的组才创建 namespace
//	if group.ParentID == 0 {
//		_, err = g.kubernetesService.CreateNamespaceSvc(ctx, k8sNameSpace)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//	}
//	tx.Commit()
//	return &pb_user_v1.GroupResponse{Code: 0}, nil
//}
//
//// GroupQueryWithQuotaByConditionSvc 根据条件查询组信息和其配额信息
//func (g *GroupService) GroupQueryWithQuotaByConditionSvc(ctx context.Context, data *pb_user_v1.GroupQueryWithQuotaByConditionRequest) (*pb_user_v1.GroupQueryWithQuotaByConditionResponse, error) {
//
//	condition := &models.GroupQueryByCondition{
//		ID:       data.Id,
//		Name:     data.Name,
//		ParentID: data.ParentId,
//	}
//
//	querySlice, err := g.groupRepo.GroupQueryWithQuotaByConditionRepo(condition, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var groupData = make(map[int64]*pb_user_v1.GroupWithQuota)
//	var groupQuotaTemp = make(map[int64]map[int]map[string]*pb_user_v1.Quota)
//
//	l := len(querySlice)
//	for i := 0; i < l; i++ {
//		r := querySlice[i]
//		if _, ok := groupData[r.ID]; !ok {
//			groupData[r.ID] = new(pb_user_v1.GroupWithQuota)
//			groupData[r.ID].ParentId = r.ParentID
//			groupData[r.ID].Id = r.ID
//			groupData[r.ID].Name = r.Name
//			groupData[r.ID].Description = r.Description
//			groupData[r.ID].Quotas = make([]*pb_user_v1.Quota, 0)
//			levelPath := strings.Split(r.LevelPath, "-")
//			var topParent string
//			if levelPath[1] == "" {
//				topParent = "0"
//			} else {
//				topParent = levelPath[1]
//			}
//			topParentID, err := strconv.ParseInt(topParent, 10, 64)
//			if err != nil {
//				return nil, errors.New("转换顶级组ID失败:" + err.Error())
//			}
//			groupData[r.ID].TopParentId = topParentID
//		}
//
//		if models.ResourceType(r.Type) == models.ResourceDisk {
//			groupData[r.ID].DiskQuotaSize = int64(r.Total)
//			continue
//		}
//
//		// 对零散的配额数据进行分组
//		// 判断 group
//		if _, ok := groupQuotaTemp[r.ID]; !ok {
//			groupQuotaTemp[r.ID] = make(map[int]map[string]*pb_user_v1.Quota)
//		}
//		// 判断 is_share 配额类型,使用 is_share + resources_id 进行分组
//		if _, ok := groupQuotaTemp[r.ID][r.IsShare]; !ok {
//			groupQuotaTemp[r.ID][r.IsShare] = make(map[string]*pb_user_v1.Quota)
//		}
//		if _, ok := groupQuotaTemp[r.ID][r.IsShare][r.ResourceID]; !ok {
//			// 创建 quota 数据,并设置 is_share 和 resources_id 字段信息
//			groupQuotaTemp[r.ID][r.IsShare][r.ResourceID] = &pb_user_v1.Quota{
//				IsShare:          int64(r.IsShare),
//				ResourcesGroupId: r.ResourceID,
//			}
//		}
//
//		// 获取 quota 设置对应的值
//		quotaData := groupQuotaTemp[r.ID][r.IsShare][r.ResourceID]
//		switch models.ResourceType(r.Type) {
//		case models.ResourceCpu:
//			quotaData.Cpu = int64(r.Total)
//		case models.ResourceGpu:
//			quotaData.Gpu = int64(r.Total)
//		case models.ResourceMemory:
//			quotaData.Memory = int64(r.Total)
//		}
//	}
//
//	// 组装 quota 数据到 group 数据中
//	for groupID, v1 := range groupQuotaTemp {
//		quotas := make([]*pb_user_v1.Quota, 0)
//		for _, v2 := range v1 {
//			for _, q := range v2 {
//				quotas = append(quotas, q)
//			}
//		}
//		groupData[groupID].Quotas = quotas
//	}
//
//	// 结果 map -> slice
//	var result = make([]*pb_user_v1.GroupWithQuota, 0)
//	for _, g := range groupData {
//		result = append(result, g)
//	}
//
//	return &pb_user_v1.GroupQueryWithQuotaByConditionResponse{
//		Groups: result,
//	}, nil
//}
//
//// GroupUpdateSvc 组信息更新
//func (g *GroupService) GroupUpdateSvc(ctx context.Context, data *pb_user_v1.GroupUpdateRequest) (*pb_user_v1.GroupResponse, error) {
//	var tx = g.groupRepo.GetTx()
//
//	var parentGroupQuota *models.QueryQuota
//	if data.UseParentId {
//		fmt.Println("获取顶级组信息")
//		// 通过父级查询顶级组信息
//		parentGroup, err := g.groupRepo.GroupQueryByIDRepo(data.ParentId, tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//		// 获取顶级父ID
//		var topGroupID int64
//		if parentGroup.ParentID == 0 {
//			topGroupID = int64(parentGroup.ID)
//		} else {
//			topGroupIDStr := strings.Split(parentGroup.LevelPath, "-")[1]
//			topGroupID, err = strconv.ParseInt(topGroupIDStr, 10, 64)
//			if err != nil {
//				tx.Rollback()
//				return nil, err
//			}
//		}
//
//		topGroup, err := g.groupRepo.GroupQueryByIDRepo(topGroupID, tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//
//		fmt.Println("获取顶级组信息: ", topGroup)
//
//		// 查询顶级组的资源组和配额
//		parentGroupQuota, err = g.groupRepo.QueryQuota(int64(topGroup.ID), tx)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//	}
//	fmt.Println("处理配额")
//
//
//	// 处理配额
//	quotaTypeMap := map[string]models.ResourceType{
//		"cpu":    models.ResourceCpu,
//		"gpu":    models.ResourceGpu,
//		"memory": models.ResourceMemory,
//		"disk":   models.ResourceDisk,
//	}
//
//	// 资源组ID, 用于更新资源组ID
//	var resourcesGroupIDMap = make(map[int64]string)
//
//	// 展开上传数据的配额数据
//	quotasLen := len(data.Quotas)
//	var quotasUpdateData = make([]*models.QuotaUpdateRequest, 0)
//	for i := 0; i < quotasLen; i++ {
//		q := data.Quotas[i]
//
//		if _, ok := resourcesGroupIDMap[q.IsShare]; !ok {
//			resourcesGroupIDMap[q.IsShare] = q.ResourcesGroupId
//		}
//
//		var parentQuota *models.QueryQuotaInfo
//
//		if parentGroupQuota != nil {
//			if q.IsShare == 1 {
//				parentQuota = parentGroupQuota.ShareQuota
//			} else if q.IsShare == 2 {
//				parentQuota = parentGroupQuota.NonShareQuota
//			}
//		}
//
//		valMap := map[string]int64{
//			"cpu":    q.Cpu,
//			"gpu":    q.Gpu,
//			"memory": q.Memory,
//		}
//
//		for kind, val := range valMap {
//			if parentQuota != nil {
//				switch quotaTypeMap[kind] {
//
//				case models.ResourceCpu:
//					if int(val) > parentQuota.CpuTotal {
//						return nil, fmt.Errorf("资源类型: %d,CPU配额超出父级大小", parentQuota.IsShare)
//					}
//				case models.ResourceGpu:
//					if int(val) > parentQuota.GpuTotal {
//						return nil, fmt.Errorf("资源类型: %d,GPU配额超出父级大小", parentQuota.IsShare)
//					}
//				case models.ResourceMemory:
//					if int(val) > parentQuota.MemoryTotal {
//						return nil, fmt.Errorf("资源类型: %d,Memory配额超出父级大小", parentQuota.IsShare)
//					}
//				}
//			}
//
//			_tmp := &models.QuotaUpdateRequest{
//				GroupID:     data.Id,
//				IsShare:     q.IsShare,
//				ResourcesID: q.ResourcesGroupId,
//				QuotaType:   int64(quotaTypeMap[kind]),
//				Total:       val,
//			}
//			quotasUpdateData = append(quotasUpdateData, _tmp)
//		}
//	}
//
//	if parentGroupQuota != nil {
//		if int64(parentGroupQuota.DiskQuotaTotal) < data.DiskQuotaSize {
//			return nil, errors.New("磁盘配额超出父级大小")
//		}
//	}
//
//	// 磁盘配额单独添加
//	quotasUpdateData = append(quotasUpdateData, &models.QuotaUpdateRequest{
//		GroupID:     data.Id,
//		IsShare:     0,
//		ResourcesID: "",
//		QuotaType:   int64(models.ResourceDisk),
//		Total:       data.DiskQuotaSize,
//	})
//
//
//	// 更新配额信息
//	// 按照 组ID 共享类型 数据类型更新总量数据
//	// 不存在的配额进行创建
//	err := g.groupRepo.QuotaUpdateRepo(quotasUpdateData, data.Id, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 更新资源组信息
//	// 一个组只能有一个 共享 一个非共享, 这里按照共享类型更新 资源组ID
//	// 仅限不存在子级,不存在用户的组使用更新
//	// 如果原本有两个类型的配额(共享,非共享),执行删除操作
//	err = g.groupRepo.UpdateQuotaResourceID(data.Id, resourcesGroupIDMap, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 更新组表信息
//	d := &models.GroupUpdateRequest{
//		ID:          data.Id,
//		Name:        data.Name,
//		Description: data.Description,
//	}
//
//	// 必须传递 ParentID 才能校验配额, 否则不校验
//	if data.UseParentId {
//		d.ParentID = &data.ParentId
//	}
//
//	fmt.Println("更新组信息")
//	err = g.groupRepo.GroupUpdateRepo(d, tx)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	tx.Commit()
//
//	return &pb_user_v1.GroupResponse{Code: 0}, nil
//}
//
//// QuotaUpdateSvc 配额信息更新
//func (g *GroupService) QuotaUpdateSvc(ctx context.Context, data *pb_user_v1.QuotaUpdateRequest) (*pb_user_v1.GroupResponse, error) {
//
//	var err error
//	d := &models.QuotaUpdateRequest{
//		GroupID:     data.GroupId,
//		IsShare:     data.IsShare,
//		ResourcesID: data.ResourcesId,
//		QuotaType:   data.QuotaType,
//		Total:       data.Total,
//		Used:        data.Used,
//	}
//
//	_data := []*models.QuotaUpdateRequest{d}
//
//	err = g.groupRepo.QuotaUpdateRepo(_data, data.GroupId, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &pb_user_v1.GroupResponse{Code: 0}, nil
//}
//
//// GroupDeleteSvc 组删除(软删除)
//func (g *GroupService) GroupDeleteSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupResponse, error) {
//	if data.Id == 0 {
//		return nil, errors.New("组 id 不允许为空")
//	}
//
//	// 验证组下是否存在用户
//	users, err := g.userRepo.GetUserListRepo(models.User{GroupID: int(data.Id)}, nil, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	if len(users) != 0 {
//		return nil, fmt.Errorf("无法删除组,组内存在用户: %d 个", len(users))
//	}
//
//	// 验证是否存在下级组
//	groups, err := g.groupRepo.GroupListWithChangedLevelPathRepo(data.Id, nil)
//	if err != nil {
//		return nil, fmt.Errorf("无法删除组,查询是否包含下级组时错误: %s", err.Error())
//	}
//
//	if len(groups) != 0 {
//		return nil, fmt.Errorf("无法删除组,包含下级组: %d 个", len(groups))
//	}
//
//	// 删除组
//	err = g.groupRepo.GroupDeleteRepo(data.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb_user_v1.GroupResponse{Code: 0}, nil
//}
//
//// GroupTreeQuerySvc 组树查询
//func (g *GroupService) GroupTreeQuerySvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupTreeResponse, error) {
//
//	groupList, err := g.groupRepo.GroupListWithChangedLevelPathRepo(data.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	tree := generateGroupTree(groupList, int(data.Id))
//	if tree == nil {
//		return nil, errors.New("生成失败,结果为空")
//	}
//
//	jsonByte, err := json.Marshal(tree)
//	if err != nil {
//		return nil, err
//	}
//
//	res := &pb_user_v1.GroupTreeResponse{
//		TreeJson: jsonByte,
//	}
//
//	return res, nil
//}
//
//// insertToParentChildren generateGroupTree 中查找父节点的递归方法
//func insertToParentChildren(currentNode *models.GroupTreeNode, node *models.GroupTreeNode, targetParentID string, levelPath []string) error {
//	// 如果currentNode的ID是targetParentID,那么就直接加入该节点的children
//	if currentNode.ID == targetParentID {
//		currentNode.Children = append(currentNode.Children, node)
//		return nil
//	}
//	// 根据levelPath获取下一个节点的ID
//	nextNodeID := levelPath[0]
//
//	// 查找currentNode的children中是否包含nextNodeID,不包含则执行最后的return返回错误信息,包含则继续递归查找
//	for i := 0; i < len(currentNode.Children); i++ {
//		if currentNode.Children[i].ID == nextNodeID {
//			return insertToParentChildren(currentNode.Children[i], node, targetParentID, levelPath[1:])
//		}
//	}
//	return fmt.Errorf("未找到父ID: nodeName:%s nodeID:%s parentID:%s\n", node.Name, node.ID, targetParentID)
//}
//
//// generateGroupTree 根据数
//func generateGroupTree(data []*models.Group, rootParentID int) []*models.GroupTreeNode {
//
//	var result []*models.GroupTreeNode
//	var children []*models.Group
//
//	// 分离根节点和子节点
//	for i := 0; i < len(data); i++ {
//		if data[i].ParentID == rootParentID {
//			result = append(result, &models.GroupTreeNode{
//				Name:     data[i].Name,
//				ID:       strconv.Itoa(data[i].ID),
//				Children: make([]*models.GroupTreeNode, 0),
//			})
//		} else {
//			children = append(children, data[i])
//		}
//	}
//
//	// 如果没有子节点,那么直接返回结果
//	if len(children) == 0 {
//		return result
//	}
//	// 遍历子节点,依次插入
//	for i := 0; i < len(children); i++ {
//		_t := strings.Split(children[i].LevelPath, "-")
//		topParentID := _t[1]
//		cNode := &models.GroupTreeNode{
//			Name:     children[i].Name,
//			ID:       strconv.Itoa(children[i].ID),
//			Children: make([]*models.GroupTreeNode, 0),
//		}
//		child := children[i]
//		for i := 0; i < len(result); i++ {
//			// 找到子节点的顶级父
//			if result[i].ID == topParentID {
//				// 进行递归查找
//				err := insertToParentChildren(result[i], cNode, strconv.Itoa(child.ParentID), _t[2:])
//				if err != nil {
//					log.Logger.Info(err.Error())
//				}
//				break // 递归完成直接 break 内层循环
//			}
//		}
//	}
//	return result
//}
//
//// QueryGroupAndSubGroupsUsersSvc 查询组及其子组下的所有用户
//func (g *GroupService) QueryGroupAndSubGroupsUsersSvc(ctx context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error) {
//
//	groupIDs, err := g.groupRepo.QueryGroupIDAndSubGroupsID(data.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	users, err := g.userRepo.GetUserListRepo(models.User{}, nil, nil, groupIDs...)
//	if err != nil {
//		return nil, err
//	}
//
//	var result []*pb_user_v1.UserProto
//	l := len(users)
//	for i := 0; i < l; i++ {
//		user := users[i]
//
//		_user := &pb_user_v1.UserProto{
//			Id:        &pb_user_v1.Index{Id: int64(user.ID)},
//			UserName:  user.UserName,
//			LoginName: user.LoginName,
//			Mobile:    user.Mobile,
//			GroupId:   int64(user.GroupID),
//			UserType:  int64(user.UserType),
//			RoleIds:   nil,
//		}
//		// TODO: 查询组及其子组下的所有用户 -> 添加roleIDs
//		result = append(result, _user)
//	}
//
//	return &pb_user_v1.Users{
//		Users: result,
//	}, nil
//}
//
//// SetGroupQuotaUsedSvc 设置组配额已使用数值
//func (g *GroupService) SetGroupQuotaUsedSvc(_ context.Context, data *pb_user_v1.SetGroupQuotaUsedRequest) (*pb_user_v1.GroupResponse, error) {
//	var err error
//	d := &models.SetGroupQuotaRequest{
//		GroupID:   data.GroupId,
//		IsShare:   data.IsShare,
//		QuotaType: data.QuotaType,
//		Used:      data.Used,
//	}
//
//	err = g.groupRepo.SetGroupQuotaUsedRepo(d, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &pb_user_v1.GroupResponse{Code: 0}, nil
//}
//
//// QueryGroupIDAndSubGroupsIDSvc 查询组及其子组ID
//func (g *GroupService) QueryGroupIDAndSubGroupsIDSvc(_ context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.GroupIDsResponse, error) {
//
//	groupIDs, err := g.groupRepo.QueryGroupIDAndSubGroupsID(data.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb_user_v1.GroupIDsResponse{
//		Ids: groupIDs,
//	}, nil
//}
//
//// QueryQuotaByCondition 通过条件查询配额详情
//func (g *GroupService) QueryQuotaByConditionSvc(_ context.Context, data *pb_user_v1.QueryQuotaByCondition) (*pb_user_v1.QueryQuotaByConditionResponse, error) {
//
//	d := &models.QuotaQueryByCondition{
//		GroupID:    data.GroupId,
//		Type:       int(data.Type),
//		ResourceID: data.ResourceId,
//		IsShare:    int(data.IsShare),
//	}
//
//	resp, err := g.groupRepo.QuotaQueryByConditionRepo(d, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var result = make([]*pb_user_v1.QuotaRecord, 0)
//	l := len(resp)
//	for i := 0; i < l; i++ {
//		_item := resp[i]
//
//		_tmp := &pb_user_v1.QuotaRecord{
//			IsShare:    int64(_item.IsShare),
//			ResourceId: _item.ResourceID,
//			Type:       int64(_item.Type),
//			GroupId:    int64(_item.GroupID),
//			Total:      int64(_item.Total),
//			Used:       int64(_item.Used),
//		}
//		result = append(result, _tmp)
//	}
//
//	return &pb_user_v1.QueryQuotaByConditionResponse{Records: result}, nil
//}
//
//// QuerySubGroupsUsersSvc 查询子组下的用户
//func (g *GroupService) QuerySubGroupsUsersSvc(_ context.Context, data *pb_user_v1.GroupID) (*pb_user_v1.Users, error) {
//	groupIDs, err := g.groupRepo.QueryGroupIDAndSubGroupsID(data.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var groupIDSlice = make([]int64, 0, len(groupIDs)-1)
//
//	for _, v := range groupIDs {
//		if v == data.Id {
//			continue
//		}
//		groupIDSlice = append(groupIDSlice, v)
//	}
//
//	if len(groupIDSlice) == 0 {
//		return nil, errors.New("不存在下级组")
//	}
//
//	users, err := g.userRepo.GetUserListRepo(models.User{}, nil, nil, groupIDSlice...)
//	if err != nil {
//		return nil, err
//	}
//
//	var result []*pb_user_v1.UserProto
//	l := len(users)
//	for i := 0; i < l; i++ {
//		user := users[i]
//
//		_user := &pb_user_v1.UserProto{
//			Id:        &pb_user_v1.Index{Id: int64(user.ID)},
//			UserName:  user.UserName,
//			LoginName: user.LoginName,
//			Mobile:    user.Mobile,
//			GroupId:   int64(user.GroupID),
//			UserType:  int64(user.UserType),
//			RoleIds:   nil,
//		}
//		// TODO: 查询子组下的用户 -> 添加roleIDs
//		result = append(result, _user)
//	}
//
//	return &pb_user_v1.Users{
//		Users: result,
//	}, nil
//}
//
//// GetAllGroup 查询查询所有的组
//func (g *GroupService) GetAllGroup(_ context.Context, groupId *pb_user_v1.GroupID) (*pb_user_v1.Groups, error) {
//	var (
//		groups   []models.Group
//		pbGroups = &pb_user_v1.Groups{}
//	)
//	groups = g.groupRepo.GetAllGroup()
//	for _, v := range groups {
//		var (
//			pbGroup = &pb_user_v1.Group{
//				Id:                   int64(v.ID),
//				Name:                 v.Name,
//				Description:          v.Description,
//				ParentId:             int64(v.ParentID),
//				LevelPath:            v.LevelPath,
//				NameSpace:            v.NameSpace,
//				Status:               int64(v.Status),
//			}
//		)
//		pbGroups.Groups = append(pbGroups.Groups, pbGroup)
//	}
//	return pbGroups, nil
//}
//
//// QueryQuota 查询所有配额信息
//func (g *GroupService) QueryQuotaSvc(_ context.Context, groupID *pb_user_v1.GroupID) (*pb_user_v1.QueryQuotaResponse, error) {
//
//	res, err := g.groupRepo.QueryQuota(groupID.Id, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	result := &pb_user_v1.QueryQuotaResponse{
//		ShareQuota: &pb_user_v1.QueryQuotaInfo{
//			IsShare:          res.ShareQuota.IsShare,
//			ResourcesGroupId: res.ShareQuota.ResourcesGroupId,
//			GpuTotal:         int64(res.ShareQuota.GpuTotal),
//			CpuTotal:         int64(res.ShareQuota.CpuTotal),
//			MemoryTotal:      int64(res.ShareQuota.MemoryTotal),
//			GpuUsed:          int64(res.ShareQuota.GpuUsed),
//			CpuUsed:          int64(res.ShareQuota.CpuUsed),
//			MemoryUsed:       int64(res.ShareQuota.MemoryUsed),
//		},
//		NonShareQuota: &pb_user_v1.QueryQuotaInfo{
//			IsShare:          res.NonShareQuota.IsShare,
//			ResourcesGroupId: res.NonShareQuota.ResourcesGroupId,
//			GpuTotal:         int64(res.NonShareQuota.GpuTotal),
//			CpuTotal:         int64(res.NonShareQuota.CpuTotal),
//			MemoryTotal:      int64(res.NonShareQuota.MemoryTotal),
//			GpuUsed:          int64(res.NonShareQuota.GpuUsed),
//			CpuUsed:          int64(res.NonShareQuota.CpuUsed),
//			MemoryUsed:       int64(res.NonShareQuota.MemoryUsed),
//		},
//		DiskQuotaTotal: int64(res.DiskQuotaTotal),
//		DiskQuotaUsed:  int64(res.DiskQuotaUsed),
//	}
//
//	return result, nil
//}
//
//// QueryTopGroupExcludeSelfUsersSvc 查询顶级组及其下面子组的用户,不包含传入的组
//func (g *GroupService) QueryTopGroupExcludeSelfUsersSvc(ctx context.Context, data *pb_user_v1.GroupIDWithPage) (*pb_user_v1.GroupUsersWithPage, error) {
//
//	group, err := g.GroupQueryWithQuotaByConditionSvc(ctx, &pb_user_v1.GroupQueryWithQuotaByConditionRequest{
//		Id: []int64{data.Id},
//	})
//	if err != nil {
//		return nil, err
//	}
//	if len(group.Groups) != 1 {
//		return nil, fmt.Errorf("查询组信息失败: %v", group.Groups)
//	}
//	_group := group.Groups[0]
//
//
//
//	groupIDs, err := g.groupRepo.QueryGroupIDAndSubGroupsID(_group.TopParentId, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	var groupIDSlice = make([]int64, 0, len(groupIDs)-1)
//
//	for _, v := range groupIDs {
//		if v == data.Id {
//			continue
//		}
//		groupIDSlice = append(groupIDSlice, v)
//	}
//
//	if len(groupIDSlice) == 0 {
//		return &pb_user_v1.GroupUsersWithPage{}, nil
//	}
//
//	fmt.Println("DEBUG", _group.TopParentId, groupIDSlice, data.Id)
//
//	page := &models.Page{
//		PageSize:  int(data.PageSize),
//		PageNum:   int(data.Page),
//	}
//
//	users, err := g.userRepo.GetUserListRepo(models.User{}, page, nil, groupIDSlice...)
//	if err != nil {
//		return nil, err
//	}
//
//	var result []*pb_user_v1.UserProto
//	l := len(users)
//	for i := 0; i < l; i++ {
//		user := users[i]
//
//		_user := &pb_user_v1.UserProto{
//			Id:        &pb_user_v1.Index{Id: int64(user.ID)},
//			UserName:  user.UserName,
//			LoginName: user.LoginName,
//			Mobile:    user.Mobile,
//			GroupId:   int64(user.GroupID),
//			UserType:  int64(user.UserType),
//			RoleIds:   nil,
//		}
//		result = append(result, _user)
//	}
//
//	return &pb_user_v1.GroupUsersWithPage{
//		Users:     result,
//		Page:      int64(page.PageNum),
//		PageSize:  int64(page.PageSize),
//		TotalPage: int64(page.TotalPage),
//		TotalNum:  page.Total,
//	}, nil
//}