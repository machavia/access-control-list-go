package internal

import (
	"fmt"
	"strings"
)

func PrintObjectTree(ob *Object, level int) {
	permStr := ""
	for _, perm := range ob.permissions {
		permStr += fmt.Sprintf("%s>%t, ", perm.subject.name, perm.permission)
	}
	fmt.Printf("|%s %s (%s)\n", strings.Repeat("- ", level), ob.name, permStr)
	if len(ob.children) > 0 {
		for _, child := range ob.children {
			PrintObjectTree(child, level+1)
		}
	}
}
