package main

type User struct {
	Name    string `json:"Name"`
	Age     uint   `json:"Age"`
	Freinds []int
}

var (
	Users map[int]User
)

func main() {

}
