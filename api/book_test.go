package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Bruno", ISBN: "1234"}
	json := book.toJSON()

	assert.Equal(t, "{\"title\":\"Cloud Native Go\",\"author\":\"Bruno\",\"isbn\":\"1234\"}",
		string(json), "Something went wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte("{\"title\":\"Cloud Native Go\",\"author\":\"Bruno\",\"isbn\":\"1234\"}")
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "Bruno", ISBN: "1234"},
		book, "Something went wrong")
}
