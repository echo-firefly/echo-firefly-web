package Services

import (
	"echo-firefly-web/app/Models"
)

func (this *TestService)GetUserList() ([]Models.User, error) {

	user, err := this.Models.User.GetAll()
	if err != nil {
		return user, err
	}
	return user, nil

}