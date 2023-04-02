package collections

type dictionaryViaMap[K comparable, V any] struct {
	m map[K]V
}

var _ interface{ Dictionary[int, int] } = (*dictionaryViaMap[int, int])(nil)

func newDictionaryViaMap[K comparable, V any]() *dictionaryViaMap[K, V] {
	return &dictionaryViaMap[K, V]{
		m: make(map[K]V),
	}
}

func (d *dictionaryViaMap[K, V]) Clear() {
	d.m = make(map[K]V)
}

func (d *dictionaryViaMap[K, V]) Delete(key K) {
	delete(d.m, key)
}

func (d *dictionaryViaMap[K, V]) ForEach(callback func(entry Pair[K, V]) bool) {
	for k, v := range d.m {
		if !callback(NewPair(k, v)) {
			return
		}
	}
}

func (d *dictionaryViaMap[K, V]) Get(key K) (value V, ok bool) {
	value, ok = d.m[key]
	return value, ok
}

func (d *dictionaryViaMap[K, V]) Keys() Collection[K] {
	return Map[Pair[K, V]](d, func(pair Pair[K, V]) K {
		return pair.First
	})
}

func (d *dictionaryViaMap[K, V]) Set(key K, value V) {
	d.m[key] = value
}

func (d *dictionaryViaMap[K, V]) Size() int {
	return len(d.m)
}

func (d *dictionaryViaMap[K, V]) Values() Collection[V] {
	return Map[Pair[K, V]](d, func(pair Pair[K, V]) V {
		return pair.Second
	})
}

type setViaMap[T comparable] struct {
	m map[T]struct{}
}

var _ interface{ Set[int] } = (*setViaMap[int])(nil)

func newSetViaMap[T comparable]() *setViaMap[T] {
	return &setViaMap[T]{m: make(map[T]struct{})}
}

func (s *setViaMap[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *setViaMap[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *setViaMap[T]) Delete(value T) {
	delete(s.m, value)
}

func (s *setViaMap[T]) ForEach(callback func(T) bool) {
	for v := range s.m {
		if !callback(v) {
			break
		}
	}
}

func (s *setViaMap[T]) Has(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *setViaMap[T]) Size() int {
	return len(s.m)
}
