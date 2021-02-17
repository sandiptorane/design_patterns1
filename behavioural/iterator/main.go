package main

import "fmt"

func main() {
	user1 := &User{
		Name: "Sandip",
		Age:  23,
	}

	user2 := &User{
		Name: "Dinesh",
		Age:  24,
	}

	var userCollection Collection = &UserCollection{
		Users: []*User{user1, user2},
	}

	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)

	}
}