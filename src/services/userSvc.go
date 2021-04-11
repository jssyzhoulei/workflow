package services

// user 所提供的能力抽象成接口以便复用
type UserInterface interface {
	// 从组下随机获取用户
	GetUserRandomWithGroup(groupId int) (int, error)
	// 获取组下管理人员 即角色数据权限为 Admin的用户
	GetAdminWithGroup(groupId int) (int, error)
	// 获取组下普通人员 即角色数据权限为 Self的用户
	GetGeneralWithGroup(groupId int) (int, error)
	IsGroupExist(groupId int) bool
	IsUserExist(groupId int) bool
}
