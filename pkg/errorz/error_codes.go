package errorz

import "net/http"

// ErrorCode
// error code numbering format <tribeId><serviceId><logic_error><error_code>
// we will 8 digits in total. 1 digit for tribeId, 2 digits for prefixed serviceId, 1 digit logic_error category and 4 digits for error_code
// tribeId = 3
// serviceId = 02 (based on index in the repo list sort by name asc)
// logic_error category = (0: internal, 1: unprocessable, 2: not found, 3: validation)
// error_code = 0001 - 9999
//
// logic_error category definition:
// 0: internal error - error that is not caused by user input, will result in 500 code
// 1: unprocessable - error caused by user input, will result in 422 code
// 2: not found - error caused by data cannot found, will result in 404 code
// 3: validation - error caused by user input, will result in 400 code
// 4: forbidden - error caused by user input, will result in 403 code
type ErrorCode int

const (
	//group for logic_error 0 -> 500
	ErrInternalServerGeneral ErrorCode = 30200001

	//group for logic_error 1 -> 422
	ErrUnprocessedGeneral    ErrorCode = 30210001
	ErrItemDraftAlreadyExist ErrorCode = 30210002
	ErrItemIdNotInUserDraft  ErrorCode = 30210003

	//group for logic_error 2 -> 404
	ErrNotFound        ErrorCode = 30220000
	ErrNotFoundGeneral ErrorCode = 30220001

	//group for logic_error 3 -> 400
	ErrValidationGeneral            ErrorCode = 30230001
	ErrItemDraftInvalidFileContents ErrorCode = 30230002
	ErrItemDraftInvalidFormat       ErrorCode = 30230003
	ErrItemDraftExceedMaxLength     ErrorCode = 30230004

	//group for logic_error 4 -> 403
	ErrItemDraftEditHasNoPermission ErrorCode = 30240001
)

var errorCodeToString = map[ErrorCode]string{
	ErrInternalServerGeneral:        "Internal server error",
	ErrUnprocessedGeneral:           "Unprocessable Data error",
	ErrValidationGeneral:            "Validation Failed",
	ErrNotFoundGeneral:              "%s not found",
	ErrNotFound:                     "Page not found",
	ErrItemDraftAlreadyExist:        "Item draft already exist",
	ErrItemDraftInvalidFileContents: "Invalid excel file format. Please check correction file",
	ErrItemDraftInvalidFormat:       "Invalid excel file format",
	ErrItemDraftExceedMaxLength:     "Item draft exceed max length",
	ErrItemDraftEditHasNoPermission: "User don't have access to edit item draft.",
	ErrItemIdNotInUserDraft:         "Some item IDs are not in current user draft",
}
var errorCodeToHttpCode = []int{
	0: http.StatusInternalServerError,
	1: http.StatusUnprocessableEntity,
	2: http.StatusNotFound,
	3: http.StatusBadRequest,
	4: http.StatusForbidden,
}

func (e ErrorCode) ToHttpCode() int {
	logicCategory := (int(e) / 10000) % 10
	return errorCodeToHttpCode[logicCategory]
}

func (e ErrorCode) ToString() string {
	if msg, ok := errorCodeToString[e]; ok {
		return msg
	}
	return "Unknown error"
}
func (e ErrorCode) ToInt() int {
	return int(e)
}

// ParseErrorCode converts string message to ErrorCode
func ParseErrorCode(message string) (ErrorCode, bool) {
	for code, msg := range errorCodeToString {
		if msg == message {
			return code, true
		}
	}
	return 0, false
}
