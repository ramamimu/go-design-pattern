package api

import (
	"encoding/json"
	helloworld "hello_world/modules"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listAddr string) *Server {
	return &Server{
		listenAddr: listAddr,
	}
}

func (s *Server) Start() error {
	// r := mux.NewRouter()
	http.HandleFunc("/hello-world", s.handleHelloWorld)
	// http.HandleFunc("/hello/{firstParam}/{secondParam}", s.handleGetJson)
	http.HandleFunc("/hello", s.handleGetJson)
	http.HandleFunc("/hello-more", s.handleGetJsonStruct)
	http.HandleFunc("/trigger-context", s.handleContext)
	http.HandleFunc("/test-mock-http", s.HandleMockHTTP)
	// http.HandleFunc("/call-context", s.handleContext)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleHelloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write(helloworld.OrdinaryHelloWorld("https://cataas.com/api/tags"))
}

func (s *Server) handleGetJson(w http.ResponseWriter, req *http.Request) {
	print(req.URL.String())
	w.Header().Set("content-type", "application/json")
	w.Write(helloworld.JsonHelloWorld())
}

func (s *Server) handleGetJsonStruct(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	data := helloworld.JsonHelloWorldStruct()
	jsonData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
