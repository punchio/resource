package resource

import (
	"context"
	"errors"
	"fmt"
	"game/def"
)

type Owner interface {
	context.Context
	GetOwnerInfo() (def.EOwnerType, int32)
	CheckResource(def.EResourceType) bool
}

type LogParam interface {
	GetAction() int32
	GetParams() []int64
}

func register(ownerType def.EOwnerType, resourceType def.EResourceType, method Method) error {
	tRouter, err := mgr.get(ownerType)
	if err != nil {
		tRouter = &typeRouter{route: make(map[def.EResourceType]Method)}
		mgr.route[ownerType] = tRouter
	}

	_, err = tRouter.get(resourceType)
	if err == nil {
		return fmt.Errorf("resource register repeat, owner:%d, resource:%d", ownerType, resourceType)
	}

	tRouter.route[resourceType] = method
	return nil
}

var mgr ownerRouter
var errIllegalParam = errors.New("illegal param")

type ownerRouter struct {
	route map[def.EOwnerType]*typeRouter
}

func (o *ownerRouter) get(typ def.EOwnerType) (*typeRouter, error) {
	r, ok := o.route[typ]
	if !ok {
		return nil, fmt.Errorf("resource owner not found, type:%d", typ)
	}
	return r, nil
}

type typeRouter struct {
	route map[def.EResourceType]Method
}

func (i *typeRouter) get(typ def.EResourceType) (Method, error) {
	r, ok := i.route[typ]
	if !ok {
		return nil, fmt.Errorf("resource type not found, type:%d", typ)
	}
	return r, nil
}
func (i *typeRouter) getModifier(typ def.EResourceType, method methodType) (modifier, error) {
	r, ok := i.route[typ]
	if !ok {
		return nil, fmt.Errorf("resource type not found, type:%d", typ)
	}
	switch method {
	case methodTypeAdd:
		if r.Add != nil {
			return r.Add, nil
		}
	case methodTypeRemoveTid:
		if r.RemoveByTid != nil {
			return r.RemoveByTid, nil
		}
	case methodTypeRemoveUid:
		if r.RemoveByUid != nil {
			return r.RemoveByUid, nil
		}
	case methodTypeSetTid:
		if r.SetByTid != nil {
			return r.SetByTid, nil
		}
	case methodTypeSetUid:
		if r.SetByUid != nil {
			return r.SetByUid, nil
		}
	}
	return nil, fmt.Errorf("resource type method not implement, type:%d, method:%d", typ, method)
}

type modifier func(owner Owner, lp LogParam, res []*def.Resource) (entities *def.EntityGroup, err error)
type methodType int32

const (
	methodTypeAdd methodType = iota
	methodTypeRemoveTid
	methodTypeRemoveUid
	methodTypeSetTid
	methodTypeSetUid
)

func modify(owner Owner, lp LogParam, res []*def.Resource, method methodType) (*def.EntityGroup, error) {
	r, g, err := prepare(owner, res)
	if err != nil {
		return nil, err
	}

	var m modifier
	var oneGroup *def.EntityGroup
	resultGroup := &def.EntityGroup{}
	for rType, resGroup := range g {
		m, err = r.getModifier(rType, method)
		if err != nil {
			return nil, err
		}
		oneGroup, err = m(owner, lp, resGroup)
		if err != nil {
			return nil, err
		}

		resultGroup.Append(oneGroup)
	}
	return resultGroup, nil
}

func checkEnough(owner Owner, res []*def.Resource) error {
	r, g, err := prepare(owner, res)
	if err != nil {
		return err
	}

	filters := make([]*def.Filter, len(res))
	for i, v := range res {
		filters[i] = &def.Filter{
			Type:  v.Type,
			Attr:  def.EResourceAttributeTid,
			Value: int64(v.Id),
			Mode:  def.CompareModeEqual,
		}
	}

	var tRouter Method
	for rType, resGroup := range g {
		tRouter, err = r.get(rType)
		if err != nil {
			return err
		}
		resGroupCount := tRouter.Collect(owner, filters).GetEachCount()
		for _, v := range resGroup {
			count := resGroupCount[v.Id]
			if count < v.Count {
				return fmt.Errorf("resource not enough, id:%d, need:%d, has:%d",
					v.Id, v.Count, count)
			}
		}
	}
	return nil
}

func checkSpace(owner Owner, res []*def.Resource) error {
	r, g, err := prepare(owner, res)
	if err != nil {
		return err
	}
	var tRouter Method
	for rType, resGroup := range g {
		tRouter, err = r.get(rType)
		if err != nil {
			return err
		}
		err = tRouter.CheckSpace(owner, resGroup)
		if err != nil {
			return err
		}
	}
	return nil
}

func collect(owner Owner, filter []*def.Filter) (*def.EntityGroup, error) {
	r, g, err := prepare(owner, filter)
	if err != nil {
		return nil, err
	}

	entities := &def.EntityGroup{}

	var tRouter Method
	var oneResGroup *def.EntityGroup
	for rType, resGroup := range g {
		tRouter, err = r.get(rType)
		if err != nil {
			return nil, err
		}
		oneResGroup = tRouter.Collect(owner, resGroup)
		entities.Append(oneResGroup)
	}
	return entities, nil
}

func prepare[T interface{ GetType() def.EResourceType }](owner Owner, res []T) (*typeRouter, map[def.EResourceType][]T, error) {
	if len(res) == 0 {
		return nil, nil, errIllegalParam
	}
	typ, _ := owner.GetOwnerInfo()
	v, err := mgr.get(typ)
	if err != nil {
		return nil, nil, err
	}

	resGroup := makeGroup(res)
	return v, resGroup, nil
}

func makeGroup[T interface{ GetType() def.EResourceType }](res []T) map[def.EResourceType][]T {
	m := make(map[def.EResourceType][]T)
	for _, v := range res {
		m[v.GetType()] = append(m[v.GetType()], v)
	}
	return m
}
