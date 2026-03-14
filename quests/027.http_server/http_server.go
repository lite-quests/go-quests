package http_server

// TODO: Implement StartCatalogServer
// Read README.md for the instructions

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func StartCatalogServer(port int) error {
	// In-memory store

	// GET /products

	// GET /products/{id}

	// POST /products
	return nil
}
