package internal

import (
	"fmt"
)

type Acl struct {
	objects *Object
	subjects []*Subject
	subjectGroups []*SubjectGroup
}

func NewAcl(objects *Object, subjects []*Subject) *Acl {
	return &Acl{
		objects:       objects,
		subjects:      subjects,
		subjectGroups: nil,
	}
}

func (a *Acl) Eval(objectName string, subjectName string) (bool, error) {
	ob := a.objects.getChild(objectName)
	if ob == nil {
		return false, fmt.Errorf("unknown object %s", objectName)
	}

	var sub *Subject
	for _, elem := range a.subjects {
		if elem.name == subjectName {
			sub = elem
		}
	}
	if sub == nil {
		return false, fmt.Errorf("unknown subject %s", subjectName)
	}

	current := ob
	for current != nil{
		subPerm := current.GetPermission(sub)
		if subPerm != nil {
			return *subPerm, nil
		}
		current = current.parent

	}

	return false, nil
}