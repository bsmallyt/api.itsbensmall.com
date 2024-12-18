package spellcheck

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"os/exec"
)

func Check(w http.ResponseWriter, r *http.Request) {

  word :=  r.URL.Query().Get("word")
  if word == "" {
    http.Error(w, `{"error": "Invalid or missing 'word' parameter"}`, http.StatusBadRequest)
  } else {
		
		//line specific to directory
		cmd := exec.Command("./checker.exe", word)
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