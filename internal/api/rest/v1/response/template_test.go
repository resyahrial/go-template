package response_test

import "github.com/resyahrial/go-template/internal/api/rest/v1/response"

func (s *ResponseConverterTestSuite) TestWrapSingleResponse() {
	data := map[string]interface{}{
		"message": "OK",
	}
	wrappedRes := response.WrapSingleData(data)
	s.Equal(&response.Success{
		Data: data,
	}, wrappedRes)
}

func (s *ResponseConverterTestSuite) TestWrapPaginatedResponse() {
	data := []map[string]interface{}{
		{
			"id": "id",
		},
	}

	testCases := []struct {
		name      string
		count     int64
		inputPage int
		limit     int
		pageInfo  response.PageInfo
	}{
		{
			name:      "should have PageInfo{1, 1, 1}",
			count:     1,
			inputPage: 0,
			limit:     10,
			pageInfo: response.PageInfo{
				CurrentPage: 1,
				TotalPage:   1,
				Count:       1,
			},
		},
		{
			name:      "should have PageInfo{1, 10, 100}",
			count:     100,
			inputPage: 0,
			limit:     10,
			pageInfo: response.PageInfo{
				CurrentPage: 1,
				TotalPage:   10,
				Count:       100,
			},
		},
		{
			name:      "should have PageInfo{4, 10, 100}",
			count:     100,
			inputPage: 3,
			limit:     10,
			pageInfo: response.PageInfo{
				CurrentPage: 4,
				TotalPage:   10,
				Count:       100,
			},
		},
		{
			name:      "should have PageInfo{4, 3, 30}",
			count:     30,
			inputPage: 3,
			limit:     10,
			pageInfo: response.PageInfo{
				CurrentPage: 4,
				TotalPage:   3,
				Count:       30,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			wrappedRes := response.WrapPaginatedResponse(data, tc.count, tc.inputPage, tc.limit)
			s.Equal(&response.Success{
				Data:     data,
				PageInfo: tc.pageInfo,
			}, wrappedRes)
		})
	}
}
