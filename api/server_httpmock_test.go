package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	hw "hello_world/api"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMockHTTP(t *testing.T) {
	wRecorder := httptest.NewRecorder()

	person := hw.NewPerson()
	person.IsDefineError = false
	person.TimeSleep = 10

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(person); err != nil {
		t.Errorf("Error convert struct to body io %s", err)
	}

	rRequest := httptest.NewRequest(http.MethodPost, "http://localhost:9090?p1=hai&p1=hei&p2=world", &buf)
	rRequest.Header.Add("Authentication", "hola")

	s := hw.Server{}
	s.HandleMockHTTP(wRecorder, rRequest)

	if wRecorder.Result().StatusCode == http.StatusBadRequest {
		t.Error("the param is empty")
	}

	if wRecorder.Result().StatusCode == http.StatusMethodNotAllowed {
		t.Error("error method")
	}

	b, err := io.ReadAll(wRecorder.Result().Body)
	if err != nil {
		t.Error(err.Error())
	}

	expect := fmt.Sprintf("hello %t %d hai hei world", person.IsDefineError, person.TimeSleep)

	bString := string(b)
	if expect != string(b) {
		t.Errorf("expected: %s != %s", expect, bString)
	}
}
