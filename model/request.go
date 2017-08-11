package model

import (
	"cetm/qapi/o/house/counter"
	"cetm/qapi/o/org/customer"
	"cetm/qapi/o/report/transaction"
	"strconv"
)

//InboundRequest :
type InboundRequest struct {
	TransID string  `json:"clientTransactionId"`
	ReqType string  `json:"requestType"`
	Params  []*Pair `json:"param"`
}

//RequestBody :
type RequestBody struct {
	Body InboundRequest `json:"InboundRequest"`
}

//NewRequestBody :
func NewRequestBody(t *transaction.Transaction) (*RequestBody, error) {
	var ctm, err = customer.GetByID(t.CustomerID)
	if nil != err {
		return nil, err
	}
	ct, err := counter.GetByID(t.CounterID)
	if nil != err {
		return nil, err
	}
	var req = &RequestBody{}
	req.Body.TransID = t.ID
	req.Body.ReqType = "QMS_SURVEY"
	req.Body.Params = make([]*Pair, 5)
	req.Body.Params[0] = newPair("ticketnumber", t.ID)
	req.Body.Params[1] = newPair("msisdn", ctm.PhoneNumber)
	req.Body.Params[2] = newPair("store_name", t.BranchID)
	req.Body.Params[3] = newPair("counter_number", ct.Name)
	req.Body.Params[4] = newPair("eventdate", strconv.FormatInt(t.CTime.Unix(), 10))
	return req, nil
}
