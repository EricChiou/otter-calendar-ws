package types

type UserRole string

const (
	Normal = "normal"
)

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)
