package transaction

type transactionService struct {
	repository Repository
	serializer Serializer
}

func (t *transactionService) Decode(input []byte) (*Model, error) {
	return t.serializer.Decode(input)
}

func (t *transactionService) Store(transaction *Model) error {
	return t.repository.Store(transaction)
}
