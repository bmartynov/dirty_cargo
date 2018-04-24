package cargo

type Result interface {
	Id() string
	Value() interface{}
	Error() error
}

type ResultHandler interface {
	Handle(Result)
}
