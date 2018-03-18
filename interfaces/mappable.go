package interfaces

type Mappable interface {
	ToJSON() ([]byte, error)
}
