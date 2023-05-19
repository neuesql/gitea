package structs

// WhoAmI Reference: https://github.com/huggingface/huggingface.js/blob/main/packages/hub/src/types/api/api-who-am-i.d.ts

type AccessTokenResponse struct {
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
}
type AuthResponse struct {
	Type        string              `json:"type"`
	AccessToken AccessTokenResponse `json:"accessToken"`
}

type WhoAmIResponse struct {
	// Base info
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	FullName  string `json:"fullname"`
	Email     string `json:"email"`
	CanPay    bool   `json:"canPay"`
	Plan      string `json:"plan"`
	AvatarUrl string `json:"avatarUrl"`
	PeriodEnd int32  `json:"periodEnd"`

	Auth AuthResponse `json:"auth"`

	// Type = User
	EmailVerified bool             `json:"emailVerified"`
	IsPro         bool             `json:"isPro"`
	Orgs          []WhoAmIResponse `json:"orgs"`

	// Type = Org

	// Type = App
	Scope map[string]string `json:"scope"`
}
