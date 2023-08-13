package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"io"
	"net/url"
	"time"
)

type DailyCost struct {
	Timestamp float64 `json:"timestamp"`
	LineItems []struct {
		Name string  `json:"name"`
		Cost float64 `json:"cost"`
	} `json:"line_items"`
}

type UsageResponse struct {
	Object     string      `json:"object"`
	DailyCosts []DailyCost `json:"daily_costs"`
	TotalUsage float64     `json:"total_usage"`
}

func (u *UserClient) UsageWithSecretKey(sk, StartDate, EndDate string) (UsageResponse, error) {
	if sk == "" {
		return UsageResponse{}, errors.New("GetUsage with no access token is defined")
	}
	return u.usageWithCustomToken(sk, StartDate, EndDate)
}

func (u *UserClient) UsageWithSessionToken(StartDate, EndDate string) (UsageResponse, error) {
	if u.SessionKey() == "" {
		return UsageResponse{}, errors.New("GetUsage get empty session key")
	}
	return u.usageWithCustomToken(u.SessionKey(), StartDate, EndDate)
}

func (u *UserClient) usageWithCustomToken(token, StartDate, EndDate string) (UsageResponse, error) {
	Params := url.Values{
		"end_date":   {EndDate},
		"start_date": {StartDate},
	}

	req, err := http.NewRequest(http.MethodGet, PlatformApiUrlPrefix+"/dashboard/billing/usage?"+Params.Encode(), nil)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set(AuthorizationHeader, "Bearer "+token)
	resp, err := u.client.Do(req)
	if err != nil {
		return UsageResponse{}, errors.Join(
			errors.New("usage error"),
			err,
		)
	}
	if resp.StatusCode != http.StatusOK {
		return UsageResponse{}, errors.Join(
			errors.New(fmt.Sprintf("Usage found non 200 response, StatusCode: %v", resp.StatusCode)),
			err)
	}
	data, _ := io.ReadAll(resp.Body)

	var response UsageResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return UsageResponse{}, errors.Join(
			errors.New("usage Unmarshal error"),
			err)
	}

	return response, nil
}

func GetLastMonth() (string, string) {
	Now := time.Now()
	year := Now.Year()
	month := Now.Month()

	lastMonth := month - 1
	lastYear := year
	if month == time.January {
		lastYear = year - 1
		lastMonth = time.December
	}
	return fmt.Sprintf("%v-%v-01", lastYear, lastMonth.String()),
		fmt.Sprintf("%v-%v-01", year, month.String())
}
