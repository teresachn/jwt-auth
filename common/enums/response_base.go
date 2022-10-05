package enums

type ResponseBase struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

var (
	SUCCESS = ResponseBase{
		ResponseCode:    "00",
		ResponseMessage: "SUCCESS",
	}
	BAD_REQUEST = ResponseBase{
		ResponseCode:    "01",
		ResponseMessage: "BAD REQUEST",
	}
	NOT_FOUND = ResponseBase{
		ResponseCode:    "02",
		ResponseMessage: "NOT_FOUND",
	}
	UNAUTHORIZED = ResponseBase{
		ResponseCode:    "03",
		ResponseMessage: "UNAUTHORIZED",
	}
	INTERNAL_ERROR = ResponseBase{
		ResponseCode:    "04",
		ResponseMessage: "INTERNAL ERROR",
	}
)
