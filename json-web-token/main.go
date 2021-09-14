package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

// var mySigningKey = os.get("MY_JWT_TOKEN")
var mySigningKey = []byte("my-secret")

func homePage(w http.ResponseWriter, r *http.Request)  {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "ashikur"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests()  {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main()  {
	fmt.Println("my simple client")

	handleRequests()
}
