package htmlgo

//<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">`
type THead struct {
	HttpEquiv string
	Content   string
	strTitle  string
}

func NewHead() *THead {
	return &THead{
		HttpEquiv: "Content-Type",
		Content:   `text/html; charset=utf-8`,
	}
}

func (m *THead) String() string {
	return `<head>` + `<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">` +
		`<meta name="viewport" content="width=device-width,initial-scale=1.0">` +
		`<title>` + m.strTitle + `</title>` + `</head>`
}

func (m *THead) SetTitle(strTitle string) {
	m.strTitle = strTitle
}
