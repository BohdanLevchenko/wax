package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "strconv"
)

type MappingEntry struct {
  Source string
  Target string
}

type Configuration struct {
  Mappings []MappingEntry
}

var mappings map[string]string
var port *int
var code *int

func redirect(w http.ResponseWriter, r *http.Request) {
  target, found := mappings[r.URL.Path]
  if found {
    fmt.Printf("Redirect %s => %s\n", r.URL.Path, target)
    http.Redirect(w, r, target, *code)
  } else {
    fmt.Printf("Redirect %s => 404\n", r.URL.Path)
    w.WriteHeader(http.StatusNotFound)
    io.WriteString(w, "404 Not Found\n")
  }
}

func main() {
  port = flag.Int("port", 8001, "http port")
  code = flag.Int("code", 301, "Redirect status code")
  flag.Parse()

  file, _ := os.Open("wax.json")
  decoder := json.NewDecoder(file)
  config := Configuration{}
  err := decoder.Decode(&config)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  mappings = make(map[string]string)
  for _, me := range config.Mappings {
    mappings[me.Source] = me.Target
  }

  http.HandleFunc("/", redirect)
  fmt.Printf("Listening on http://0.0.0.0:%d\nHit Ctrl+C to exit\n", *port)
  fmt.Printf("Registered mappings: %v\n", mappings)
  err = http.ListenAndServe(":"+strconv.Itoa(*port), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
