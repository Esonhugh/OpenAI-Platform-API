package platform

//goland:noinspection SpellCheckingInspection
const (
	apiCreateChatCompletions = PlatformApiUrlPrefix + "/v1/chat/completions"
	apiCreateCompletions     = PlatformApiUrlPrefix + "/v1/completions"

	platformAuthClientID      = "DRivsnm2Mu42T3KOpqdtwB3NYviHYzwD"
	platformAuthAudience      = "https://api.openai.com/v1"
	platformAuthRedirectURL   = "https://platform.openai.com/auth/callback"
	platformAuthScope         = "openid profile email offline_access"
	platformAuthResponseType  = "code"
	platformAuthGrantType     = "authorization_code"
	platformAuth0Url          = Auth0Url + "/authorize?"
	getTokenUrl               = Auth0Url + "/oauth/token"
	auth0Client               = "eyJuYW1lIjoiYXV0aDAtc3BhLWpzIiwidmVyc2lvbiI6IjEuMjEuMCJ9" // '{"name":"auth0-spa-js","version":"1.21.0"}'
	auth0LogoutUrl            = Auth0Url + "/v2/logout?returnTo=https%3A%2F%2Fplatform.openai.com%2Floggedout&client_id=" + platformAuthClientID + "&auth0Client=" + auth0Client
	dashboardLoginUrl         = "https://api.openai.com/dashboard/onboarding/login"
	getSessionKeyErrorMessage = "Failed to get session key."
)

const (
	PlatformApiPrefix    = "/platform"
	PlatformApiUrlPrefix = "https://api.openai.com"

	defaultErrorMessageKey             = "errorMessage"
	AuthorizationHeader                = "Authorization"
	XAuthorizationHeader               = "X-Authorization"
	ContentType                        = "application/x-www-form-urlencoded"
	UserAgent                          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	Auth0Url                           = "https://auth0.openai.com"
	LoginUsernameUrl                   = Auth0Url + "/u/login/identifier?state="
	LoginPasswordUrl                   = Auth0Url + "/u/login/password?state="
	ParseUserInfoErrorMessage          = "Failed to parse user login info."
	GetAuthorizedUrlErrorMessage       = "Failed to get authorized url."
	GetStateErrorMessage               = "Failed to get state."
	EmailInvalidErrorMessage           = "Email is not valid."
	EmailOrPasswordInvalidErrorMessage = "Email or password is not correct."
	GetAccessTokenErrorMessage         = "Failed to get access token."
	GetArkoseTokenErrorMessage         = "Failed to get arkose token."
	defaultTimeoutSeconds              = 600 // 10 minutes

	EmailKey                       = "email"
	AccountDeactivatedErrorMessage = "Account %s is deactivated."

	ReadyHint = "Service go-chatgpt-api is ready."
)
