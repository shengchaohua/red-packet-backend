package config

import (
	"fmt"
	"strings"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleAPI   Role = "api"
)

func mustParseRole(role string) Role {
	roleEnum := Role(strings.ToLower(role))
	switch roleEnum {
	case RoleAdmin, RoleAPI:
		return roleEnum
	}
	panic(fmt.Errorf("unknown role: %s", role))
}

func (role Role) IsAdmin() bool {
	return role == RoleAdmin
}

func (role Role) IsAPI() bool {
	return role == RoleAPI
}
