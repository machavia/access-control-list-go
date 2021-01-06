package internal

func NewObjectTree() *Object{
	return &Object{
		name:     "root",
		parent:   nil,
		children: []*Object{},
	}
}

type Object struct {
	name     string
	parent   *Object
	children []*Object
	permissions []*Permission
}

func (o *Object) AddChild(name string) *Object {
	child := &Object{
		name:     name,
		parent:   o,
		children: []*Object{},
	}
	o.children = append(o.children, child)
	return child
}

func (o *Object) AttachPermission(permission *Permission) {
	o.permissions = append(o.permissions, permission)
}

func (o *Object) getParent() *Object {
	return o.parent
}

func (o *Object) getChild(name string) *Object {
	for _, elem := range o.children {
		if elem.name == name {
			return elem
		}
		if res := elem.getChild(name); res != nil {
			return res
		}
	}
	return nil
}

func (o *Object) GetPermission(subject *Subject) *bool {
	for _, perm := range o.permissions{
		if perm.subject == subject {
			return &perm.permission
		}
	}

	return nil
}
