package model

type Menu struct {
	Mid int `json:"mid" gorm:"primary_key;autoIncrement"`

	MenuName string `json:"menuName" gorm:"type:varchar(50);not null"`

	Level int `json:"level" gorm:"type:int(2);not null"`

	ParentId int `json:"parentId" gorm:"type:int(11);not null;default:0"`

	Path string `json:"path" gorm:"type:varchar(64);not null"`

	Component string `json:"component" gorm:"type:varchar(64);not null"`

	Roles []*Role `json:"-" gorm:"many2many:menu_role;"`

	Children []Menu `json:"children" gorm:"-"`
}

type Router struct {
	Path string `json:"path"`

	Name string `json:"name"`

	Component string `json:"component"`

	Children []*Router `json:"children"`
}

func ArrayToTree(array []*Menu) []*Router {

	m := make(map[int]*Router)
	var tree []*Router
	for _, i := range array {
		var r Router
		r.Path = i.Path
		r.Name = i.MenuName
		r.Component = i.Component
		if i.ParentId == 0 {
			tree = append(tree, &r)
		} else {
			m[i.ParentId].Children = append(m[i.ParentId].Children, &r)
		}
		m[i.Mid] = &r
	}
	return tree
}
