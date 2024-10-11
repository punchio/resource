package def

type EOwnerType int32

const (
	EOwnerTypeRole EOwnerType = 1
)

type EResourceType int32

const (
	EResourceTypeNone EResourceType = iota
	EResourceTypeItem
	EResourceTypeEvent
)

type EResourceAttribute int32

const (
	EResourceAttributeNone EResourceAttribute = iota
	// 通用类型
	EResourceAttributeTid

	EResourceAttributeItemClass EResourceAttribute = 100
	EResourceAttributeItemQuality

	EResourceAttributeEventItem EResourceAttribute = 200
)

type Resource struct {
	Type  EResourceType
	Id    int32
	Count int64
	Param any
}

func (r *Resource) GetType() EResourceType {
	return r.Type
}

type Entity struct {
	Type  EResourceType
	Tid   int32 // 如果是删除，通过uid拿不到对应数据
	Uid   int32
	Count int64
}

type EntityGroup struct {
	entities []*Entity
}

func (e *EntityGroup) Append(other *EntityGroup) {
	e.entities = append(e.entities, other.entities...)
}

func (e *EntityGroup) GetCount() int64 {
	if e == nil {
		return 0
	}
	count := int64(0)
	for _, entity := range e.entities {
		count += entity.Count
	}
	return count
}

func (e *EntityGroup) GetEachCount() map[int32]int64 {
	if e == nil {
		return nil
	}
	count := make(map[int32]int64)
	for _, entity := range e.entities {
		count[entity.Tid] += entity.Count
	}
	return count
}

type CompareMode int32

const (
	CompareModeEqual   = 0
	CompareModeGreater = 1
)

type Filter struct {
	Type  EResourceType
	Attr  EResourceAttribute
	Value int64
	Mode  CompareMode
}

func (f *Filter) GetType() EResourceType {
	return f.Type
}
