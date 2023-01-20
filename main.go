package main

import (
    "fmt"
    "net/http"
    "github.com/go-redis/redis"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    client := redis.NewClient(&redis.Options{
        Addr: "pinners-db-redis-ams3-dev-do-user-13352444-0.b.db.ondigitalocean.com:25061",
        Password: "AVNS_5jbeoQuN0HH-s4_jMFJ",
        DB: 0,
    })
    
    a := time.Now()
    
    err := client.Set(ctx, "key", a, 0).Err()
    if err != nil {
        panic(err)
    }
    
    fmt.Fprintf(w, "Hello Shit, %s!", r.URL.Path[1:])
}
