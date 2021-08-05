package htmlgo

import "strconv"

type TDiv struct {
	strBackgroundColor string
	strFlex            string
	strDisplay         string
	strDirection       string
	strAlignItem       string
	strMargin          string
	strPadding         string
	strHeight          string
	children           []IChild
}

func NewDiv() *TDiv {
	return &TDiv{}
}
func NewRow() *TDiv {
	return &TDiv{
		// strBackgroundColor: "red",
		// strFlex:      "1",
		strDisplay:   "flex",
		strDirection: "row",
		strAlignItem: "center",
		children:     []IChild{},
	}
}

func NewColumn() *TDiv {
	return &TDiv{
		// strBackgroundColor: "red",
		// strFlex:      "1",
		strDisplay:   "flex",
		strDirection: "column",
		strAlignItem: "center",
		children:     []IChild{},
	}
}

// AddTo
func (m *TDiv) AddTo(pParent IParent) *TDiv {
	pParent.AddChild(m)
	return m
}

//
func (m *TDiv) String() string {
	strResult := `<div style="`
	if m.strAlignItem != "" {
		strResult += `align-items: ` + m.strAlignItem + `;`
	}
	if m.strHeight != "" {
		strResult += `height: ` + m.strHeight + `;`
	}
	if m.strMargin != "" {
		strResult += `margin: ` + m.strMargin + `;`
	}
	if m.strPadding != "" {
		strResult += `padding: ` + m.strPadding + `;`
	}

	if m.strBackgroundColor != "" {
		strResult += `background-color: ` + m.strBackgroundColor + `;`
	}

	if m.strFlex != "" {
		strResult += `flex: ` + m.strFlex + `;`
	}

	if m.strDisplay != "" {
		strResult += `display: ` + m.strDisplay + ";"
	}

	if m.strDirection != "" {
		strResult += "flex-direction: " + m.strDirection
	}

	strResult += `">`

	for _, v := range m.children {
		strResult += v.String()
	}

	return strResult + `</div>`
}

//
func (m *TDiv) SetBackground(strBackgroundColor string) *TDiv {
	m.strBackgroundColor = strBackgroundColor
	return m
}

func (m *TDiv) SetFlex(nIndex int) *TDiv {
	m.strFlex = strconv.Itoa(nIndex)
	return m
}

func (m *TDiv) SetMargin(strMargin string) *TDiv {
	m.strMargin = strMargin
	return m
}

func (m *TDiv) SetPadding(strPadding string) *TDiv {
	m.strPadding = strPadding
	return m
}

func (m *TDiv) SetHeight(strHeight string) *TDiv {
	m.strHeight = strHeight
	return m
}

func (m *TDiv) SetAlignItem(strAlignItem string) *TDiv {
	m.strAlignItem = strAlignItem
	return m

}

func (m *TDiv) AddChild(pChild IChild) {
	m.children = append(m.children, pChild)
}
