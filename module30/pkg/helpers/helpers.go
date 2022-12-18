package helpers

func RemoveElementFromArray(friends []string, friendName string) []string {
	needIndex := 0
	for index, friend := range friends {
		if friend == friendName {
			needIndex = index
			break
		}
	}
	newFriends := append(friends[:needIndex], friends[needIndex+1:]...)
	return newFriends
}
