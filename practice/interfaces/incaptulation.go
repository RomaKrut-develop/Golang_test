package somepackage

type SomeInterface interface {
	SomeMethod(in string) (out string)
}

type someStruct struct {
	prefix string
}

func (s SomeStruct) SomeMethod(in string) (out string){
	return s.prefix + in
}

func NewSomeInterface(prefix string) SomeInterface {
	return someStruct{prefix}
}