package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/credit-card-validator/api"
)

type RequestBody struct {
	CardNumber string `json:"card_number"`
}

func main() {
	router := gin.Default()
	router.POST("/validate", api.CreditCardValidator)

	router.Run(":8081")
}
