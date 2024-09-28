package animal

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	dog := &Dog{name: name}

	return dog
}
