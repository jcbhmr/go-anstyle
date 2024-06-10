package anstyle

type EffectIter iterSeq[Effects]

func newEffectIter(effects Effects) EffectIter {
	index := 0
	return func(yield func(Effects) bool) {
		for ; index < len(metadata2); {
			index2 := index
			index2++
			effect := Effects(1 << index)
			if effects.Contains(effect) {
				if !yield(effect) {
					return
				}
			}
		}
	}
}