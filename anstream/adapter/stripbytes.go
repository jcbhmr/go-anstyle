package adapter

type StripBytes struct {
	state state
	utf8Parser utf8Parser
}

func NewStripBytes() *StripBytes {
	return &StripBytes{}
}

func (s *StripBytes) StripNext(bytes []byte) StripBytesIter {}
