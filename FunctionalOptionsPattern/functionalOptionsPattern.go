package main

import "fmt"

type identity struct {
	Name         string
	IdCardNumber string
	Age          int
	Email        string
}

//可选参数列表
type IdentityOptions struct {
	Age   int
	Email string
}

type IdentityOption func(options *IdentityOptions)

func WithAge(age int) IdentityOption {
	return func(options *IdentityOptions) {
		options.Age = age
	}
}

func WithEmail(email string) IdentityOption {
	return func(options *IdentityOptions) {
		options.Email = email
	}
}

var defaultIdentity = IdentityOptions{18, "hdu_wangyujiang@163.com"}

func NewIdentity(name string, idCard string, opts ...IdentityOption) *identity {
	id := defaultIdentity
	for _, o := range opts {
		o(&id)
	}
	return &identity{name, idCard, id.Age, id.Email}
}
func main() {
	x := NewIdentity("WANG", "340", WithAge(20), WithEmail("hdu@hdu.edu.cn"))
	fmt.Println(x)

}
