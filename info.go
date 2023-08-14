package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"io"
)

// /dashboard/billing/subscription
// SessionKey

type SubscriptionResponse struct {
	Object             string      `json:"object"`
	HasPaymentMethod   bool        `json:"has_payment_method"`
	Canceled           bool        `json:"canceled"`
	CanceledAt         interface{} `json:"canceled_at"`
	Delinquent         interface{} `json:"delinquent"`
	AccessUntil        int         `json:"access_until"`
	SoftLimit          int         `json:"soft_limit"`
	HardLimit          int         `json:"hard_limit"`
	SystemHardLimit    int         `json:"system_hard_limit"`
	SoftLimitUsd       float64     `json:"soft_limit_usd"`
	HardLimitUsd       float64     `json:"hard_limit_usd"`
	SystemHardLimitUsd float64     `json:"system_hard_limit_usd"`
	Plan               struct {
		Title string `json:"title"`
		ID    string `json:"id"`
	} `json:"plan"`
	Primary        bool        `json:"primary"`
	AccountName    string      `json:"account_name"`
	PoNumber       interface{} `json:"po_number"`
	BillingEmail   interface{} `json:"billing_email"`
	TaxIds         interface{} `json:"tax_ids"`
	BillingAddress struct {
		City       string `json:"city"`
		Line1      string `json:"line1"`
		Line2      string `json:"line2"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
	} `json:"billing_address"`
	BusinessAddress interface{} `json:"business_address"`
}

func (u *UserClient) DashboardSubscription() (SubscriptionResponse, error) {
	if u.SessionKey() == "" {
		return SubscriptionResponse{}, errors.New("getting dashboard Subscription, but sessionkey is empty, you need re-login")
	}

	req, err := http.NewRequest(http.MethodGet, PlatformApiUrlPrefix+"/dashboard/billing/subscription", nil)
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.SessionKey())
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return SubscriptionResponse{}, errors.Join(
			errors.New("GetSubscription Request Error"), err)
	}
	u.lastResponse = resp
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return SubscriptionResponse{}, err
	}
	// data, statusCode, err := u.RawDashboardLogin(u.AccessToken())
	statusCode := resp.StatusCode
	if err != nil {
		return SubscriptionResponse{}, errors.Join(
			errors.New(
				fmt.Sprintf("Try Login to dashboard, got HTTP Status: %v", statusCode),
			),
			err,
		)
	}
	if statusCode != http.StatusOK {
		return SubscriptionResponse{}, errors.Join(
			errors.New(fmt.Sprintf("Try GetSubscription.got non-200 HTTP Status: %v", statusCode)),
			err,
		)
	}

	var subscriptionResponse SubscriptionResponse
	err = json.Unmarshal([]byte(data), &subscriptionResponse)
	if err != nil {
		return SubscriptionResponse{}, errors.Join(
			errors.New("GetSubscription error, bad json data, data: "+string(data)),
			err)
	}

	return subscriptionResponse, nil
}

type PaymentHistoryResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object            string      `json:"object"`
		ID                string      `json:"id"`
		Number            string      `json:"number"`
		AmountDue         int         `json:"amount_due"`
		AmountPaid        int         `json:"amount_paid"`
		Tax               int         `json:"tax"`
		TotalExcludingTax int         `json:"total_excluding_tax"`
		Total             int         `json:"total"`
		Created           int         `json:"created"`
		PdfURL            string      `json:"pdf_url"`
		HostedInvoiceURL  string      `json:"hosted_invoice_url"`
		CollectionMethod  string      `json:"collection_method"`
		DueDate           interface{} `json:"due_date"`
		Status            string      `json:"status"`
	} `json:"data"`
}

func (u *UserClient) PaymentHistory() (PaymentHistoryResponse, error) {
	if u.SessionKey() == "" {
		return PaymentHistoryResponse{}, errors.New("dashboard Onboarding but accessToken is empty, you need re-login")
	}

	req, err := http.NewRequest(http.MethodGet, PlatformApiUrlPrefix+"/dashboard/billing/invoices?system=api", nil)
	req.Header.Set("Authorization", "Bearer "+u.SessionKey())
	req.Header.Set("User-Agent", UserAgent)
	resp, err := u.client.Do(req)
	if err != nil {
		return PaymentHistoryResponse{}, errors.Join(
			errors.New("GetHistoryPayment Request Error"), err)
	}
	u.lastResponse = resp
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PaymentHistoryResponse{}, err
	}

	// data, statusCode, err := u.RawDashboardLogin(u.AccessToken())
	statusCode := resp.StatusCode
	if err != nil {
		return PaymentHistoryResponse{}, errors.Join(
			errors.New(
				fmt.Sprintf("Try Login dashboard, got HTTP Status: %v", statusCode),
			),
			err,
		)
	}
	if statusCode != http.StatusOK {
		return PaymentHistoryResponse{}, errors.Join(
			errors.New(fmt.Sprintf("Try GetHistoryPayment.got non-200 HTTP Status: %v", statusCode)),
			err,
		)
	}

	var subscriptionResponse PaymentHistoryResponse
	err = json.Unmarshal([]byte(data), &subscriptionResponse)
	if err != nil {
		return PaymentHistoryResponse{}, errors.Join(
			errors.New("GetHistoryPayment error, bad json data, data: "+string(data)),
			err)
	}

	return subscriptionResponse, nil
}

func (u *UserClient) GetInvoices() (PaymentHistoryResponse, error) {
	return u.PaymentHistory()
}
