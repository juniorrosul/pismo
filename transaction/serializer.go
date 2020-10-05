package transaction

// Serializer interface
type Serializer interface {
	Decode(input []byte) (*Model, error)
	Encode(input []byte) ([]byte, error)
}
