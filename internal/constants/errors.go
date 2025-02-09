package constants

import "fmt"

var ErrorInvliadDescriptionLenght = fmt.Errorf("invalid length for given description")
var ErrorInvliadTimeFormat = fmt.Errorf("transaction with invalid time format")
var ErrorTransactionPurchaseAmount = fmt.Errorf("transaction with invalid purchase amount")
