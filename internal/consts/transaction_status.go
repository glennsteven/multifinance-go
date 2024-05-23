package consts

const (
	Pending  = 0
	Approved = 1
	Rejected = 2
)

var TransactionStatusMap = map[int]string{
	0: "pending",
	1: "approved",
	2: "rejected",
}

// ConvertStatusToString converts numeric status to its string representation
func ConvertStatusToString(status int) string {
	str, ok := TransactionStatusMap[status]
	if !ok {
		return "unknown"
	}
	return str
}
