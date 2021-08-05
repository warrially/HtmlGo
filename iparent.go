package htmlgo

type IParent interface {
	AddChild(IChild)
}

type IChild interface {
	String() string
}
