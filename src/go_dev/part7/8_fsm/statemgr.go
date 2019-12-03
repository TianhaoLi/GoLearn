package main

import "errors"

type StateManager struct {
	//已添加的状态
	stateByName map[string]State
	
	//状态改变时的回调
	OnChange func(from,to State)
	
	curr State
}

//状态信息
func (sm *StateManager) Add(s State) {
	name := StateName(s)


	//将s转换为能设置名字的接口 然后调用该接口
	s.(interface{
		setName(name string)
	}).setName(name)

	if sm.Get(name) != nil{
		panic("duplicate state:"+name)
	}

	sm.stateByName[name] = s
}

func (sm *StateManager) Get(name string) State{
	if v,ok := sm.stateByName[name]; ok{
		return v
	}
	return nil
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName:make(map[string]State),
	}
}


var ErrStateNotFound = errors.New("state not found")

//状态转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")

//不能转移到指定的状态
var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) CurrState() State{
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool{
	if sm.curr == nil{
		return true
	}

	//相同状态无需转换
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit(){
		return false
	}

	//使用当前状态检查是否能转移到指定名字状态
	return sm.curr.CanTransitTo(name)
}

//转移到制定状态
func (sm *StateManager) Transit(name string) error{
	//获取目标状态
	next := sm.Get(name)

	//目标不存在
	if next == nil{
		return ErrStateNotFound
	}

	//记录转移前状态
	pre := sm.curr

	//当前有状态
	if sm.curr !=nil{
		//相同状态无需转移
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit(){
			return ErrForbidSameStateTransit
		}

		//不能转移到目标状态
		if !sm.curr.CanTransitTo(name){
			return ErrCannotTransitToState
		}

		//结束当前状态
		sm.curr.OnEnd()
	}

	//切换当前状态
	sm.curr = next

	//调用新的状态的开始
	sm.curr.OnBegin()

	//通知回调
	if sm.OnChange !=nil{
		sm.OnChange(pre,sm.curr)
	}

	return nil
}
