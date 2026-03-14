package http_client

// Read README.md for the instructions

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// TODO: Implement FetchProducts
func FetchProducts(apiURL string) ([]Product, error) {
	return nil, nil
}

// TODO: Implement CreateProduct
func CreateProduct(apiURL string, p Product) error {
	return nil
}

// TODO: Implement FetchProduct
func FetchProduct(apiURL string, id string) (Product, bool, error) {
	return Product{}, false, nil
}

// TODO: Implement FetchMultipleProducts
func FetchMultipleProducts(apiURL string, ids []string) (map[string]Product, error) {
	return nil, nil
}
