package htmlgo

type TBody struct {
	children []IChild
}

func NewBody() *TBody {
	return &TBody{}
}

func (m *TBody) String() string {
	strResult := `<body  style="display: flex; flex :1; flex-direction: column;">`

	for _, v := range m.children {
		strResult += v.String()
	}

	return strResult + "</body>"
}

func (m *TBody) AddChild(pChild IChild) {
	m.children = append(m.children, pChild)
}
