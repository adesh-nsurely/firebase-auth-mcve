package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	firebase "firebase.google.com/go/v4"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// Set the number of users to create.
	numUsers := 50

	// Initialize Firebase SDK.
	opt := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v\n", err)
		return
	}

	// Obtain a reference to the auth service.
	client, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error getting Auth client: %v\n", err)
		return
	}

	// Create the users parallelly.
	var wg sync.WaitGroup
	for i := 0; i < numUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Generate a random email ID.
			randomEmail := fmt.Sprintf("%s@nsurely.com", generateRandomString(15))

			// Set the params for creating the user.
			params := (&firebaseAuth.UserToCreate{}).Email(randomEmail)

			// Create the new user.
			u, err := client.CreateUser(context.Background(), params)
			if err != nil {
				fmt.Printf("error creating user: %v\n", err)
			} else {

				// User creation successful.

				// Set the time when the user was created.
				time := time.Now().Format("2006-01-02 15:04:05")

				// Convert the default claims to JSON for printing.
				claims, err := json.Marshal(u.CustomClaims)
				if err != nil {
					fmt.Printf("error marshalling user: %v\n", err)
				}
				fmt.Printf("%v: Successfully created user %v and default claims %s\n", time, u.Email, claims)
			}

			// set custom claims on the user
			claims := map[string]interface{}{
				"orgGroup": generateRandomString(20),
				"orgs":     []string{generateRandomString(20), generateRandomString(20)},
			}

			err = client.SetCustomUserClaims(context.Background(), u.UID, claims)
			if err != nil {
				fmt.Printf("error setting custom claims: %v\n", err)
			}

			// wait for 5 seconds
			time.Sleep(5 * time.Second)

			// fetch the user again and print the claims
			u, err = client.GetUser(context.Background(), u.UID)
			if err != nil {
				fmt.Printf("error fetching user: %v\n", err)
			} else {
				time := time.Now().Format("2006-01-02 15:04:05")

				// Convert the custom claims to JSON for printing.
				claims, err := json.Marshal(u.CustomClaims)
				if err != nil {
					fmt.Printf("error marshalling claims: %v\n", err)
				}
				fmt.Printf("%v: Successfully fetched user %v custom claims: %s\n", time, u.Email, claims)
			}
		}()
	}
	wg.Wait()
}

// Generate a random string
func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
