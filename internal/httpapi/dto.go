package httpapi

type postNameReqest struct{
	UserID int64 `json:"user_id"`
	Name string `json:"name"`
}