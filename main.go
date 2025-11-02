package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type user struct {
	Name string `json:"name"`
}

var userCache = make(map[int]user)
var cacheMutex sync.RWMutex

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /Users", createUser)
	mux.HandleFunc("GET /Users/{id}", getUser)
	mux.HandleFunc("DELETE /Users/{id}", deleteUser)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var User user
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if User.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
	}

	cacheMutex.Lock()
	userCache[len(userCache)+1] = User
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
