package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "io/ioutil"
    "encoding/json"
)

type Joke struct {
    ID string `json:"id"`
    JOKE string `json:"joke"`
    STATUS int `json:"status"`
}


type Response struct {
    StatusCode int               `json:"statusCode"`
    Headers    map[string]string `json:"headers"`
    Body       string            `json:"body"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {

    // cors and some req headers
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
    if err != nil {
        log.Fatalln(err)
    }
    req.Header.Set("Accept", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    joke := Joke{}
    jsonErr := json.Unmarshal(bytes,&joke)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }


    jokeJson, jsonError := json.Marshal(joke)

    fmt.Println(jsonError)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jokeJson)
}


// 1. Set up a router so we can browse to our local api in go  
// 2. set up a handler for an endpoint /dad-joke
// 3. set up integration to dad jokes api to fetch data joke data
// 4. parse response from dad jokes api and return only joke 
func main() {
    router := mux.NewRouter().StrictSlash(true)

    // http://localhost:8080/
    router.HandleFunc("/", homeLink)
    log.Fatal(http.ListenAndServe(":8080", router))
}
