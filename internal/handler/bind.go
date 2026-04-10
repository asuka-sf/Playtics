package handler

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// bind and validate JSON request
func bindJSON[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		var ve validator.ValidationErrors
		// check if the error is validator.ValidationErrors and get the error information
		if errors.As(err, &ve) {
			// get the struct type for field lookup
			structType := reflect.TypeOf(*req)
			var msgs []string
			for _, fieldErr := range ve {
				// check if the field exists in the struct and get the json tag name
				if structField, ok := structType.FieldByName(fieldErr.Field()); ok {
					jsonTag := structField.Tag.Get("json")
					// remove option suffix from json tag (e.g. "name,omitempty" -> "name")
					if idx := strings.Index(jsonTag, ","); idx != -1 {
						jsonTag = jsonTag[:idx]
					}
					msgs = append(msgs, jsonTag+" is "+fieldErr.Tag())
				}
			}
			// set a message (e.g. "name is required, email is required")
			message := strings.Join(msgs, ", ")
			if message == "" {
				message = "validation failed"
			}
			errorResponse(c, http.StatusBadRequest, message)
			return false
		}
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return false
	}
	return true
}
