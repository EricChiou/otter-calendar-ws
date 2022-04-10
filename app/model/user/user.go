package user

const (
	Table string = "public.user"
	PK    string = "id"
)

const (
	ID        string = "id"
	Account   string = "account"
	Password  string = "password"
	Name      string = "name"
	Role      string = "role"
	Status    string = "status"
	CreatedAt string = "created_at"
	UpdatedAt string = "updated_at"
)

type Entity struct {
	ID        int    `json:"id,omitempty"`
	Account   string `json:"account,omitempty"`
	Password  string `json:"password,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
