package platform

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strings"

	http "github.com/bogdanfinn/fhttp"
)

//goland:noinspection GoUnhandledErrorResult,GoErrorStringFormat,GoUnusedParameter
func (u *UserClient) GetAuthorizedUrl(csrfToken string) (string, int, error) {
	urlParams := url.Values{
		"client_id":     {platformAuthClientID},
		"audience":      {platformAuthAudience},
		"redirect_uri":  {platformAuthRedirectURL},
		"scope":         {platformAuthScope},
		"response_type": {platformAuthResponseType},
	}
	req, _ := http.NewRequest(http.MethodGet, platformAuth0Url+urlParams.Encode(), nil)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	u.lastResponse = resp
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(GetAuthorizedUrlErrorMessage)
	}

	return resp.Request.URL.String(), http.StatusOK, nil
}

func (u *UserClient) GetState(authorizedUrl string) (string, int, error) {
	split := strings.Split(authorizedUrl, "=")
	return split[1], http.StatusOK, nil
}

//goland:noinspection GoUnhandledErrorResult,GoErrorStringFormat
func (u *UserClient) CheckUsername(state string, username string) (int, error) {
	formParams := url.Values{
		"state":                       {state},
		"username":                    {username},
		"js-available":                {"true"},
		"webauthn-available":          {"true"},
		"is-brave":                    {"false"},
		"webauthn-platform-available": {"false"},
		"action":                      {"default"},
	}
	req, err := http.NewRequest(http.MethodPost, LoginUsernameUrl+state, strings.NewReader(formParams.Encode()))
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	u.lastResponse = resp

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, errors.New(EmailInvalidErrorMessage)
	}

	return http.StatusOK, nil
}

//goland:noinspection GoUnhandledErrorResult,GoErrorStringFormat
func (u *UserClient) CheckPassword(state string, username string, password string) (string, int, error) {
	formParams := url.Values{
		"state":    {state},
		"username": {username},
		"password": {password},
		"action":   {"default"},
	}
	req, err := http.NewRequest(http.MethodPost, LoginPasswordUrl+state, strings.NewReader(formParams.Encode()))
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	u.lastResponse = resp

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(EmailOrPasswordInvalidErrorMessage)
	}

	return resp.Request.URL.Query().Get("code"), http.StatusOK, nil
}

//goland:noinspection GoUnhandledErrorResult,GoErrorStringFormat
func (u *UserClient) GetAccessToken(code string) (string, int, error) {
	jsonBytes, _ := json.Marshal(GetAccessTokenRequest{
		ClientID:    platformAuthClientID,
		Code:        code,
		GrantType:   platformAuthGrantType,
		RedirectURI: platformAuthRedirectURL,
	})
	req, err := http.NewRequest(http.MethodPost, getTokenUrl, strings.NewReader(string(jsonBytes)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	u.lastResponse = resp

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(GetAccessTokenErrorMessage)
	}

	data, _ := io.ReadAll(resp.Body)
	return string(data), http.StatusOK, nil
}

type DashboardLoginResponse struct {
	Object  string        `json:"object"`
	User    User          `json:"user"`
	Invites []interface{} `json:"invites"`
}

//goland:noinspection GoUnhandledErrorResult,GoErrorStringFormat
func (u *UserClient) DashboardLogin(AccessToken string) (string, int, error) {
	req, err := http.NewRequest(http.MethodPost, dashboardLoginUrl, strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	u.lastResponse = resp
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(GetAccessTokenErrorMessage)
	}

	data, _ := io.ReadAll(resp.Body)
	return string(data), http.StatusOK, nil
}
