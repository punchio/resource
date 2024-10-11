package item

import (
	"game/def"
	"game/framework/resource"
)

func init() {
	if err := resource.Register(def.EOwnerTypeRole, def.EResourceTypeItem, &sItem{}); err != nil {
		panic(err)
	}
}

type sItem struct {
}

func (s *sItem) Collect(owner resource.Owner, filter []*def.Filter) (entities *def.EntityGroup) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) Add(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) RemoveByTid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) RemoveByUid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) SetByTid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) SetByUid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sItem) CheckSpace(owner resource.Owner, res []*def.Resource) error {
	//TODO implement me
	panic("implement me")
}
