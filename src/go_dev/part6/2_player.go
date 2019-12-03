package main

type Player struct {
	currPos Vec2
	targetPos Vec2
	speed float32
}
//设置玩家移动的目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos()Vec2 {
	return p.currPos
}

func (p *Player) IsArrived()bool {
	//计算但前位置与目标位置不超过移动的步长
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

//更新玩家的位置上
func (p *Player) Update() {
	if !p.IsArrived(){
		dir := p.targetPos.Sub(p.currPos).Normalize()

		newPos := p.currPos.Add(dir.Scale(p.speed))

		//移动完成后更新当前的位置、
		p.currPos = newPos
	}
}

func NewPlayer(speed float32) * Player {
	return &Player{
		speed: speed,
	}
}
