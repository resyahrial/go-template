package response

type Success struct {
	Data     interface{} `json:"data"`
	PageInfo interface{} `json:"pageInfo,omitempty"`
}

type PageInfo struct {
	CurrentPage int   `json:"currentPage,omitempty"`
	TotalPage   int   `json:"totalPage,omitempty"`
	Count       int64 `json:"count,omitempty"`
}

func WrapSingleData(data interface{}) *Success {
	return &Success{
		Data: data,
	}
}

func WrapPaginatedResponse(data interface{}, count int64, currentPage int, limit int) *Success {
	totalPage := 1
	if limit < int(count) {
		addtional := int(count) / limit
		if int(count)%limit == 0 {
			addtional -= 1
		}
		totalPage += addtional
	}
	return &Success{
		Data: data,
		PageInfo: PageInfo{
			CurrentPage: currentPage + 1,
			TotalPage:   totalPage,
			Count:       count,
		},
	}
}
