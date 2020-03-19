package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title	string `json:"title"`
	Author 	string `json:"author"`
	ISBN 	string `json:"isbn"`
}

func (b Book) toJSON() []byte {
	bookJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return bookJSON
}

func fromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	{
		Title:  "Title 1",
		Author: "Author 1",
		ISBN:   "ISBN 1",
	},
	{
		Title:  "Title 2",
		Author: "Author 2",
		ISBN:   "ISBN 2",
	},
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}