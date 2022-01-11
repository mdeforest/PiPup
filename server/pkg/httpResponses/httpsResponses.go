package httpresponses

func ErrorResponse(message string) map[string]string {
	return map[string]string{"status": "error", "reason": message}
}

func OkResponse() map[string]string {
	return map[string]string{"status": "success"}
}
