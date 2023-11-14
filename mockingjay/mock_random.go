package mockingjay

type mockRandom struct{}

func NewMockRandom() *mockRandom {
	return &mockRandom{}
}

func (r mockRandom) RandNumb() int {
	return 420
}
