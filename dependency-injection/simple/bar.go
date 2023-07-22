package simple

// provider set or grouping

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{BarRepository: barRepository}
}