package apiengine

import (
	"encoding/json"

	"github.com/romeq/testaustime-cli/utils"
)

type User struct {
	Id         int
	RegTime    DateFormat `json:"registration_time"`
	Username   string     `json:"username"`
	FriendCode string     `json:"friend_code"`
}

func (a *Api) GetProfile() (user User) {
	res := a.getRequest("users/@me")
	verifyRequest(res.StatusCode, 200)
	defer res.Body.Close()

	utils.Check(json.NewDecoder(res.Body).Decode(&user))
	return user
}
