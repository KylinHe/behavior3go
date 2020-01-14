package actions

import (
	b3 "github.com/KylinHe/behavior3go"
	. "github.com/KylinHe/behavior3go/core"
)

type Failer struct {
	Action
}

func (this *Failer) OnTick(tick *Tick) b3.Status {
	return b3.FAILURE
}
