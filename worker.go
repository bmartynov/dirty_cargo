package cargo

type Worker interface {
	Process(job Job) Result
}
