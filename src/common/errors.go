package common

import (
	"errors"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest 		= errors.New("error bad request")
	ErrInternalFailure	= errors.New("error internal failure")
	ErrNotFound			= errors.New("error not found")
)

func ErrStatusChecker(err error, c *gin.Context) error {
	if errors.Is(err, ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{	
			"status": "access_is_true",										
			"body": err.Error(),
		})
	}
	if errors.Is(err, ErrInternalFailure) {
		c.JSON(http.StatusInternalServerError, gin.H{	
			"status": "access_is_true",										
			"body": err.Error(),
		})
	}
	if errors.Is(err, ErrBadRequest) {
		c.JSON(http.StatusBadRequest, gin.H{	
			"status": "access_is_true",										
			"body": err.Error(),
		})
	}
	
	return nil
}
