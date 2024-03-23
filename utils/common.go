package utils

type Status struct {
	Code   int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status  Status	`json:"status"`
	Data    interface{} `json:"data"`
}

var StatusCodeMessage = map[int]Status{
	200: {Code: 200, Message: "Success"},
	201: {Code: 201, Message: "Created successfully"},
	204: {Code: 204, Message: "No content"},
	400: {Code: 400, Message: "Bad request"},
}