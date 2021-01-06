package main

import (
	"acl/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initAcl() *internal.Acl {
	celia := internal.NewSubject("celia")
	machavia := internal.NewSubject("machavia")
	subjects := []*internal.Subject{celia, machavia}

	root := internal.NewObjectTree()

	//settings
	settings := root.AddChild("settings")
	settings.AddChild("platform")
	settingUser := settings.AddChild("users")

	settingUser.AttachPermission(internal.NewPermission(machavia, true))

	//accounts
	accounts := root.AddChild("accounts")
	accounts.AddChild("google")
	accounts.AddChild("facebook")

	accounts.AttachPermission(internal.NewPermission(machavia, true))

	//accounts
	dashboard := root.AddChild("dashboard")
	roi := dashboard.AddChild("roi")
	dashboard.AddChild("ltv")

	dashboard.AttachPermission(internal.NewPermission(machavia, true))
	roi.AttachPermission(internal.NewPermission(machavia, false))



	internal.PrintObjectTree(root, 0)
	return internal.NewAcl(root, subjects)
}

func Test_Permission(t *testing.T) {
	acl := initAcl()
	res, _ := acl.Eval("accounts", "machavia")
	assert.True(t, res)

	res, _ = acl.Eval("google", "machavia")
	assert.True(t, res)

	res, _ = acl.Eval("settings", "machavia")
	assert.False(t, res)

	res, _ = acl.Eval("users", "machavia")
	assert.True(t, res)

	res, _ = acl.Eval("dashboard", "machavia")
	assert.True(t, res)

	res, _ = acl.Eval("roi", "machavia")
	assert.False(t, res)
}
