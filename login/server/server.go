package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nerzal/gocloak/v11"
)

var (
	realm        = os.Getenv("REALM")
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	keycloakUrl  = os.Getenv("KEYCLOAK_URL")
)

func main() {
	http.HandleFunc("/", loginHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TODO: only allow from top-level domain

	err := r.ParseForm()
	if err != nil {
		log.Fatalf("parse form error: %v", err)
	}
	login := r.FormValue("login")
	password := r.FormValue("password")
	client := gocloak.NewClient(keycloakUrl)
	restyClient := client.RestyClient()
	restyClient.SetDebug(false)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	ctx := context.Background()
	jwt, err := client.Login(ctx, clientID, clientSecret, realm, login, password)
	if err != nil {
		log.Printf("could not login at keycloak: %v", err)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "could not login at keycloak: %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, jwt.AccessToken)
}
