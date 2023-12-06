package auth

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	NotBeforePolicy  int64  `json:"not-before-policy"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string
}

var tokenUri string = os.Getenv("OAUTH_TOKEN_URI")
var clientId string = os.Getenv("OAUTH_CLIENT_ID")
var clientSecret string = os.Getenv("OAUTH_CLIENT_SECRET")

var token Token

func GetToken() string {
	ts := time.Now().Unix()

	if len(token.AccessToken) > 0 || ts > token.RefreshExpiresIn {
		// if no token is available or refreshToken is expired, get new token
		newToken()
	} else if ts > token.ExpiresIn {
		// if no token is expired, refresh it
		refreshToken()
	}

	return token.AccessToken
}

func refreshToken() {
	slog.Info("Refreshing token...")

	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("grant_type", `refresh_token`)
	params.Add("refresh_token", token.RefreshToken)

	authRequest(params)

	slog.Info("Refreshing token done.")
}

func newToken() {
	slog.Info("Getting new token...")
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("grant_type", `client_credentials`)

	authRequest(params)

	slog.Info("Getting new token done.")
}

func authRequest(params url.Values) {
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", tokenUri, body)
	if err != nil {
		slog.Error("error", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("error", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error("error", err)
		}

		err = json.Unmarshal(bodyBytes, &token)
		if err != nil {
			slog.Error("error", err)
		}
	}
}
