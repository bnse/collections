package collections

import "github.com/tidwall/btree"

type dictViaBTree[K, V any] struct {
	b    *btree.Generic[Pair[K, V]]
	less func(K, K) bool
}

var _ interface{ Dictionary[int, int] } = (*dictViaBTree[int, int])(nil)

func newDictionaryViaBTree[K, V any](less func(K, K) bool) *dictViaBTree[K, V] {
	return &dictViaBTree[K, V]{
		b: btree.NewGeneric(func(a, b Pair[K, V]) bool {
			return less(a.First, b.First)
		}),
		less: less,
	}
}

func (d *dictViaBTree[K, V]) Clear() {
	less := d.less
	d.b = btree.NewGeneric(func(a, b Pair[K, V]) bool {
		return less(a.First, b.First)
	})
	d.less = less
}

func (d *dictViaBTree[K, V]) Delete(key K) {
	var zero V
	d.b.Delete(NewPair(key, zero))
}

func (d *dictViaBTree[K, V]) ForEach(callback func(Pair[K, V]) bool) {
	d.b.Scan(func(item Pair[K, V]) bool {
		return callback(item)
	})
}

func (d *dictViaBTree[K, V]) Get(key K) (value V, ok bool) {
	item, ok := d.b.Get(NewPair(key, value))
	if !ok {
		return value, ok
	}

	return item.Second, true
}

func (d *dictViaBTree[K, V]) Keys() Collection[K] {
	return Map[Pair[K, V]](d, func(pair Pair[K, V]) K {
		return pair.First
	})
}

func (d *dictViaBTree[K, V]) Set(key K, value V) {
	d.b.Set(NewPair(key, value))
}

func (d *dictViaBTree[K, V]) Size() int {
	return d.b.Len()
}

func (d *dictViaBTree[K, V]) Values() Collection[V] {
	return Map[Pair[K, V]](d, func(pair Pair[K, V]) V {
		return pair.Second
	})
}

type setViaBTree[T any] struct {
	b    *btree.Generic[T]
	less func(T, T) bool
}

var _ interface{ Set[int] } = (*setViaBTree[int])(nil)

func newSetViaBTree[T any](less func(T, T) bool) *setViaBTree[T] {
	return &setViaBTree[T]{
		b:    btree.NewGeneric(less),
		less: less,
	}
}

func (s *setViaBTree[T]) Add(value T) {
	s.b.Set(value)
}

func (s *setViaBTree[T]) Clear() {
	s.b = btree.NewGeneric(s.less)
}

func (s *setViaBTree[T]) Delete(value T) {
	s.b.Delete(value)
}

func (s *setViaBTree[T]) ForEach(callback func(T) bool) {
	s.b.Scan(callback)
}

func (s *setViaBTree[T]) Has(value T) bool {
	_, ok := s.b.Get(value)
	return ok
}

func (s *setViaBTree[T]) Size() int {
	return s.b.Len()
}
