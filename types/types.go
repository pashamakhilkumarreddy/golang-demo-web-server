package types

type Founder struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Age     uint32 `json:"age"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

type Response struct {
	Message string `json:"message"`
	StatusCode uint16 `json:"statusCode"`
}