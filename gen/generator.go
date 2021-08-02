package gen

type Generator interface {
	FromFile(fileName string) error
}
