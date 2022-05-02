package cli

import (
	"fmt"

	"github.com/GabrielBrotas/hexagonal-architeture/application"
)

func Run(service application.IProductService, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
		case "create":
			product, err := service.Create(productName, price)
			
			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with Name %s has been created with price %f and status %s", 
				product.GetID(), 
				product.GetName(), 
				product.GetPrice(),
				product.GetStatus(),
			)
		
		case "enable":
			product, err := service.GetById(productId)

			if err != nil {
				return result, err
			}

			_, err = service.Enable(product)
			
			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with Name %s has been Enabled", 
				product.GetID(), 
				product.GetName(), 
			)
		
		case "disable":
			product, err := service.GetById(productId)

			if err != nil {
				return result, err
			}

			_, err = service.Disable(product)
			
			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID %s with Name %s has been Disabled", 
				product.GetID(), 
				product.GetName(), 
			)
		
		default:
			product, err := service.GetById(productId)

			if err != nil {
				return result, err
			}

			result = fmt.Sprintf(
				"Product ID: %s \nName: %s\nPrice: %f\nStatus: %s", 
				product.GetID(), 
				product.GetName(),
				product.GetPrice(),
				product.GetStatus(), 
			)
	}


	return result, nil
}