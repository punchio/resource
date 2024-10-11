package resource

import "game/def"

// Method 资源操作方式
// 有些资源类型是抽象出来的一类资源，没有修改的功能，如Add、Deduct，只有读的功能
type Method interface {
	// Add 添加物品
	// 可以实现普通的根据配置id添加物品，group：Item；id：配置id；count：添加数量
	// 也可以实现根据属性添加物品，group：ItemAttr；id：品质、类别等属性id；count：添加数量
	// 如抽象的事件资源组，就不支持添加删除
	Add(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
	// RemoveByTid 删除物品
	// 通过配置id删除物品与 Add 同理
	RemoveByTid(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
	// RemoveByUid
	// 所有通过uid删除物品含义都一样
	RemoveByUid(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
	// SetByTid 设置物品数量
	// 通过配置id设置物品数量与 Add 同理
	SetByTid(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
	// SetByUid
	// 所有通过uid设置物品含义都一样
	SetByUid(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
	// CheckSpace 检查剩余空间
	// 检查剩余空间只能根据基础的配置id检查，ItemAttr不能检查，uid也不能检查，没有意义。
	CheckSpace(owner Owner, res []*def.Resource) error
	// Collect 收集物品
	// 收集所有满足 filter 的物品
	// 实现事件属性，type：EventAttr；id：道具配置id、品质，养成线id、等级等等
	Collect(owner Owner, filter []*def.Filter) (entities *def.EntityGroup)
}

func Register(ownerType def.EOwnerType, resourceType def.EResourceType, method Method) error {
	return register(ownerType, resourceType, method)
}

func Add(owner Owner, lp LogParam, res []*def.Resource) (*def.EntityGroup, error) {
	return modify(owner, lp, res, methodTypeAdd)
}

func DeductByTid(owner Owner, lp LogParam, res []*def.Resource) (*def.EntityGroup, error) {
	return modify(owner, lp, res, methodTypeRemoveTid)
}
func DeductByUid(owner Owner, lp LogParam, res []*def.Resource) (*def.EntityGroup, error) {
	return modify(owner, lp, res, methodTypeRemoveUid)
}

func SetByTid(owner Owner, lp LogParam, res []*def.Resource) (*def.EntityGroup, error) {
	return modify(owner, lp, res, methodTypeSetTid)
}
func SetByUid(owner Owner, lp LogParam, res []*def.Resource) (*def.EntityGroup, error) {
	return modify(owner, lp, res, methodTypeSetUid)
}

func CheckEnough(owner Owner, res []*def.Resource) error {
	return checkEnough(owner, res)
}

func CheckSpace(owner Owner, res []*def.Resource) error {
	return checkSpace(owner, res)
}

func Collect(owner Owner, filter []*def.Filter) (*def.EntityGroup, error) {
	return collect(owner, filter)
}
