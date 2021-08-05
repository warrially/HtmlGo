package htmlgo

import "io/ioutil"

type THTML struct {
	pHead *THead
	pBody *TBody
}

func NewHTML() *THTML {
	return &THTML{
		pHead: NewHead(),
		pBody: NewBody(),
	}
}

func (m *THTML) String() string {
	return m.document() + `<html style="height: 100%;width:100%;display: flex;>` + m.pHead.String() + m.pBody.String() + `</html>`
}

func (m *THTML) document() string {
	return `<!DOCTYPE html>`
}

// 设置标题
func (m *THTML) SetTitle(strTitle string) {
	m.pHead.SetTitle(strTitle)
}

func (m *THTML) SaveToFile(strFile string) {
	ioutil.WriteFile(strFile, []byte(m.String()), 0666)
}

func (m *THTML) GetBody() *TBody {
	return m.pBody
}
