package tests

import (
	// helloworld "hello_world/modules"
	// "testing"

	"testing"

	"github.com/stretchr/testify/mock"
)

type HelloWorldMock struct {
	mock.Mock
}

func (m *HelloWorldMock) GetCatTags(url string) ([]byte, error) {
	args := m.Called(url)
	return []byte(args.String(0)), args.Error(1)
}

func TestMockHelloWorld(t *testing.T) {
	argMock := HelloWorldMock{}
	argMock.On("GetCatTags", "https://cataas.com/api/tags").Return("hello world", nil)
}

// // func (m *HelloWorldMock) CatTags() ([]byte, error) {
// // 	args := m
// // }

// // func  (mock *HelloWorldMock) GetCatTags() ([]byte, error) {
// // 	args := mock.Called()
// // 	result := args.Get(0)
// // 	return
// // }

// func TestHelloWorld(t *testing.T) {
// 	// result := string(helloworld.OrdinaryHelloWorld())

// 	// mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.WriteHeader(http.StatusOK)
// 	// 	w.Write([]byte(`hello ww wolrd`))
// 	// }))

// 	t.Run("test mock hello world", func(t *testing.T) {
// 		mockFunction := new(HelloWorldMock)
// 		mockFunction.On("GetCatTags").Return([]byte("hello ww world"), nil)
// 		// result := string(helloworld.OrdinaryHelloWorld())
// 		result, _ := helloworld.GetCatTags()
// 		// defer mockServer.Close()
// 		// testService = mockFunctio

// 		// resp := helloworld.GetCatTags()
// 		if "hello ww world" != string(result) {
// 			t.Errorf("Expected 'hello world', got %s", result)
// 		}
// 	})

// 	// if "hello ww world" != result {
// 	// 	t.Errorf("Expected 'hello world', got %s", result)
// 	// }
// }
