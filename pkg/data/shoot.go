package data

type SendShootRequest struct {
	FileData []byte
	FileName string
	ToEmail  string `json:"to_email"`
}

type SendShootResponse struct {
	Status string `json:"status"`
}
