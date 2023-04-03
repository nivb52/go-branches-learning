package commonController

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpCommonResponse interface {
	NotFound()
	NotImplemented()
	InternalServerError()
	MakeSuccessResponse()
}

type HttpRouter struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (h HttpRouter) NotFound() {
	h.Res.WriteHeader(http.StatusNotFound)
	message := "not found"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) NotImplemented() {
	h.Res.WriteHeader(http.StatusNotImplemented)
	message := "not implemented"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) InternalServerError() {
	h.Res.WriteHeader(http.StatusInternalServerError)
	h.Res.WriteHeader(http.StatusNotImplemented)
	message := "internal server error"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) MakeSuccessResponse(data any) {
	fmt.Println("make success response |  data:", data)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.WriteHeader(http.StatusOK)
	h.Res.Write(jsonBytes)
	return
}
