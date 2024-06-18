package adapter

type StripBytesIter iterSeq[byte]

type stripBytesIterImpl struct {
	bytes []byte
	state state
	utf8Parser utf8Parser
}

func (s *stripBytesIterImpl) Next() (byte, bool) {

}

func (s *stripBytesIterImpl) ForEach(yield func(byte) bool) {
	for {
		b, ok := s.Next()
		if !ok {
			break
		}
		if !yield(b) {
			break
		}
	}
}

func newStripBytesIter() StripBytesIter {
	s := &stripBytesIterImpl{}
	return s.ForEach
}