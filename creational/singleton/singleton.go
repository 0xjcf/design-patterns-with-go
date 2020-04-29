package singleton

// Operators to increment and decrement a value
type Operators interface {
	AddOne() int
	SubtractOne() int
}

// Counter contains numeric value that increments and decrements by one
type Counter struct {
	value int
}

var counter *Counter

// GetInstance returns an instance of the Counter
func GetInstance() *Counter {
	if counter == nil {
		counter = new(Counter)
	}
	return counter
}

// AddOne increments the Counters value by one
func (c *Counter) AddOne() int {
	c.value++
	return c.value
}

// SubtractOne decrements the Counter value by one
func (c *Counter) SubtractOne() int {
	c.value--
	return c.value
}
