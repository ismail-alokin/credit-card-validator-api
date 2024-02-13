package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/credit-card-validator/utils"
)

func CreditCardValidator(c *gin.Context) {
	var requestBody struct {
		CardNumber string `json:"cardNumber"`
	}

	err := c.BindJSON(&requestBody)
	utils.CheckHttpBadRequest(err, c)

	creditCardNumber := requestBody.CardNumber

	valid, err := verifyLuhnsAlgorithm(creditCardNumber)
	utils.HandleServerError(err, c)

	jsonResponse := formResponse(valid)
	utils.SendSuccessJSONResponse(jsonResponse, c)
}

func verifyLuhnsAlgorithm(cardNumber string) (bool, error) {
	fmt.Println("Veryfying Card Number: ", cardNumber)

	runes := []rune(cardNumber)
	sumOfDigits := 0
	numberOfDigits := len(runes)

	for i := numberOfDigits - 1; i >= 0; i-- {

		digit, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			fmt.Println("Not a digit")
			return false, nil
		}

		if i%2 == 0 || numberOfDigits%2 != 0 {
			doubleOfDigitStr := strconv.Itoa(2 * digit)
			for _, d := range doubleOfDigitStr {
				innerDigit, _ := strconv.Atoi(string(d))
				sumOfDigits += innerDigit
			}
		} else {
			sumOfDigits += digit
		}
	}
	fmt.Println("Final sumOfDigits", sumOfDigits)

	if sumOfDigits%10 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func formResponse(valid bool) map[string]interface{} {
	var msg string

	if valid {
		msg = "Credit Card number is valid!"
	} else {
		msg = "Invalid Credit Card number!"
	}

	return map[string]interface{}{
		"success": valid,
		"message": msg,
	}
}
