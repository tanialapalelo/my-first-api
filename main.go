package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TodoItem struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

func main() {
	var todos = make([]TodoItem, 0)
	var nextId = 1

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todo", func(writer http.ResponseWriter, request *http.Request) {
		b, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
		}
		_, err = writer.Write(b)
		if err != nil {
			log.Println(err)
		}
	})
	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		t.Id = nextId
		nextId++
		todos = append(todos, t)
		writer.WriteHeader(http.StatusCreated)
		return
	})
	mux.HandleFunc("DELETE /todo/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.PathValue("id")
		id, err := strconv.Atoi(idStr) // ASCII to integer
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		for i, t := range todos {
			if t.Id == id {
				// IMPORTANT NOTE:
				//  todos[:i] is everything before the match, todos[i+1:] is everything after it, and
				//  append glues them together — effectively squeezing the matched element out. It works in-place (reuses the same backing array), which is the idiomatic Go way to delete from a
				//  slice since there's no .remove() method like in other languages.
				todos = append(todos[:i], todos[i+1:]...)
				writer.WriteHeader(http.StatusNoContent)
				return
			}
		}
		// NOTE:
		//  if the loop finishes without finding a match, the handler just falls through and returns without writing any status. In Go's net/http, that means it silently
		//  sends 200 OK by default, even though nothing was actually deleted. That's misleading for a client
		writer.WriteHeader(http.StatusNotFound)
	})
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
