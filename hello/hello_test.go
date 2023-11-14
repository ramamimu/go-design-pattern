package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockText struct {
	mock.Mock
}

func (m *mockText) GetHelloWorld() string {
	args := m.Called()
	return args.String(0)
}

func (m *mockText) GetHeiWorld() string {
	args := m.Called()
	return args.String(0)
}

func TestHelloWorld(t *testing.T) {
	// mocker := new(mockText)
	mocker := &mockText{}
	mocker.On("GetHelloWorld").Return("hello bang")

	textGenerator := NewText(mocker)
	result := textGenerator.service.GetHelloWorld()
	var expected string = "hello bang"
	assert.Equal(t, expected, result)
}
