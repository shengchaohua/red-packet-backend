package userdm

type DataManager interface {
	LoadById(userID uint64) error
}

func InitUserDM() {

}
