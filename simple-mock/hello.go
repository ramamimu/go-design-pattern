package hello

type TextService interface {
	GetHelloWorld() string
}

type Text struct {
	TextService
}

func NewText(t TextService) *Text {
	return &Text{
		TextService: t,
	}
}

func (t *Text) GetHelloWorld() string {
	return "haloo"
}
