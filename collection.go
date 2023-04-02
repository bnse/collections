package collections

type collectionImpl[T any] struct {
	forEach func(callback func(T) bool)
	size    func() int
}

func (c collectionImpl[T]) ForEach(callback func(T) bool) {
	c.forEach(callback)
}

func (c collectionImpl[T]) Size() int {
	return c.size()
}

// Map maps all the values in a collection to a new collection.
func Map[from, to any](c Collection[from], f func(from) to) Collection[to] {
	return collectionImpl[to]{
		forEach: func(callback func(to) bool) {
			c.ForEach(func(v from) bool {
				return callback(f(v))
			})
		},
		size: func() int {
			return c.Size()
		},
	}
}
