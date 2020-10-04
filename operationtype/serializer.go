package operationtype

// Serializer interface
type Serializer interface {
	Decode(input []byte) (*OperationType, error)
	Encode(input *OperationType) ([]byte, error)
}
