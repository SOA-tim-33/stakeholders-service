package dto

type AccountRegistrationDto struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Pasysword      string `json:"password"`
	ProfilePicture string `json:"profilePicture"`
	Biography      string `json:"biography"`
	Motto          string `json:"motto"`
}
