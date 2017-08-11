package home

import (
	"cetm/qapi/o/report/transaction"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mobiera/config"
	"mobiera/model"
	"net/http"
	"strings"
)

type Server struct {
	*http.ServeMux
}

//NewServer :
func NewServer() *Server {
	var res = &Server{}
	res.ServeMux = http.NewServeMux()
	res.ServeMux.HandleFunc("/transaction", res.handleDoneTransaction)
	res.ServeMux.HandleFunc("/response", res.handleResponse)
	return res
}

//ServerHTTP :
func (s *Server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeMux.ServeHTTP(w, r)
}

func (s *Server) handleDoneTransaction(w http.ResponseWriter, r *http.Request) {
	var trans = &transaction.Transaction{}
	raw, err := ioutil.ReadAll(r.Body)
	if nil != err {
		fmt.Println("Body request error: ", err.Error())
		return
	}
	err = json.Unmarshal(raw, trans)
	if nil != err {
		fmt.Println("Body request must be json format: ", err.Error())
		return
	}

	request, err := model.NewRequestBody(trans)
	if nil != err {
		fmt.Println("Cant create request body: ", err.Error())
		return
	}

	raw, err = json.Marshal(request)
	if nil != err {
		fmt.Println("Cant marshall request body: ", err.Error())
		return
	}

	resp, err := http.Post(config.GetSurveyHost(), "text/json", strings.NewReader(string(raw)))
	if nil != err {
		fmt.Printf("Cant do request to server %s : %s\n", config.GetSurveyHost(), err.Error())
		return
	}
	defer resp.Body.Close()
}

func (s *Server) handleResponse(w http.ResponseWriter, r *http.Request) {

}
