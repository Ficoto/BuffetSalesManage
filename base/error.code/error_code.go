package ec

// ErrorCode - base error struct
type ErrorCode struct {
	Err int    `json:"err"`
	Msg string `json:"msg"`
}

// Clone - clone method for error struct
func (ec *ErrorCode) Clone() *ErrorCode {
	var clone = *ec
	return &clone
}

// Error - return error message
func (ec ErrorCode) Error() string {
	return ec.Msg
}

// ErrorCodeEx - Extension of ErrorCode
type ErrorCodeEx struct {
	ErrorCode
	External interface{} `json:"external"`
}

// Clone - clone method for ex error struct
func (ece *ErrorCodeEx) Clone() *ErrorCodeEx {
	var clone = *ece
	return &clone
}
