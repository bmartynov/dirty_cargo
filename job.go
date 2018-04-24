package cargo

type Job interface {
	Id() string
	Payload() interface{}
}
