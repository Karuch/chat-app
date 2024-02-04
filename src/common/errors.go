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
	switch err {
    case ErrNotFound:
		c.JSON(http.StatusOK, gin.H{	
			"status": "access_is_true",										
			"body": err.Error(),
		})
    case ErrInternalFailure:
		c.JSON(http.StatusInternalServerError, gin.H{	
			"status": "access_is_true",										
			"body": ErrInternalFailure.Error(),
		})
    case ErrBadRequest:
		c.JSON(http.StatusBadRequest, gin.H{	
			"status": "access_is_true",										
			"body": err.Error(),
		})
    }
	return nil
}