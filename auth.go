package scrolls

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// asks the Mojang authentication server for an auth token to connect to the game server
func authenticate(email, password string) (mFirstConnect, error) {
	var m mFirstConnect

	// first we have to request an authorization token
	var authRequest bytes.Buffer
	json.NewEncoder(&authRequest).Encode(req{
		"agent": req{
			"name":    "Scrolls",
			"version": 1,
		},
		"username": email,
		"password": password,
	})

	resp, err := http.Post("https://authserver.mojang.com/authenticate", "application/json", &authRequest)
	if err != nil {
		return m, err
	}

	// that is copied into our FirstConnect message
	err = json.NewDecoder(resp.Body).Decode(&m.AccessToken)
	resp.Body.Close()

	return m, err
}
