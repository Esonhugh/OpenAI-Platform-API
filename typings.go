package platform

//goland:noinspection GoSnakeCaseUsage
import (
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"os"
)

type UserClient struct {
	client       tls_client.HttpClient
	accessToken  string
	sessionKey   string
	lastResponse *http.Response
}

// NewUserPlatformClient called NewHttpClient with log.
func NewUserPlatformClient(accessToken string) *UserClient {
	return &UserClient{
		client:      NewHttpClient(nil),
		accessToken: accessToken,
	}
}

func (u *UserClient) WithCustomHttpClient(client tls_client.HttpClient) *UserClient {
	u.client = client
	return u
}

func (u *UserClient) SessionKey() string {
	return u.sessionKey
}

func (u *UserClient) AccessToken() string {
	return u.accessToken
}

func (u *UserClient) LastResponse() *http.Response {
	return u.lastResponse
}

//goland:noinspection GoUnhandledErrorResult,SpellCheckingInspection
func NewHttpClient(logger tls_client.Logger) tls_client.HttpClient {
	tls_client_logger := tls_client.NewNoopLogger()
	if logger != nil {
		tls_client_logger = logger
	}
	client, _ := tls_client.NewHttpClient(tls_client_logger, []tls_client.HttpClientOption{
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
		tls_client.WithClientProfile(tls_client.Okhttp4Android13),
		tls_client.WithInsecureSkipVerify(), // for debug and proxies.
	}...)

	proxyUrl := os.Getenv("https_proxy")
	if proxyUrl != "" {
		client.SetProxy(proxyUrl)
	}

	return client
}

type GetAccessTokenRequest struct {
	ClientID    string `json:"client_id"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type GetAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Org struct {
	Object      string        `json:"object"`
	ID          string        `json:"id"`
	Created     int           `json:"created"`
	Title       string        `json:"title"`
	Name        string        `json:"name"`
	Description interface{}   `json:"description"`
	Personal    bool          `json:"personal"`
	IsDefault   bool          `json:"is_default"`
	Role        string        `json:"role"`
	Groups      []interface{} `json:"groups"`
}

type Orgs struct {
	Object string `json:"object"`
	Data   []Org  `json:"data"`
}

type Session struct {
	SensitiveID string      `json:"sensitive_id"`
	Object      string      `json:"object"`
	Name        interface{} `json:"name"`
	Created     int         `json:"created"`
	LastUse     int         `json:"last_use"`
	Publishable bool        `json:"publishable"`
}

type User struct {
	Object       string        `json:"object"`
	ID           string        `json:"id"`
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	Picture      string        `json:"picture"`
	Created      int           `json:"created"`
	Groups       []interface{} `json:"groups"`
	Session      Session       `json:"session"`
	Orgs         Orgs          `json:"orgs"`
	IntercomHash string        `json:"intercom_hash"`
	Amr          []interface{} `json:"amr"`
}

type Key struct {
	SensitiveID string `json:"sensitive_id"`
	Object      string `json:"object"`
	Name        string `json:"name"`
	Created     int    `json:"created"`
	LastUse     int    `json:"last_use"`
	Publishable bool   `json:"publishable"`
}
