package transaction

type transactionService struct {
	transactionRepository Repository
}

// NewTransactionService initializer
func NewTransactionService(transactionRepository Repository) Service {
	return &transactionService{
		transactionRepository,
	}
}

func (t *transactionService) Find(id int) (*Transaction, error) {
	return t.transactionRepository.Find(id)
}

func (t *transactionService) Store(transaction *Transaction) error {
	return t.transactionRepository.Store(transaction)
}
