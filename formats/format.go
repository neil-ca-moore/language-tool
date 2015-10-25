package formats

type Format interface {
	MakeWith(str string, loc string) (string, error)
}
