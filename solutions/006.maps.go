package solutions

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	m := make(map[string]int)
	return &Cache{data: m}
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (int, bool) {
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) Count() int {
	count := 0
	for _, v := range c.data {
		v++
		count++
	}
	return count
}

func (c *Cache) AllKeys() []string {
	var res []string
	for k, _ := range c.data {
		res = append(res, k)
	}
	return res
}

func (c *Cache) RemoveBelow(limit int) {
	for k, v := range c.data {
		if v < limit {
			delete(c.data, k)
		}
	}
}
