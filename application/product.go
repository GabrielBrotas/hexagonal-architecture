package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// exeutada por padrao primeiro
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type IProductService interface {
	GetByID(id string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error) 
}

type IProductReader interface {
	GetById(id string) (IProduct, error)
}

type IProductWriter interface {
	Save(product IProduct) (IProduct, error)
}

// composição de interface
type IProductPersistance interface {
	IProductReader
	IProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string `valid:"uuidv4"`
	Name   string `valid:"required"`
	Status string `valid:"float,optional"`
	Price  float64 `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		ID: uuid.New().String(),
		Status: DISABLED,
	}

	// retornar a localização do ponteiro e não o valor em si
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Price < 0 {
		return false, errors.New("product price must be greater than 0")
	}

	if p.Status == DISABLED {
		return false, nil
	}
	
	if p.Status == ENABLED {
		return true, nil
	}

	if p.Status == "" {
		return false, errors.New("invalid status")
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("invalid status")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

// se a primeira letra da funcao for mínuscula significa que é um método privado
// letra maiuscula public
func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("Product price must be greater than 0")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("Product price must be equal to 0")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
