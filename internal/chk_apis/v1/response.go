package v1

import "net/http"

type requestResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type errResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

func successResponse(data interface{}) (res *requestResponse) {

	if data != nil {
		res = &requestResponse{
			Status: http.StatusOK,
			Msg:    "success",
			Data:   data,
		}
	} else {
		res = &requestResponse{
			Status: http.StatusOK,
			Msg:    "success",
		}
	}

	return
}

func invalidRequest() (res *errResponse) {

	res = &errResponse{
		Status: http.StatusBadRequest,
		Msg:    "invalid_request",
	}
	return
}

// func badRequest() (res *errResponse) {

// 	res = &errResponse{
// 		Status: http.StatusBadRequest,
// 		Msg:    "bad_request",
// 	}
// 	return
// }

func failedRequest(stt int, msg string) (res *errResponse) {

	res = &errResponse{
		Status: stt,
		Msg:    msg,
	}
	return
}
