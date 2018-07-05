package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "encoding/json"
)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func populateJson(n int) string{
    mapD := map[string]int{"n": fibonacci(n), "n-1": fibonacci(n-1)}
    mapB, _ := json.Marshal(mapD)
    return (string(mapB))
}

func homePage(w http.ResponseWriter, r *http.Request) {
    numbers, ok := r.URL.Query()["number"]
    if !ok || len(numbers[0]) < 1 {
        log.Println("Url Param 'number' is missing")
        os.Exit(2)
    }
    number := numbers [0]
    log.Println("Url Param 'number' is: " + string(number))
    iNumber, err := strconv.Atoi(number)
    if err != nil {
       fmt.Println(err)
       os.Exit(2)
    }
    fmt.Fprintf(w, populateJson(iNumber))
}

func main() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":3004", nil))
}
