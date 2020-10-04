package operationtype

import "errors"

var (
	ErrOperationTypeNotFound = errors.New("OperationType not found")
	ErrOperationTypeStore    = errors.New("OperationType not stored")
)

type operationTypeService struct {
	operationTypeRepository Repository
}

// NewOperationTypeService service initializer
func NewOperationTypeService(operationTypeRepository Repository) Service {
	return &operationTypeService{
		operationTypeRepository,
	}
}

func (ot *operationTypeService) Find(id int) (*OperationType, error) {
	return ot.operationTypeRepository.Find(id)
}

func (ot *operationTypeService) Store(operationType *OperationType) error {
	return ot.operationTypeRepository.Store(operationType)
}
