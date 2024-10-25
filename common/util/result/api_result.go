package result

type DataResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewDataResult(success bool, message string, data interface{}) DataResult {
	return DataResult{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func NewResult(success bool, message string) Result {
	return Result{
		Success: success,
		Message: message,
	}
}
