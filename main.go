package main

import (
	"fmt"
	"log"
	"net/http"

        "bitbucket.org/ericroku/sessions"
)
var (
	// Key for encrypting the session cookie (use a secure random key in production)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/set", setSessionHandler)
	http.HandleFunc("/get", getSessionHandler)

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// homeHandler displays a welcome message
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Use /set to set a session and /get to retrieve it.")
}

// setSessionHandler sets a session value
func setSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. This will create a new session if one doesn't exist.
	session, _ := store.Get(r, "example-session")

	// Set some session values
	session.Values["username"] = "gopher"
	session.Values["email"] = "gopher@example.com"

	// Save the session
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Session values set!")
}

// getSessionHandler retrieves session values
func getSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the session
	session, _ := store.Get(r, "example-session")

	// Retrieve session values
	username, ok := session.Values["username"].(string)
	if !ok {
		username = "unknown"
	}

	email, ok := session.Values["email"].(string)
	if !ok {
		email = "unknown"
	}

	fmt.Fprintf(w, "Session values:\nUsername: %s\nEmail: %s\n", username, email)
}

