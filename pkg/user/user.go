package user

import "errors"

type User struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Friends []int  `json:"Friends"`
}

type UserDB struct {
	last int
	Db   map[int]User
}

func (udb *UserDB) Init() {
	udb.Db = make(map[int]User)
}

func (udb *UserDB) AddUser(u User) int {
	udb.Db[udb.last] = u
	tmp := udb.last
	udb.last++
	return tmp
}

func (udb *UserDB) DeleteUser(id int) {
	delete(udb.Db, id)
}

func (udb UserDB) CheckUser(id int) error {
	_, found := udb.Db[id]
	if found {
		return nil
	} else {
		return errors.New("user not exist!")
	}
}

func (udb UserDB) GetUserName(id int) string {
	u, _ := udb.Db[id]
	return u.Name
}

func (udb UserDB) GetUserAge(id int) int {
	u, _ := udb.Db[id]
	return u.Age
}

func (udb UserDB) GetUserFriends(id int) []int {
	u, _ := udb.Db[id]
	return u.Friends
}

func (udb *UserDB) MakeFriends(source, target int) error {
	userSource, found := udb.Db[source]
	if !found {
		return errors.New("unable to make friends!")
	}

	userTarget, found := udb.Db[target]
	if !found {
		return errors.New("unable to make friends!")
	}

	userSource.Friends = append(userSource.Friends, target)
	userTarget.Friends = append(userTarget.Friends, source)

	delete(udb.Db, source)
	delete(udb.Db, target)

	udb.Db[source] = userSource
	udb.Db[target] = userTarget

	return nil
}

func (udb *UserDB) UpdateUserAge(id, newAge int) error {
	if u, found := udb.Db[id]; found {
		delete(udb.Db, id)
		u.Age = newAge
		udb.Db[id] = u
		return nil
	}

	return errors.New("User not exist!")
}
