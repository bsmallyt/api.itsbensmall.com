package main

import (
	"api/spellcheck"
	"api/reload"

	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//requestID := generateRequestID()

  switch val := r.URL.Query().Get("val"); val {
  case "spellcheck":
		os.Chdir("/usr/api.itsbensmall.com/spellcheck/cpp")
    spellcheck.Check(w, r)
		os.Chdir("/usr/api.itsbensmall.com/main")
	case "reload":
		os.Chdir("/usr/api.itsbensmall.com/reload")
		reload.Fitsbensmall()
		os.Chdir("/usr/api.itsbensmall.com/main")
  default:
    http.Error(w, `{"error": "Invalid or missing 'val' parameter"}`, http.StatusBadRequest)
  }

}

func main() {
	http.HandleFunc("/", handler)

	port := ":90"
	println("Server running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}