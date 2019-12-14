package routes

type ApiMessage struct {
	Message string `json:"message"`
}

type ApiStatus struct {
	Id   int64  `json:"id"`
	Text string `json:"status"`
}

type ApiStatusSimple struct {
	Text string `json:"status"`
}

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field"`
}

type ApiErrorsWithMessage struct {
	ApiMessage
	Errors []ApiError `json:"errors"`
}
