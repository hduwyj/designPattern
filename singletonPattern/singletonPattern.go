package singletonPattern

type signleton struct {
	count int
}

var instance *signleton

func GetInstance() *signleton{
	if instance == nil{
		instance = new(signleton)
	}
	return  instance
}

func (i *signleton) AddOne() int{
	i.count++
	return i.count
}