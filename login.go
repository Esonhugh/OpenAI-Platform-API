package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
)

func (u *UserClient) Logout() error {
	_, err := u.client.Get(auth0LogoutUrl)
	return err
}

func (u *UserClient) LoginWithAuth0(Username, Password string) error {
	// hard refresh cookies
	resp, _ := u.client.Get(auth0LogoutUrl)
	defer resp.Body.Close()

	// get authorized url
	authorizedUrl, statusCode, err := u.GetAuthorizedUrl("")
	if err != nil {
		return errors.Join(
			errors.New(fmt.Sprintf("HTTP url: %v got HTTP Status: %v", authorizedUrl, statusCode)),
			err,
		)
	}

	// get state
	state, _, _ := u.GetState(authorizedUrl)

	// check username
	statusCode, err = u.CheckUsername(state, Username)
	if err != nil {
		return errors.Join(
			errors.New(fmt.Sprintf("CheckUsername, got HTTP Status: %v", statusCode)),
			err,
		)
	}

	// check password
	code, statusCode, err := u.CheckPassword(state, Username, Password)
	if err != nil {
		return errors.Join(
			errors.New(fmt.Sprintf("CheckPassword, got HTTP Status: %v", statusCode)),
			err,
		)
	}

	// get access token
	accessToken, statusCode, err := u.GetAccessToken(code)
	if err != nil {
		return errors.Join(
			errors.New(fmt.Sprintf("GetAccessToken, got HTTP Status: %v", statusCode)),
			err,
		)
	}

	// get session key
	var getAccessTokenResponse GetAccessTokenResponse
	json.Unmarshal([]byte(accessToken), &getAccessTokenResponse)
	data, statusCode, err := u.DashboardLogin(getAccessTokenResponse.AccessToken)
	if err != nil {
		return errors.Join(
			errors.New(
				fmt.Sprintf("Try Login dashboard, got HTTP Status: %v", statusCode),
			),
			err,
		)
	}
	defer resp.Body.Close()
	if statusCode != http.StatusOK {
		return errors.Join(
			errors.New(fmt.Sprintf("Try login dashboard.got non-200 HTTP Status: %v", statusCode)),
			err,
		)
	}

	var getHealthCheckResponse DashboardLoginResponse
	json.Unmarshal([]byte(data), &getHealthCheckResponse)

	u.accessToken = getAccessTokenResponse.AccessToken
	u.sessionKey = getHealthCheckResponse.User.Session.SensitiveID
	return nil
}

func (u *UserClient) LoginWithAccessToken() (DashboardLoginResponse, error) {
	return u.DashboardOnBoarding()
}

func (u *UserClient) DashboardOnBoarding() (DashboardLoginResponse, error) {
	if u.AccessToken() == "" {
		return DashboardLoginResponse{}, errors.New("dashboard Onboarding but accessToken is empty, you need re-login")
	}
	data, statusCode, err := u.DashboardLogin(u.AccessToken())
	if err != nil {
		return DashboardLoginResponse{}, errors.Join(
			errors.New(
				fmt.Sprintf("Try Login dashboard, got HTTP Status: %v", statusCode),
			),
			err,
		)
	}
	if statusCode != http.StatusOK {
		return DashboardLoginResponse{}, errors.Join(
			errors.New(fmt.Sprintf("Try login dashboard.got non-200 HTTP Status: %v", statusCode)),
			err,
		)
	}

	var getHealthCheckResponse DashboardLoginResponse
	err = json.Unmarshal([]byte(data), &getHealthCheckResponse)
	if err != nil {
		return DashboardLoginResponse{}, errors.Join(
			errors.New("GetHealthCheck error, bad json data, data: "+data),
			err)
	}

	u.sessionKey = getHealthCheckResponse.User.Session.SensitiveID
	return getHealthCheckResponse, nil
}
