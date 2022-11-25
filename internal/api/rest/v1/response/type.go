package response

const (
	SuccessKey   = "SuccessKey"
	FailureKey   = "FailureKey"
	PaginatedKey = "PaginatedKey"
)

type PaginatedResultValue struct {
	Page  int
	Limit int
	Count int64
}
