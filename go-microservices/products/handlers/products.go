// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// Consumes:
// - application
//
// Produces:
// -application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	protos "products/currency"
	"products/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:response productsResponse
type productsResponseWrapper struct {
	// All produts in the system
	// in:body
	Body []data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from database
	// in:path
	// required:true
	ID int `json:"id"`
}

type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProducts(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//
//	200:productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	rw.Header().Add("Content-Type", "application/json")
	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		return
	}
	lp := data.GetProducts()
	for _, p := range lp {
		p.Price = p.Price * resp.Rate
	}
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	p.l.Printf("Prod: %#v", prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(int(id), &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
	p.l.Printf("Prod: %#v", prod)
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.ParseJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		if err = prod.Validate(); err != nil {
			p.l.Println("[ERROR] validation product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error Validating product %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

// swagger:route DELETE /products/{id} products deleteProduct
// deletes a product
// responses:
// 201: noContent
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle Delete Product", id)
	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	rw.Write([]byte(fmt.Sprintf("Product %v deleted", id)))
	p.l.Printf("Product %v deleted", id)
}
