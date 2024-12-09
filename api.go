
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	//"os"
	//"time"
)

//CREATES A UNIQUE 40 digit ID, sperated by dashes, for every request
func generateRequestID() string {
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ID := ""
	for i:=0; i < 4; i++ {
		PT := ""
		for j:=0; j < 10; j++ {
			rand_num := rand.Intn(62)
			PT += string(letters[rand_num])
		}
		if i!=3 {ID += PT + "-"} else {ID += PT}
	}
	fmt.Println(ID)
	return ID
}

func spellcheck(w http.ResponseWriter, r *http.Request, requestID string) {

  word :=  r.URL.Query().Get("word")
  if word == "" {
    http.Error(w, `{"error": "Invalid or missing 'word' parameter"}`, http.StatusBadRequest)
  } else {
		
		//line specific to directory
		cmd := exec.Command("/usr/api.itsbensmall.com/checker.exe", word)
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error", err)
			return
		}

		words_response := []string{}
		scanner := bufio.NewScanner(&out)
		for scanner.Scan() {
			words_response = append(words_response, scanner.Text())
		}

		// Convert response to JSON and write it
		if err := json.NewEncoder(w).Encode(words_response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	requestID := generateRequestID()

  switch val := r.URL.Query().Get("val"); val {
  case "spellcheck":
    spellcheck(w, r, requestID)
  default:
    http.Error(w, `{"error": "Invalid or missing 'val' parameter"}`, http.StatusBadRequest)
  }

}

func main() {
	http.HandleFunc("/", handler)

	port := ":8090"
	println("Server running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
