package user

type CreateRequest struct{
	Username	string		`json:"username"`
	Password	string		`json:"password"`
}
type CreateReponse struct {
	Username	string		`json:username`
}