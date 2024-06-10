package anstyle

type stringerString string

func (s stringerString) String() string {
	return string(s)
}

type iterSeq[V any] func(yield func(V) bool)
type iterSeq2[K, V any] func(yield func(K, V) bool)

type option[T any] struct {
	ok    bool
	value T
}

func optionSome[T any](v T) option[T] {
	return option[T]{ok: true, value: v}
}
func optionNone[T any]() option[T] {
	return option[T]{ok: false}
}

func optionFromPtr[T any](v *T) option[T] {
	if v == nil {
		return optionNone[T]()
	} else {
		return optionSome[T](*v)
	}
}
func optionFromOk[T any](v T, ok bool) option[T] {
	if ok {
		return optionSome[T](v)
	} else {
		return optionNone[T]()
	}
}

func (o option[T]) ToPtr() *T {
	if o.ok {
		return &o.value
	} else {
		return nil
	}
}
func (o option[T]) ToOk() (T, bool) {
	return o.value, o.ok
}
