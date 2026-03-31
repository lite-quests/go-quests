package json

import "encoding/json"

// Product represents a product in an online store.
type Product struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Price   float64  `json:"price"`
	InStock bool     `json:"in_stock"`
	Tags    []string `json:"tags,omitempty"`
}

// MarshalProduct encodes a Product into a JSON byte slice.
func MarshalProduct(p Product) ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalProduct decodes a JSON byte slice into a Product.
func UnmarshalProduct(data []byte) (Product, error) {
	var p Product
	err := json.Unmarshal(data, &p)
	return p, err
}

// MarshalProducts encodes a slice of Products into a JSON byte slice.
func MarshalProducts(products []Product) ([]byte, error) {
	return json.Marshal(products)
}

// UnmarshalProducts decodes a JSON byte slice into a slice of Products.
func UnmarshalProducts(data []byte) ([]Product, error) {
	var products []Product
	err := json.Unmarshal(data, &products)
	return products, err
}

// PrettyMarshal encodes a Product into an indented (pretty-printed) JSON byte slice.
func PrettyMarshal(p Product) ([]byte, error) {
	return json.MarshalIndent(p, "", "\t")
}
