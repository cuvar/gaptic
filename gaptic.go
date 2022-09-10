package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/icrowley/fake"
)
type User struct {
	Name string
	Email string
	Age string
}

type Error struct {
	Error string
	Message string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Please use /v1/user")
	})

	http.HandleFunc("/v1/users", UserHandler)

	port := ":8001"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}

func UserHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	const errorRes = "{\"error\": \"Internal Server Error\", \"message\": \"An error occured while processing your request\"}"

	amount := r.URL.Query().Get("amount")
	if amount == "" {
		amount = "1"
	}
	
	amountNumber, _ := strconv.Atoi(amount);

	users := createUsers(amountNumber)
	res, err := json.Marshal(users)
	
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errorRes)
		return
	} 

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(res))
}


func createUsers(amount int) *[]User {
	rand.Seed(time.Now().UnixNano())

	users := make([]User, amount)
	for i := 0; i < amount; i++ {
		users[i] = User{
			Name: fake.FullName(),
			Email: fake.EmailAddress(),
			Age: strconv.Itoa(randInt(5, 10)),
		}
	}
	return &users
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}