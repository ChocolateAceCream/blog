# 2023-3-28 Gorm Usage
<a>#gorm</a><a> #cascade delete</a>
## gorm cascade delete
e.g.
for the following role and menu many2many relations

```go
type Menu struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Pid       uint      `json:"pid" gorm:"comment:parent id"`
	Name      string    `json:"name" gorm:"comment:route name"`
	Roles     []Role `json:"roles" gorm:"many2many:roleMenu;constraint:OnDelete:CASCADE;"`
	// ChildMenu []Menu `gorm:"foreignkey:Pid;constraint:OnDelete:CASCADE;"`
}

