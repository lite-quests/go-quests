package maps

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	// TODO: initialize the cache
	// Read README.md for the instructions
	return &Cache{}
}

func (c *Cache) Set(key string, value int) {
	// TODO: implement
	// Read README.md for the instructions
}

func (c *Cache) Get(key string) (int, bool) {
	// TODO: implement
	// Read README.md for the instructions
	return 0, false
}

func (c *Cache) Delete(key string) {
	// TODO: implement
	// Read README.md for the instructions
}

func (c *Cache) Count() int {
	// TODO: implement
	// Read README.md for the instructions
	return 0
}

func (c *Cache) AllKeys() []string {
	// TODO: implement
	// Read README.md for the instructions
	return []string{}
}

func (c *Cache) RemoveBelow(limit int) {
	// TODO: implement
	// Read README.md for the instructions
}
