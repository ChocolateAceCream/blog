package request

type AddRole struct {
	Pid  *uint  `json:"pid" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type EditRole struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
