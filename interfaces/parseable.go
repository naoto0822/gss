package interfaces

// Parseable is Parser interface for feed
type Parseable interface {
	// Parse parsing XML byte and return struct implmented interfaces.Mappable
	Parse([]byte) (*Mappable, error)
}
