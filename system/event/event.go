package event

import (
	"errors"
	"game/def"
	"game/framework/resource"
)

type Attribute int32

const (
	AttributeNone Attribute = iota
	AttributeType

	AttributeItemAddId
	AttributeItemAddCount
)

type eventID int32

const (
	eventNone eventID = iota
	eventItemAdd
)

var errAttributeNotFound = errors.New("attribute not found")

type ItemAddEvent struct {
	Tid   int32
	Count int64
}

func GetAttrValue(i *ItemAddEvent, id Attribute) (int64, error) {
	switch id {
	case AttributeType:
		return int64(eventItemAdd), nil
	case AttributeItemAddId:
		return int64(i.Tid), nil
	case AttributeItemAddCount:
		return i.Count, nil

	default:
		return 0, errAttributeNotFound
	}
}

type sEvent struct {
}

func (s *sEvent) Add(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEvent) RemoveByTid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEvent) RemoveByUid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEvent) SetByTid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEvent) SetByUid(owner resource.Owner, lp resource.LogParam, res []*def.Resource) (entities *def.EntityGroup, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *sEvent) Collect(owner resource.Owner, res []*def.Resource) (entities *def.EntityGroup) {
	e := &ItemAddEvent{}
	_, _ = GetAttrValue(e, 0)
	return
}

func (s *sEvent) CheckSpace(owner resource.Owner, res []*def.Resource) error {
	//TODO implement me
	panic("implement me")
}
