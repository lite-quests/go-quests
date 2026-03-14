package http_server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

// TODO: Implement StartCatalogServer
// Read README.md for the instructions

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func StartCatalogServer(port int) error {
	// In-memory store
	products := make(map[string]Product)

	// Optional but recommended for concurrent safety
	var mu sync.RWMutex

	mux := http.NewServeMux()

	// GET /products
	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()

		list := make([]Product, 0, len(products))
		for _, p := range products {
			list = append(list, p)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)
	})

	// GET /products/{id}
	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		mu.RLock()
		product, exists := products[id]
		mu.RUnlock()

		if !exists {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	})

	// POST /products
	mux.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {
		var product Product

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if product.ID == "" {
			http.Error(w, "Product ID required", http.StatusBadRequest)
			return
		}

		mu.Lock()
		products[product.ID] = product
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
	})

	return http.ListenAndServe(":"+strconv.Itoa(port), mux)
}
