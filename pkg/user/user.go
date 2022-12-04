package user

type User struct {
	Name    string `json:"Name"`
	Age     uint   `json:"Age"`
	Friends []int  `json:"Friends"`
}
