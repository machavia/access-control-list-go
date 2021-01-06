package internal

type Permission struct {
	subject *Subject
	permission bool
}

func NewPermission(subject *Subject, permission bool) *Permission {
	return &Permission{
		subject:    subject,
		permission: permission,
	}
}