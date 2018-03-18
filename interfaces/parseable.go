package interfaces

type Parseable interface {
	Parse([]byte) (*Mappable, error)
}
