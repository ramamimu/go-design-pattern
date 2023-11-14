package hello

type TextService interface {
	GetHelloWorld() string
}

type Text struct {
	service TextService
}

func NewText(s TextService) *Text {
	return &Text{service: s}
}

func GetHelloWorld() string {
	return "hello world"
}
