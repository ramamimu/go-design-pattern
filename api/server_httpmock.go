package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerQuery(s *string, p []string) {
	if len(p) > 0 {
		for _, item := range p {
			*s = fmt.Sprintf("%s %s", *s, item)
		}
	}
}

// for better understanding how to mock HTTP
// including params, headers, method, body
func (s *Server) HandleMockHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if h := r.Header.Get("Authentication"); h == "" {
		http.Error(w, "invalid header", http.StatusBadRequest)
		return
	} else {
		fmt.Printf("header: %s\n", h)
	}

	person := &Person{}
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		return
	}

	returnStatement := fmt.Sprintf("hello %t %d", person.IsDefineError, person.TimeSleep)

	// the query form is map
	// for every map return an array
	p1Param := r.URL.Query()["p1"]
	p2Param := r.URL.Query()["p2"]

	handlerQuery(&returnStatement, p1Param)
	handlerQuery(&returnStatement, p2Param)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(returnStatement))
}
