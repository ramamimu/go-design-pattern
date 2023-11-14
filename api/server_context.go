package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Person struct {
	IsDefineError bool `json:"isDefineError"`
	TimeSleep     int  `json:"timeSleep"`
}

func NewPerson() *Person {
	return &Person{}
}

func (s *Server) handleContext(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	isDefineError := &Person{}
	err := json.NewDecoder(req.Body).Decode(isDefineError)
	if err != nil {
		http.Error(w, "failed to parse body", http.StatusBadRequest)
		return
	}
	fmt.Println("got ", isDefineError)
	w.Write(s.callContext(*isDefineError))
}

type Response struct {
	message []byte
	err     error
}

func (s *Server) callContext(p Person) []byte {
	if p.IsDefineError {
		return []byte("request success but automatically define error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	respCh := make(chan Response)

	go func() {
		val, err := s.fetchWithContext(p.TimeSleep)
		respCh <- Response{
			message: val,
			err:     err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return []byte("raise error time out")
		case resp := <-respCh:
			if resp.err != nil {
				return []byte(resp.err.Error())
			}
			return []byte(resp.message)
		}

	}
}

func (s *Server) fetchWithContext(timeSleep int) ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s/hello-more", s.listenAddr)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("error creating request")
		return []byte("hello error"), err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []byte("Error creating request"), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
		return []byte("Error making request"), err
	}
	defer resp.Body.Close()
	time.Sleep(time.Duration(timeSleep) * time.Second)

	return body, nil
}
