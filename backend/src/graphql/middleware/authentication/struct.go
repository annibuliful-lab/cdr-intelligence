package authentication

type AuthorizationHeader struct {
	Token     string
	ProjectId string
	AccountId string
}

type AuthorizationContext struct {
	Token     string
	ProjectId string
	AccountId string
}
