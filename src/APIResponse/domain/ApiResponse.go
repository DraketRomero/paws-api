package domain

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []T    `json:"data,omitempty"`
}

func NewApiResponse[T any](statusMsg string, message string, data []T) ApiResponse[T] {
	response := ApiResponse[T]{
		Status:  statusMsg,
		Message: message,
		Data:    data,
	}

	return response
}
