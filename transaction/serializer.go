package transaction

// Serializer interface
type Serializer interface {
	Decode(input []byte) (*Transaction, error)
	Encode(input []byte) ([]byte, error)
}
