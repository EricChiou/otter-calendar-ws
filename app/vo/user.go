package vo

type SignUpReqVO struct {
	Account  string `json:"account" req:"true"`
	Password string `json:"password" req:"true"`
}

type LoginReqVO struct {
	Account  string `json:"account" req:"true"`
	Password string `json:"password" req:"true"`
}
