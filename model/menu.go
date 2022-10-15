package model

type Menu struct {
	Mid int `json:"mid" gorm:"primary_key;autoIncrement"`

	MenuName string `json:"menuName" gorm:"type:varchar(50);not null;uniqueIndex"`

	ParentId int `json:"parentId" gorm:"type:int(11);not null;default:0"`

	Path string `json:"path" gorm:"type:varchar(64);not null"`

	Component string `json:"component" gorm:"type:varchar(64);not null"`

	MenuStatus byte `json:"menuStatus" gorm:"type:tinyint(1);not null;default:1"`

	MenuCreateAt int `json:"menuCreateAt" gorm:"autoCreateTime"`

	MenuUpdateAt int `json:"menuUpdateAt" gorm:"autoUpdateTime"`

	Roles []*Role `json:"-" gorm:"many2many:menu_role;"`

	Children []*Menu `json:"children" gorm:"-"`
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
		}
		m[i.Mid] = &r
	}
	for _, i := range array {
		if i.ParentId == 0 {
			continue
		}
		m[i.ParentId].Children = append(m[i.ParentId].Children, m[i.Mid])
	}
	return tree
}

func MenuToTree(array []*Menu) []*Menu {

	m := make(map[int]*Menu)
	var tree []*Menu
	for _, i := range array {
		if i.ParentId == 0 {
			var r Menu
			r.Path = i.Path
			r.MenuName = i.MenuName
			r.Component = i.Component
			m[i.Mid] = &r
		}
	}
	for _, i := range array {
		var r Menu
		r.Path = i.Path
		r.MenuName = i.MenuName
		r.Component = i.Component
		if i.ParentId == 0 {
			tree = append(tree, &r)
		} else {
			arrs := m[i.ParentId].Children
			if arrs == nil {
				arrs = make([]*Menu, 0)
			}
			m[i.ParentId].Children = append(arrs, &r)
		}
	}
	for _, i := range m {
		tree = append(tree, i)
	}
	return tree
}
