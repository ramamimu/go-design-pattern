package mockingjay_test

import (
	"hello_world/mockingjay"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockerMock struct {
	mock.Mock
}

func (m *MockerMock) RandNumb() int {
	args := m.Called()
	return args.Int(0)
}

func TestMockNestify(t *testing.T) {
	randomer := new(MockerMock)
	randomer.On("RandNumb").Return(12)

	temp := mockingjay.NewNumber(randomer)
	expect := temp.AddNumb()
	assert.Equal(t, 21, expect)
}

func TestAddHelloJay(t *testing.T) {
	randomer := mockingjay.NewRandom()
	realm := mockingjay.NewNumber(randomer)

	result := realm.AddNumb()
	assert.Equal(t, 51, result)
}

func TestAddHelloJayWithMock(t *testing.T) {
	randomer := mockingjay.NewMockRandom()
	realm := mockingjay.NewNumber(randomer)

	result := realm.AddNumb()
	assert.Equal(t, 429, result)
}
