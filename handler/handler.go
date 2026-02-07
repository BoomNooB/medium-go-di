package handler

type Response struct {
	IsOK bool   `json:"isOK"`
	Msg  string `json:"msg,omitempty"`
}

func newOkResponse() Response {
	return Response{
		IsOK: true,
	}
}

const (
	badRequestJSONSyntax = "json not valid"
	badRequestNotValid   = "request not valid"
)

func newBadRequestResponse(msg string) Response {
	return Response{
		IsOK: false,
		Msg:  msg,
	}
}

func newInternalErrorResponse() Response {
	return Response{
		IsOK: false,
		Msg:  "internal server error",
	}
}
