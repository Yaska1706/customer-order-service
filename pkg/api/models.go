package api

type NewCustomerRequest struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
}
