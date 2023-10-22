package token

type TokenRequest struct {
	UserName string
	Pasword  string
} //@name TokenRequest

type TokenResponse struct {
	Token      string
	ExpiryTime int
} //@name TokenResponse

type User struct {
	User_Id string
}
