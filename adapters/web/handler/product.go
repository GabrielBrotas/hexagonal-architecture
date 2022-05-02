package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrotas/hexagonal-architeture/adapters/dto"
	"github.com/GabrielBrotas/hexagonal-architeture/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.IProductService) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		write.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)

		id := vars["id"]

		product, err := service.GetById(id)

		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(write).Encode(product)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		write.Header().Set("Content-Type", "application/json")

		var productDto dto.Product

		err := json.NewDecoder(req.Body).Decode(&productDto)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(write).Encode(product)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		write.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)

		id := vars["id"]

		product, err := service.GetById(id)

		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)
		
		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			write.Write(jsonError(err.Error()))
			return
		}
	
		err = json.NewEncoder(write).Encode(result)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.IProductService) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		write.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)

		id := vars["id"]

		product, err := service.GetById(id)

		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)
		
		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			write.Write(jsonError(err.Error()))
			return
		}
	
		err = json.NewEncoder(write).Encode(result)

		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
