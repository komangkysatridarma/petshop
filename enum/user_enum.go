package enum

type UserRole = string

const (
	userAdmin UserRole = "Admin"
	userOwner UserRole = "Owner"
	userStaff UserRole = "Staff"
)