package util

type commonConst struct {
	Roles
	Status
}

type Roles struct {
	SuperAdmin string
	Admin      string
}

type Status struct {
	Active      string
	InActive    string
	NotVerified string
}

var CommonConst = commonConst{
	Roles: Roles{
		SuperAdmin: "super_admin",
		Admin:      "admin",
	},
	Status: Status{
		Active:      "active",
		InActive:    "inactive",
		NotVerified: "not_verified",
	},
}
