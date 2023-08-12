package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"io"
	"net/url"
	"strings"
)

type GetSecretKeysResponse struct {
	Object string `json:"object"`
	Data   []Key  `json:"data"`
}

func (u *UserClient) GetSecretKeys() (GetSecretKeysResponse, error) {
	if u.SessionKey() == "" {
		return GetSecretKeysResponse{}, errors.New("GetSecretKeys with no SessionKey is Defined")
	}
	formParams := url.Values{}
	req, err := http.NewRequest(http.MethodGet, PlatformApiUrlPrefix+"/dashboard/user/api_keys", strings.NewReader(formParams.Encode()))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set(AuthorizationHeader, "Bearer "+u.SessionKey())
	resp, err := u.client.Do(req)
	if err != nil {
		return GetSecretKeysResponse{}, errors.Join(
			errors.New("GetSecretKeys error"),
			err,
		)
	}
	if resp.StatusCode != http.StatusOK {
		return GetSecretKeysResponse{}, errors.Join(
			errors.New(fmt.Sprintf("GetSecretKeys found non 200 response, StatusCode: %v", resp.StatusCode)),
			err)
	}
	data, _ := io.ReadAll(resp.Body)

	var response GetSecretKeysResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return GetSecretKeysResponse{}, errors.Join(
			errors.New("GetSecretKey Unmarshal error"),
			err)
	}

	return response, nil
}

type CreateSecretKeyResponse struct {
	Result string `json:"result"`
	Key    Key    `json:"key"`
}

type ActionMap struct {
	Action      string `json:"action"`
	Name        string `json:"name,omitempty"`
	CreateAt    int    `json:"create_at,omitempty"`
	RedactedKey string `json:"redacted_key,omitempty"`
}

func (u *UserClient) CreateSecretKey(name string) (CreateSecretKeyResponse, error) {
	if u.SessionKey() == "" {
		return CreateSecretKeyResponse{}, errors.New("CreateSecretKey with no SessionKey is Defined")
	}
	form := ActionMap{
		Action: "create",
		Name:   name,
	}
	bytedata, _ := json.Marshal(form)
	req, err := http.NewRequest(http.MethodPost, PlatformApiUrlPrefix+"/dashboard/user/api_keys", strings.NewReader(string(bytedata)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(AuthorizationHeader, "Bearer "+u.SessionKey())
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return CreateSecretKeyResponse{}, errors.Join(
			errors.New("CreateSecretKeys error"),
			err,
		)
	}
	if resp.StatusCode != http.StatusOK {
		return CreateSecretKeyResponse{}, errors.Join(
			errors.New(fmt.Sprintf("CreateSecretKeys found non 200 response, StatusCode: %v", resp.StatusCode)),
			err)
	}
	data, _ := io.ReadAll(resp.Body)

	var response CreateSecretKeyResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return CreateSecretKeyResponse{}, errors.Join(
			errors.New("GetSecretKey Unmarshal error"),
			err)
	}

	return response, nil
}

type DeleteSecretKeyResponse struct {
	Result string `json:"result"`
}

func (u *UserClient) DeleteSecretKey(key Key) (DeleteSecretKeyResponse, error) {
	if u.SessionKey() == "" {
		return DeleteSecretKeyResponse{}, errors.New("CreateSecretKey with no SessionKey is Defined")
	}
	form := ActionMap{
		Action:      "delete",
		CreateAt:    key.Created,
		RedactedKey: key.SensitiveID,
	}
	bytedata, _ := json.Marshal(form)
	req, err := http.NewRequest(http.MethodPost, PlatformApiUrlPrefix+"/dashboard/user/api_keys", strings.NewReader(string(bytedata)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.SessionKey())
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return DeleteSecretKeyResponse{}, errors.Join(
			errors.New("CreateSecretKeys error"),
			err,
		)
	}
	if resp.StatusCode != http.StatusOK {
		return DeleteSecretKeyResponse{}, errors.Join(
			errors.New(fmt.Sprintf("CreateSecretKeys found non 200 response, StatusCode: %v", resp.StatusCode)),
			err)
	}
	data, _ := io.ReadAll(resp.Body)

	var response DeleteSecretKeyResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return DeleteSecretKeyResponse{}, errors.Join(
			errors.New("GetSecretKey Unmarshal error"),
			err)
	}

	return response, nil
}
