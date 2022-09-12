package response

type PostsResponse struct {
	Post interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type PostResponse struct {
	Post interface{} `json:"data"`
}
