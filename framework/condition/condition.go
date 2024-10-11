package condition

import "game/def"

type Expr struct {
	collectors []*def.Resource
	logicMode  def.CompareMode
}
