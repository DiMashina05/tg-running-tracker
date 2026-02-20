package httpapi

type postNameReqest struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

type postRunsReqest struct {
	UserID int64  `json:"user_id"`
	Dist   string `json:"distance"`
}

type postRunsResponse struct {
	UserID int64   `json:"user_id"`
	Dist   float64 `json:"distance"`
}

type getProfileResponce struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}
