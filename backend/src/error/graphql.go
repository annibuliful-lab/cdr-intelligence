package error_utils

import "fmt"

type GraphqlError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e GraphqlError) Error() string {
	return fmt.Sprintf("error [%s]: %s", e.Code, e.Message)
}

func (e GraphqlError) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}
}
