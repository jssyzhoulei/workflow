package services

import (
	"errors"
	"github.com/jssyzhoulei/workflow/src/models"
	"github.com/jssyzhoulei/workflow/src/repositories"
	"gorm.io/gorm"
)

// 节点对象排序是要根据update time排序
// create time 节点后续可能会调整顺序位置
type workNodeSvc struct {
	repo *repositories.WorkRepo
	// skip name 跟 work node request的map
	nameObjMap map[string]*models.WorkNodeRequest
	// 节点名称不允许重复
	nameSet     map[string]bool
	workflowId  int
	waitDelList map[int]bool
}

func NewNodeSvc(repo *repositories.WorkRepo, getNodeFunc func() []models.WorkNode) *workNodeSvc {

	list := getNodeFunc()
	waitDelList := make(map[int]bool)
	if list != nil {
		for _, i := range list {
			if i.ID != 0 {
				waitDelList[i.ID] = false
			}
		}
	}
	return &workNodeSvc{
		repo:        repo,
		nameObjMap:  make(map[string]*models.WorkNodeRequest),
		nameSet:     make(map[string]bool),
		waitDelList: waitDelList,
	}
}

func (wns workNodeSvc) CreateOrUpdate(tx *gorm.DB, wfs []*models.WorkNodeRequest) error {
	err := wns.parseNodes(tx, 0, wfs)
	if err == nil {
		err = wns.deleteNodes()
	}
	return err
}

func (wns workNodeSvc) parseNodes(tx *gorm.DB, parentId int, wfs []*models.WorkNodeRequest) error {
	var lastId int
	if parentId == 0 {
		if wfs == nil || len(wfs) == 0 || wfs[0].Type != models.WorkNodeTypeHead {
			return errors.New("缺少head节点")
		}
	}
	for index, i := range wfs {
		err := wns.checkWorkNode(i)
		if err != nil {
			return err
		}
		i.ParentID = parentId
		i.LastID = lastId
		// 添加/更新当前节点
		err = wns.repo.SaveOrCreateWorkNode(tx, &i.WorkNode)
		if err != nil {
			return err
		}
		lastId = i.ID
		if len(i.Children) > 0 {
			err := wns.parseNodes(tx, i.ID, i.Children)
			if err != nil {
				return err
			}
		}
		// 保存需要更新skip id的work node map
		if i.SkipName != "" {
			wns.nameObjMap[i.SkipName] = wfs[index]
		}
		if wns.nameObjMap[i.Name] != nil {
			// 更新skip id值
			wns.nameObjMap[i.Name].SkipID = i.ID
			err = wns.repo.SaveWorkNode(tx, &wns.nameObjMap[i.Name].WorkNode)
			if err != nil {
				return err
			}
		}
		delete(wns.waitDelList, i.ID)
	}
	return nil
}

func (wns *workNodeSvc) checkWorkNode(wf *models.WorkNodeRequest) error {

	if wns.workflowId == 0 {
		wns.workflowId = wf.WorkFLowID
		if wns.workflowId == 0 {
			return errors.New("work node need work flow id ")
		}
	} else {
		wf.WorkFLowID = wns.workflowId
	}
	if _, ok := wns.nameSet[wf.Name]; ok {
		return errors.New("节点名称重复")
	}
	if _, ok := wns.nameSet[wf.SkipName]; ok {
		return errors.New("不允许反向跳转节点")
	}
	wns.nameSet[wf.Name] = true

	switch wf.Type {
	case models.WorkNodeTypeGeneral, models.WorkNodeTypeHead:
		// 此两种情况不支持嵌套
		if wf.Children != nil && len(wf.Children) != 0 {
			wf.Children = nil
		}
	default:
		if wf.Children == nil || len(wf.Children) == 0 {
			return errors.New("复杂节点的子节点为空")
		}
	}
	if wf.AuditType == models.AuditTypeAbsolute {
		if wf.PrincipleID == 0 {
			// 此时principle 为user id
			return errors.New("审批人类型固定时，principle需要指定负责人")
		} else if wf.AuditType == models.AuditTypeUAnyGroup {
			if wf.PrincipleID == 0 {
				// 此时principle id为 group id
				return errors.New("审批人类型为组下人时，principle需要指定组织id")
			}
		}
	}
	return nil
}

// 清除多余的node 节点
func (wns workNodeSvc) deleteNodes() error {
	var toD []int
	for i := range wns.waitDelList {
		toD = append(toD, i)
	}
	return wns.repo.DelWorkNode(toD)
}

func buildNodeTree(nodes []models.WorkNode) []*models.WorkNodeRequest {
	var (
		// id node map
		nodeMap = make(map[int]*models.WorkNodeRequest)
		// parent id node list map
		nodeMapList = make(map[int][]*models.WorkNodeRequest)
		topList     = make([]*models.WorkNodeRequest, 0)
	)
	for _, i := range nodes {
		newNode := models.WorkNodeRequest{}
		newNode.WorkNode = i
		if i.ParentID == 0 {
			topList = append(topList, &newNode)
		} else {
			if _, ok := nodeMapList[i.ParentID]; ok {
				nodeMapList[i.ParentID] = append(nodeMapList[i.ParentID], &newNode)
			} else {
				nodeMapList[i.ParentID] = []*models.WorkNodeRequest{&newNode}
			}
		}
		nodeMap[i.ID] = &newNode
	}
	for parentId, list := range nodeMapList {
		nodeMap[parentId].Children = list
	}
	return topList
}
