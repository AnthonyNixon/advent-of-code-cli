package token

var sessionToken string

func SetSessionToken(input string) {
	sessionToken = input
}

func GetSessionToken() (token string) {
	token = sessionToken
	return
}
