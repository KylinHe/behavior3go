package core

import (
	b3 "github.com/KylinHe/behavior3go"
	. "github.com/KylinHe/behavior3go/config"
)

//子树，通过Name关联树ID查找
type SubTree struct {
	Action
	properties map[string]interface{}
}

func (this *SubTree) Initialize(setting *BTNodeCfg) {
	this.properties = setting.Properties
	this.Action.Initialize(setting)
}

func (this *SubTree) OnTick(tick *Tick) b3.Status {

	//使用子树，必须先SetSubTreeLoadFunc
	//子树可能没有加载上来，所以要延迟加载执行
	sTree := subTreeLoadFunc(this.GetName())
	if nil == sTree {
		return b3.ERROR
	}
	if tick.GetTarget() == nil {
		panic("SubTree tick.GetTarget() nil !")
	}
	tar := tick.GetTarget()
	for key, value := range this.properties {
		tick.Blackboard.Set(key, value, sTree.GetID(), "")
	}
	//	glog.Info("subtree: ", this.treeName, " id ", player.id)
	return sTree.Tick(tar, tick.Blackboard)
}

var subTreeLoadFunc func(string) *BehaviorTree

//获取子树的方法
func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}
