package model

//InboundResponse :
type InboundResponse struct {
	Status   string `json:"status" `
	TransID  string `json:"transactionId"`
	ClientID string `json:"clientTransactionId"`
}

//ResponseBody :
type ResponseBody struct {
	Body InboundResponse `json:"InboundResponse"`
}
