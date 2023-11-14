package modules

func OrdinaryHelloWorld(url string) []byte {
	cat := Cat{}
	res, err := cat.GetCatTags(url)
	if err != nil {
		return []byte(err.Error())
	}
	return res
}

func JsonHelloWorld() []byte {
	var jsonData = []byte(`{"name": "morpheus","job": "leader"}`)
	return jsonData
}

type message struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func JsonHelloWorldStruct() []message {
	var jsonData = []message{
		{Error: false, Message: "hello world"},
		{Error: true, Message: "another world"},
	}

	return jsonData
}
