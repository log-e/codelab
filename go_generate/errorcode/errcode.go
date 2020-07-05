package errorcode

//go:generate stringer -type ErrorCode --linecomment
type ErrorCode int

const (
	OK             ErrorCode = 0    //OK
	FILE_NOT_FOUND ErrorCode = 1001 //file not found
	FILE_MOD_ERROR ErrorCode = 1002 //file mode error
	FILE_TOO_BIG   ErrorCode = 1003 //file too big
)