package uploader

type Translator interface {
}

type translator struct {
}

func NewTranslator() Translator {
	return &translator{}
}
