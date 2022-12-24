package response

type (
	Response struct {
		Data   string `json:"data"`
		Status int    `json:"status"`
	}
)
