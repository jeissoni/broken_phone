package chat

// Room struct
type Room struct {
	Id   int
	Name string
	//list of users
	Users []*User
}

// maxUsers for room
const maxUsers = 2

// RoomsMap for rooms
var RoomsMap = make(map[int]*Room)

// CreateRoom function
func CreateRoom(name string) int {
	r := Room{Id: len(RoomsMap) + 1, Name: name}
	RoomsMap[r.Id] = &r
	return r.Id
}

// AddUserToRoom function
func AddUserToRoom(roomId int, user *User) {
	r := RoomsMap[roomId]
	if len(r.Users) < maxUsers {
		r.Users = append(r.Users, user)
	}
}

// DeleteUserFromRoom delete user from room
func DeleteUserFromRoom(roomId int, user *User) {

	r := RoomsMap[roomId]

	if len(r.Users) == 0 {
		return
	}

	for i, u := range r.Users {
		if u == user {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
			return
		}
	}

}

func GetRoomEmpty() int {
	for _, r := range RoomsMap {
		if len(r.Users) < maxUsers {
			return r.Id
		}
	}
	return 0
}
