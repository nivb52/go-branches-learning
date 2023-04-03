package commonController

import (
	"encoding/json"
	"net/http"
)

type HttpRouter struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (h HttpRouter) NotFound() {
	h.Res.WriteHeader(http.StatusNotFound)
	jsonBytes, err := json.Marshal("{message: not found, code: 404}")
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) NotImplemented() {
	h.Res.WriteHeader(http.StatusNotImplemented)
	jsonBytes, err := json.Marshal("{message: not implemented, code: 501}")
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) InternalServerError() {
	h.Res.WriteHeader(http.StatusInternalServerError)
	jsonBytes, err := json.Marshal("{message: internal server error, code: 500}")
	if err != nil {
		return
	}
	h.Res.Write(jsonBytes)
}

func (h HttpRouter) MakeSuccessResponse(jsonBytes []byte, err error) {
	if err != nil {
		h.InternalServerError()
		return
	}
	h.Res.WriteHeader(http.StatusOK)
	h.Res.Write(jsonBytes)
	return
}

func (h HttpRouter) BedRequestServerError(msg error) {
	h.Res.WriteHeader(http.StatusBadRequest)
	jsonBytes, err := json.Marshal("{message: " + msg.Error() + ", code: 400 }")
	if err != nil {
		return
	}
	h.Res.Write(jsonBytes)
}
