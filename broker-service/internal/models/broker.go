package models

type Request struct {
	Action string           `json:"action"`
	Data   CreateUserParams `json:"data,omitempty"`
	Log    LogParams        `json:"log,omitempty"`
}

type CreateUserParams struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	DisplayName          string `json:"displayName"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type LogParams struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}
