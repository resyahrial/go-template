package response

type Success struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
}
