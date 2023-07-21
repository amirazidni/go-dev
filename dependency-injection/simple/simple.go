package simple

import "errors"

type SimpleRepository struct {
	Error bool // error for detect wire error
}

// provider
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

// provider depend to simple repo pointer
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error { // error for detect wire error
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
