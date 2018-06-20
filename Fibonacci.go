package main

import (
    "fmt"
    "log"
    "strconv"
    "net/http"
)

func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Fibonacci: " + strconv.Itoa(Fibonacci(8)))
}

func main() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
