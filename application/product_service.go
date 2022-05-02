package application

// import "errors"

// injecao de dependencia, o service pode receber qualquer tipo de struct que implemento o IProductPersistance
type ProductService struct {
	Persistance IProductPersistance
}

func NewProductService(persistence IProductPersistance) *ProductService {
	return &ProductService{Persistance: persistence}
}

func (s *ProductService) GetById(id string) (IProduct, error) {
	// nao sabemos de onde ele vai pegar ou como, só sabemos que tem esse metodo
	product, err := s.Persistance.GetById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()

	if err != nil {
		// produto em branco
		return &Product{}, err
	}

	result, err := s.Persistance.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProductService) Enable(product IProduct) (IProduct, error) {
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}
	
	result, err := s.Persistance.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, err
} 

func (s *ProductService) Disable(product IProduct) (IProduct, error) {
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}
	
	result, err := s.Persistance.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, err
} 