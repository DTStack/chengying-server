package response

type Component struct {
	Role    string   `json:"role"`
	Status  int      `json:"status"`
	Message []string `json:"errors"`
}

type ComponentResponse struct {
	List []Component `json:"list"`
}
