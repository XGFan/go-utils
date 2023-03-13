package utils

var EMPTY_STRUCT = struct{}{}

func SliceToChan[T any](list []T) chan T {
	c := make(chan T, 0)
	go func() {
		for _, s := range list {
			c <- s
		}
		close(c)
	}()
	return c
}

type Set[T comparable] interface {
	Contains(item T) bool
	Add(item T)
	Remote(item T)
	Size() int
	Values() []T
}

func NewSet[T comparable]() Set[T] {
	return &DefaultSet[T]{
		m: make(map[T]struct{}, 0),
	}
}

func NewSetWithSlice[T comparable](slice []T) Set[T] {
	set := &DefaultSet[T]{
		m: make(map[T]struct{}, 0),
	}
	for _, t := range slice {
		set.Add(t)
	}
	return set
}

type DefaultSet[T comparable] struct {
	m map[T]struct{}
}

func (d DefaultSet[T]) Contains(item T) bool {
	_, exist := d.m[item]
	return exist
}

func (d DefaultSet[T]) Add(item T) {
	d.m[item] = EMPTY_STRUCT
}

func (d DefaultSet[T]) Remote(item T) {
	delete(d.m, item)
}

func (d DefaultSet[T]) Size() int {
	return len(d.m)
}

func (d DefaultSet[T]) Values() []T {
	keys := make([]T, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func AppendAll[T any](a, b []T) []T {
	for _, t := range b {
		a = append(a, t)
	}
	return a
}

func Distinct[T comparable](src []T) []T {
	return NewSetWithSlice(src).Values()
}
