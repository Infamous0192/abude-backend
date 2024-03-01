package common

type Creds struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role" enums:"superadmin,owner,employee"`
	Status   bool   `json:"status"`
}
