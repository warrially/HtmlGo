package htmlgo

import "strconv"

//<font color="gray" size="2">%d年</font>
type TFont struct {
	strSize  string
	strColor string
	strText  string
	bP       bool
}

func NewFont() *TFont {
	return &TFont{}
}

func (m *TFont) AddTo(pParent IParent) *TFont {
	pParent.AddChild(m)
	return m
}

//
func (m *TFont) String() string {
	strResult := "<font"

	if m.strColor != "" {
		strResult += ` color="` + m.strColor + `"`
	}

	if m.strSize != "" {
		strResult += ` size="` + m.strSize + `"`
	}

	strResult += `>`

	strResult = strResult + m.strText + `</font>`

	if m.bP {
		strResult = `<p>` + strResult + `</p>`
	}
	return strResult
}

//
func (m *TFont) SetSize(nSize int) *TFont {
	m.strSize = strconv.Itoa(nSize)
	return m
}

//
func (m *TFont) SetColor(strColor string) *TFont {
	m.strColor = strColor
	return m
}

func (m *TFont) SetText(strText string) *TFont {
	m.strText = strText
	return m
}

func (m *TFont) P() *TFont {
	m.bP = true
	return m
}
