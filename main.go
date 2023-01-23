package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

var ctx = context.Background()

func HelloServer(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis-10835.c55.eu-central-1-1.ec2.cloud.redislabs.com:10835",
		//Addr: "localhost:6379",
		Username: "default",
		Password: "BWz6EQvY7TOrYGYXqWzmqDb1yG5eYUup",
		DB:       0,
	})

	a := time.Now()
	fmt.Println(a)

	id := uuid.New().String()

	err := client.Set(ctx, id, a, 0).Err()
	if err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Fprintf(w, "Hello Shit, %s!", r.URL.Path[1:])

}

/*

	fmt.Println("Ping")
	client.Ping(ctx)
	fmt.Println("Pong")

	val2, err := client.Get(ctx, "key2").Result()

	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("key2", val2)
	}

*/
