package userpkg

type Manager interface {
}

var defaultManagerInstance Manager

func InitManager() {
}

func GetDefaultManager() Manager {
	return defaultManagerInstance
}
