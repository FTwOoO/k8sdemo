package rpc

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type k8sdemoHTTPResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type k8sdemoErrorI interface {
	Error() string
	GetMsg() string
	GetCode() int
}

func RegisterK8sdemoForHTTP(service K8sdemo) {

	http.DefaultServeMux.HandleFunc("/k8sdemo/helloNoReponse", func(w http.ResponseWriter, r *http.Request) {
		var httpStatusCode = http.StatusOK
		var code int
		var msg string
		var err error

		req := new(HelloRequest)
		body, _ := ioutil.ReadAll(r.Body)

		if err = json.Unmarshal(body, req); err != nil {
			httpStatusCode = http.StatusBadRequest
			code = http.StatusBadRequest
			msg = err.Error()
		}

		var reqI interface{} = req
		if v, ok := reqI.(interface{ Validate() error }); ok {
			err := v.Validate()
			if err != nil {
				httpStatusCode = http.StatusBadRequest
				code = http.StatusBadRequest
				msg = err.Error()
			}
		}

		if code == 0 {
			inCtx, _ := context.WithTimeout(context.TODO(), 3*time.Second)
			err = service.HelloNoReponse(inCtx, req)

			if err != nil {
				if v, ok := err.(k8sdemoErrorI); ok {
					code = v.GetCode()
					msg = v.GetMsg()
				} else {
					code = http.StatusInternalServerError
					msg = err.Error()
				}
			} else {
				code = 0
				msg = ""
			}
		}

		respD, _ := json.Marshal(k8sdemoHTTPResp{
			Code: code,
			Msg:  msg,
			Data: nil,
		})

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(httpStatusCode)
		w.Write([]byte(respD))
	})

}
