package main

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string{
	return s.name
}

//提供给内部设置名字

func (s *StateInfo) setName(name string) {
	s.name = name
}

func (s *StateInfo) EnableTransit() bool{
	return false
}

func (s *StateInfo) OnBegin() {

}

func (s *StateInfo) OnEnd() {

}

func (s *StateInfo) CanTransitTo(name string) bool{
	return true
}
