package dto

type CreateLogInput struct {
	Msg  string `json:"msg"`
	Type string `json:"type"`
}
