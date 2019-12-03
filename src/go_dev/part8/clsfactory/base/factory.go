package base

type Class interface {
	Do()
}

var(
	factoryByName = make(map[string] func() Class)
)

func Register(name string, factory func() Class)  {
	factoryByName[name] = factory;
}

//根据名称创建对应的类

func Create(name string) Class {
	if f,ok := factoryByName[name]; ok{
		return f()
	}else {
		panic("name not found")
	}
}
