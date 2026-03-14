package mutexes

// TODO:Implement producer-consumer problem using mutexes
// Read README.md for the instructions
type Counter struct {
	items int
}

func (c *Counter) Produce(amount int) {
	c.items += amount
}

func (c *Counter) Consume(amount int) {
	c.items -= amount
}

func (c *Counter) GetCount() int {
	return c.items
}
