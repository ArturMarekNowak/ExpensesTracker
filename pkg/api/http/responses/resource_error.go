package responses

import "net/http"

type ResourceError struct {
	StatusCode int
	Resource   string
	Message    string
}

func NewNotFoundError(resource string) ResourceError {
	resourceError := ResourceError{}
	resourceError.StatusCode = http.StatusNotFound
	resourceError.Resource = resource
	resourceError.Message = "Resource not found"
	return resourceError
}

func NewBadRequestBodyError(resource string) ResourceError {
	resourceError := ResourceError{}
	resourceError.StatusCode = http.StatusBadRequest
	resourceError.Resource = resource
	resourceError.Message = "Invalid request body"
	return resourceError
}

func NewBadPathParameterError(resource string) ResourceError {
	resourceError := ResourceError{}
	resourceError.StatusCode = http.StatusBadRequest
	resourceError.Resource = resource
	resourceError.Message = "Invalid path parameter"
	return resourceError
}
