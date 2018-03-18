package interfaces

// Mappable is converted to gss.Feed
type Mappable interface {
	// ToJSON convert gss.Feed
	ToJSON() ([]byte, error)
}
