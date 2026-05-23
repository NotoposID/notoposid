package categories

type Service interface {
	// Add business logic methods here if needed
}

type categoryService struct{}

func NewService() Service {
	return &categoryService{}
}
