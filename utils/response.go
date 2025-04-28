package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StandardResponses struct {
	Status  string            `json:"status" example:"OK"`
	Message string            `json:"message" example:"Successfully fetched note"`
	Data    any               `json:"data"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type PaginatedMeta struct {
	Page       int   `json:"page" example:"1"`
	Limit      int   `json:"limit" example:"10"`
	Total      int64 `json:"total" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
}

type PaginatedResponses struct {
	Status  string        `json:"status" example:"OK"`
	Message string        `json:"message" example:"Successfully fetched notes"`
	Data    any           `json:"data"`
	Meta    PaginatedMeta `json:"meta"`
}

func StatusMessage(code int, method string, entity any) string {
	messages := map[int]string{
		http.StatusBadRequest:          "Validation error %v",
		http.StatusCreated:             "Successfully created %v",
		http.StatusNotFound:            "%v not found",
		http.StatusInternalServerError: "Internal server error while processing %v",
		http.StatusUnauthorized:        "Unauthorized %v",
	}

	methodMessages := map[string]string{
		"GET":    "Successfully fetched %v",
		"POST":   "Successfully proceed %v",
		"PATCH":  "Successfully updated %v",
		"DELETE": "Successfully deleted %v",
	}

	if template, ok := messages[code]; ok {
		return fmt.Sprintf(template, entity)
	}

	if template, ok := methodMessages[method]; ok {
		return fmt.Sprintf(template, entity)
	}

	return "Unknown status"
}

func HttpResponse(c *gin.Context, data any, code int, entity any, method string, errors map[string]string) {
	response := gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, method, entity),
	}

	if len(errors) > 0 {
		response["errors"] = errors
		response["data"] = nil
	} else {
		response["data"] = data
	}

	c.JSON(code, response)
}

func PaginatedResponse(c *gin.Context, data any, code int, entity any, method string, total int64, page int, limit int) {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(code),
		"message": StatusMessage(code, method, entity),
		"data":    data,
		"meta": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
