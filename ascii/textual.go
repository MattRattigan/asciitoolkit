package ascii

type Textual interface {
	Text | []rune | []string
}

type Text interface {
	~string | ~rune
}
