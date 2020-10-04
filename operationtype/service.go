package operationtype

// Service interface
type Service interface {
	Find(id int) (*OperationType, error)
}
