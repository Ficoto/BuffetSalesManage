package ec

// 模块接口基础错误码
const (
	InitCode         int = 10000
	ErrorInitCode    int = iota * InitCode
	ErrorSysCode
	ErrorVideoCode
	ErrorImageCode
	ErrorStorageCode
)

// 基础接口错误码
var (
	Success                = ErrorCode{Err: 0, Msg: "success"}
	InvalidArgument        = ErrorCode{Err: 1, Msg: "invalid argument"}
	LackOfArguments        = ErrorCode{Err: 2, Msg: "lack of arguments"}
	ResourceUnavailable    = ErrorCode{Err: 3, Msg: "resource unavailable"}
	WrongMimetypeJSON      = ErrorCode{Err: 4, Msg: "expect application/json"}
	Retry                  = ErrorCode{Err: 5, Msg: "please try again"}
	TemporarilyUnavailable = ErrorCode{Err: 6, Msg: "temporarily unavailable"}
	Unknown                = ErrorCode{Err: 7, Msg: "unknown error"}
	Fail                   = ErrorCode{Err: 8, Msg: "fail"}
	MongodbOp              = ErrorCode{Err: 9, Msg: "mongodb operation error"}
	LiveRoomQuestion       = ErrorCode{Err: 10, Msg: "not found that question"}
)

// 业务级通用错误码
var (
	InvalidPhone = ErrorCode{Err: ErrorSysCode + 1, Msg: "invalid phone number"}
	SentSmsFail  = ErrorCode{Err: ErrorSysCode + 2, Msg: "sent sms fail"}

	Authentication     = ErrorCode{Err: ErrorStorageCode + 1, Msg: "authenticte error."}
	ObjectNotExist     = ErrorCode{Err: ErrorStorageCode + 2, Msg: "oss object not exist."}
	VideoNotTranscoded = ErrorCode{Err: ErrorStorageCode + 3, Msg: "video not transcode."}
	BucketInvalid      = ErrorCode{Err: ErrorStorageCode + 4, Msg: "bucket invalid."}
	ObjectInvalidSize  = ErrorCode{Err: ErrorStorageCode + 6, Msg: "oss object invalid size."}

	AccountIsExists    = ErrorCode{Err: 10001, Msg: "this account is exists!"}
	AccountIsNotExists = ErrorCode{Err: 10002, Msg: "this account is not exists!"}
	InvalidPassword    = ErrorCode{Err: 10003, Msg: "invalid password!"}
)

const (
	ACCOUNT_IS_EXISTE     = "account is exists"
	ACCOUNT_IS_NOT_EXISTS = "account is not exists"
	INVALID_PASSWORD      = "invalid password"
	LOGIN_SUCCESS         = "login success"
)
