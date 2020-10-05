package operationtype

// Serializer interface
type Serializer interface {
	Decode(input []byte) (*Model, error)
	Encode(input *Model) ([]byte, error)
}
