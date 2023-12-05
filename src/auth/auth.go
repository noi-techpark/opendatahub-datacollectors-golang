package auth

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        uint64 `json:"expires_in"`
	NotBeforePolicy  uint64 `json:"not-before-policy"`
	RefreshExpiresIn uint64 `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	Scope            string
}

var tokenUri string = os.Getenv("OAUTH_TOKEN_URI")
var clientId string = os.Getenv("OAUTH_CLIENT_ID")
var clientSecret string = os.Getenv("OAUTH_CLIENT_SECRET")
var scope string = os.Getenv("OAUTH_CLIENT_SCOPE")

var token Token

func GetToken() string {
	if len(token.AccessToken) > 0 {
		return token.AccessToken
	}

	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("grant_type", `client_credentials`)
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

	return token.AccessToken
}
