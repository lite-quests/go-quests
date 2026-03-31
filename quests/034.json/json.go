package json

// TODO: Implement the functions below.
// Read README.md for the instructions.

// Product represents a product in an online store.
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	InStock  bool    `json:"in_stock"`
	Tags     []string `json:"tags,omitempty"`
}

// MarshalProduct encodes a Product into a JSON byte slice.
func MarshalProduct(p Product) ([]byte, error) {
	// TODO: Implement
	return nil, nil
}

// UnmarshalProduct decodes a JSON byte slice into a Product.
func UnmarshalProduct(data []byte) (Product, error) {
	// TODO: Implement
	return Product{}, nil
}

// MarshalProducts encodes a slice of Products into a JSON byte slice.
func MarshalProducts(products []Product) ([]byte, error) {
	// TODO: Implement
	return nil, nil
}

// UnmarshalProducts decodes a JSON byte slice into a slice of Products.
func UnmarshalProducts(data []byte) ([]Product, error) {
	// TODO: Implement
	return nil, nil
}

// PrettyMarshal encodes a Product into an indented (pretty-printed) JSON byte slice.
// Use "" as prefix and "\t" as indent.
func PrettyMarshal(p Product) ([]byte, error) {
	// TODO: Implement
	return nil, nil
}
