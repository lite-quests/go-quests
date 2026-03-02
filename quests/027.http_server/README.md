# HTTP Server: Product Catalog API

## Concept

Think of an HTTP Server like a restaurant waiter. They stand by taking orders (requests), routing those orders to the right kitchen station (handlers), and then bringing the food (responses) back to the customer.

Historically, Go's standard library server was very basic. But starting in Go 1.22+, `http.ServeMux` received a massive upgrade. It now supports enhanced routing patterns directly in the standard library, making it easy to build full REST APIs without needing any third-party frameworks!

- **`http.NewServeMux()`**: Creates a new multiplexer (the "waiter").
- **`mux.HandleFunc("GET /products", handler)`**: Registers a handler function for a specific HTTP method and path pattern.
- **`mux.HandleFunc("GET /products/{id}", handler)`**: Captures wildcard variables in the path (like `{id}`). You can read the value inside the handler using `r.PathValue("id")`.
- **`json.NewDecoder(r.Body).Decode(&obj)`**: Parses JSON data sent by the client into a Go struct.
- **`json.NewEncoder(w).Encode(obj)`**: Converts a Go struct back into JSON and sends it to the client.

## Quest

### Objective
You will build an **E-Commerce Product Catalog API**.

### Requirements

Implement the function:
`func StartCatalogServer(port int) error`

1. Create a `map[string]Product` to store products in memory. 
2. Initialize an `http.ServeMux`.
3. Create a route for **`GET /products`**:
   - Write all map values as a JSON list back to the client.
4. Create a route for **`GET /products/{id}`**:
   - Get the `{id}` using `r.PathValue("id")`.
   - If the product exists in the map, return it as JSON.
   - If it doesn't exist, return a `404 Not Found` using `http.Error(w, "Not Found", http.StatusNotFound)`.
5. Create a route for **`POST /products`**:
   - Decode the JSON request body into a `Product` struct.
   - Store it in your map using `product.ID` as the key.
   - Reply with a `201 Created` status code (`w.WriteHeader(http.StatusCreated)`).
6. Start the server via `http.ListenAndServe` using the formatted `port` string and your mux.

### Inputs
- `port`: The port to start the server on, e.g., `8080`.

### Outputs
- The Go function should block and return an `error` if `http.ListenAndServe` fails.
- The HTTP server routes output JSON strings and appropriate status codes (`200 OK`, `201 Created`, `404 Not Found`).

### Models
```go
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
```

### Testing
```bash
go test -v ./quests/1004.http_server
```
