package types

type Role struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

type RoleCreateParams struct {
	Name string `json:"name"`
}

type PermisioParams struct {
	IDRole      int   `json:"id_role"`
	Permissions []int `json:"permissions"`
}

type RoleWithOperations struct {
	Role
	Operations []Operation
}
