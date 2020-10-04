package account

// Serializer interface
type Serializer interface {
	Decode(input []byte) (*Account, error)
	Encode(input *Account) ([]byte, error)
}
