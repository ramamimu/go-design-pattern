package mockingjay

type random struct{}

func NewRandom() *random {
	return &random{}
}

func (r *random) RandNumb() int {
	return 42
}
