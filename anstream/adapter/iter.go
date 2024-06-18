package adapter

type iterSeq[V any] func(yield func(V) bool)
type iterSeq2[K, V any] func(yield func(K, V) bool)
