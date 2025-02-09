package constants

import "fmt"

var ErrorInvliadDescriptionLenght = fmt.Errorf("invalid length for given description")
var ErrorInvliadTimeFormat = fmt.Errorf("transaction with invalid time format")
var ErrorTransactionPurchaseAmount = fmt.Errorf("transaction with invalid purchase amount")

// Exchange errors
var ErrorExchangeCreateRequest = fmt.Errorf("failed to create request for treasury API")
var ErrorExchangeRequestAPI = fmt.Errorf("request to treasury API failed")
var ErrorExchangeRequestUnsuccessful = fmt.Errorf("request to treasury API unsuccessful")
var ErrorExchangeDecodeResponse = fmt.Errorf("failed to decode treasury API response")
var ErrorExchangeRequestWithoutResults = fmt.Errorf("treasury API search without results")
