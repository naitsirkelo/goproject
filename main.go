package main

import(
  "fmt"
  "os"
  "log"
  "net/http"
)

type Movie struct {
  Title    string `json="Title"`
  Genre    string `json="Genre"`
  Language string `json="Language"`
  Country  string `json="Country"`
  Runtime  string `json="Runtime"`
  Director string `json="Director"`
  Released string `json="Released"`
  // Poster   string `json="Poster"`
}

// - - - - - - - - - -  Port  - - - - - - - - - - - - - -

func GetPort() string {
	 	var port = os.Getenv("PORT")
 				// Port sets to :8080 as a default
 		if (port == "") {
 			port = "8080"
			fmt.Println("No PORT variable detected, defaulting to " + port)
 		}
 		return (":" + port)
}

// - - - - - - - - - -  Main  - - - - - - - - - - - - -

func main() {
  fmt.Println("Hello World!")

  http.HandleFunc("/", Handler)
  http.HandleFunc("/fymihelp", HelpHandler)

  err := http.ListenAndServe(GetPort(), nil)
  if err != nil {
      log.Fatal("ListenAndServe Error: ", err)
  }
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - -
