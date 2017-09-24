package golang

type langStruct struct {
	ID   int    `xml:"id" json:"id"`
	Lang string `xml:"lang" json:"lang"`
	Book Book   `xml:"book"`
}

type Book struct {
	Author string `xml:"aurthor", attr`
	Name   string `xml:"name",omitempty`
}
