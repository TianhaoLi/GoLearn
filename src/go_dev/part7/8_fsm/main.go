package main

import "fmt"

//闲置状态
type IdleState struct {
	StateInfo
}

func (i *IdleState) OnBegin() {
	fmt.Println("IdleState Begin")
}

func (i *IdleState) OnEnd() {
	fmt.Println("IdleState End")
}

//移动状态
type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin(){
	fmt.Println("MOveState begin")
}

func (m *MoveState) EnableSameTransit() bool{
	return true
}

//跳跃状态
type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("JumpState begin")
}

//跳跃状态不能转移到移动
func (j *JumpState) CanTransitTo(name string) bool{
	return name != "MoveState"
}

func main() {
	//状态管理器
	sm := NewStateManager()

	//响应转移通知
	sm.OnChange = func(from, to State) {
		//打印转移流向
		fmt.Printf("%s --> %s \n\n",StateName(from),StateName(to))
	}

	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	//在不同状态转移
	transitAndReport(sm,"IdleState")

	transitAndReport(sm,"MoveState")

	transitAndReport(sm,"MoveState")

	transitAndReport(sm,"JumpState")

	transitAndReport(sm,"JumpState")

	transitAndReport(sm,"IdleState")
}

func transitAndReport(sm *StateManager,target string)  {
	if err := sm.Transit(target); err != nil{
		fmt.Printf("FAILED! %s --> %s ,%s \n\n",sm.CurrState().Name(),target,err.Error())
	}
}