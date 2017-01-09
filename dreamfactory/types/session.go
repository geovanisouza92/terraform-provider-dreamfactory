package types

type Session struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	IsSysAdmin    bool   `json:"is_sys_admin"`
	SessionToken  string `json:"session_token"`
	SessionID     string `json:"session_id"`
	LastLoginDate string `json:"last_login_date"`
	Host          string `json:"host"`
}
