package main

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {

	// set variables
	serverUrl := "http://localhost:4001"
	ctx := context.Background()

	// set credentials - usually gotten from a secret store
	creds := clientcredentials.Config{
		ClientID:     "paul",
		ClientSecret: "dirac",
		TokenURL:     serverUrl + "/oauth/token",
		Scopes:       []string{"admin"},
	}

	// get token
	token, err := creds.Token(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("token:", token.AccessToken)

	// validate token

	// use client from creds to keep auth headers
	resp, err := creds.Client(ctx).Post(serverUrl+"/validate", "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// decode validate response
	response := map[string]string{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(response["error_description"])
	}

	fmt.Println("token response:", response["data"])
}
