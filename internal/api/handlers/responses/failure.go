package response

type Failure struct {
	StatusCode int         `json:"-"`
	Error      interface{} `json:"error"`
}
