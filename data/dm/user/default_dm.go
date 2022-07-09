package userdm

import (
	"context"
)

type DefaultUserDM struct {
}

var defaultUserDM *DefaultUserDM

func initDefaultUserDM(ctx context.Context) {
	defaultUserDM = &DefaultUserDM{}

}
