# HTTP Client: Inventory Sync Client

## Concept

When communicating with external REST APIs, handling JSON encoding/decoding and HTTP status codes correctly is essential. Just like we built the server previously, now we are building the client that talks to it.

**The Golden Rule of Go HTTP Clients:** You should **never use the default HTTP client** (`http.Get`, `http.Post`) in production. The default client has _no timeout_, meaning if the server hangs, your application will hang forever waiting for a response! Instead, you must always instantiate a custom `http.Client`.

Finally, dealing with networks often means dealing with latency. If you need to fetch 10 items, doing it one by one is slow. Go makes it incredibly easy to issue multiple HTTP requests simultaneously using **Goroutines**. Think of a `sync.WaitGroup` like a school field-trip chaperone—they wait at the exit until every single student (Goroutine) has returned before letting the buses leave.

- **`client := &http.Client{Timeout: 5 * time.Second}`**: Creates a safe HTTP client.
- **`client.Get(url)`**: Issues a `GET` request using the custom client.
- **`json.NewDecoder(resp.Body).Decode(&obj)`**: Decodes a JSON response back into a Go struct.
- **`sync.WaitGroup`**: Wait for a collection of Goroutines to finish their work.
- **`sync.Mutex`**: Safely lock shared data (like a Map) so two Goroutines don't try to write to it at the exact same microsecond, which would cause a panic.

## References

- [Go by Example: HTTP Clients](https://gobyexample.com/http-clients)
- [Go by Example: WaitGroups](https://gobyexample.com/waitgroups)
- [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)

## Quest

### Objective

You will build an **Inventory Sync Client** that interacts with the Product Catalog API you built in the previous quest.

### Requirements

Implement the following functions:

**1. `FetchProducts(apiURL string) ([]Product, error)`**

- First, instantiate a custom client with a 5-second timeout: `client := &http.Client{Timeout: 5 * time.Second}`
- Make a `GET` request to `apiURL + "/products"` using your custom `client`.
- If the status isn't 200 OK, return an error.
- Decode the JSON response into a list of `Product`s and return it.

**2. `CreateProduct(apiURL string, p Product) error`**

- Instantiate the custom client with a 5-second timeout.
- Encode the given `Product` `p` into JSON. Use `bytes.Buffer` and `json.NewEncoder()`.
- Make a `POST` request to `apiURL + "/products"` with `"application/json"` content type using your custom `client`.
- If the status code is not 201 Created (or 200 OK), return an error.

**3. `FetchProduct(apiURL string, id string) (Product, bool, error)`**

- Instantiate the custom client with a 5-second timeout.
- Make a `GET` request to `apiURL + "/products/" + id` using your custom `client`.
- If the status code is 404 Not Found, return an empty `Product`, `false` (indicating not found), and `nil` error.
- If the status code is 200 OK, decode the JSON response into a `Product` and return it with `true` (found) and `nil` error.
- Return an error for any other status codes or networking issues.

**4. `FetchMultipleProducts(apiURL string, ids []string) (map[string]Product, error)`**

- This function must fetch all requested IDs simultaneously using **Goroutines**.
- Initialize a `map[string]Product` to store successfully fetched products.
- Use a `sync.WaitGroup` to wait for all Goroutines.
- Use a `sync.Mutex` to lock/unlock the map when inserting into it concurrently to avoid data races.
- Inside each Goroutine, use `FetchProduct` to get the item. Only add the product to your map if it was successfully found (ignore 404s/errors for the sake of this quest).

### Inputs

- `apiURL`: The full URL to the Product Catalog API (e.g., `"http://localhost:8080"`).
- `id` (for `FetchProduct`): The ID of the product to fetch.
- `ids` (for `FetchMultipleProducts`): A slice of IDs to fetch concurrently.
- `p` (for `CreateProduct`): The `Product` object to create.

### Outputs

- `FetchProducts`: Returns `[]Product` and `error`.
- `CreateProduct`: Returns an `error` if the creation failed or the API returned non-success.
- `FetchProduct`: Returns `(Product, bool, error)`, where the boolean indicates if the product was found (200 OK vs 404 Not Found).
- `FetchMultipleProducts`: Returns a `map[string]Product` of the successful results.

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
go test -v ./quests/028.http_client
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestInventorySyncClient
=== RUN   TestInventorySyncClient/FetchProducts
=== RUN   TestInventorySyncClient/CreateProduct_Success
=== RUN   TestInventorySyncClient/CreateProduct_Failure
=== RUN   TestInventorySyncClient/FetchProduct_Found
=== RUN   TestInventorySyncClient/FetchProduct_NotFound
=== RUN   TestInventorySyncClient/FetchMultipleProducts
--- PASS: TestInventorySyncClient (0.00s)
    --- PASS: TestInventorySyncClient/FetchProducts (0.00s)
    --- PASS: TestInventorySyncClient/CreateProduct_Success (0.00s)
    --- PASS: TestInventorySyncClient/CreateProduct_Failure (0.00s)
    --- PASS: TestInventorySyncClient/FetchProduct_Found (0.00s)
    --- PASS: TestInventorySyncClient/FetchProduct_NotFound (0.00s)
    --- PASS: TestInventorySyncClient/FetchMultipleProducts (0.00s)
PASS
```
