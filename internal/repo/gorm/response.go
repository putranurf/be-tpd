package gorm

func GetResponse() interface{} {
	return &Response{}
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
