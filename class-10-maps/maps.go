package main

import "fmt"

func main() {

	// Its like a structs, but its not allow access specific keys as user.name. To this need access as user["name"].
	// The keys and values from map need to be all same types. Can not allow create a map with strings and intergers
	user := map[string] string {
		"name": "Leonardo",
		"subname": "Suave",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	// map into a map
	user2 := map[string]map[string]string {
		"name": {
			"firstName": "Leonardo",
			"secondName": "Suave",
		},
		"curso": {
			"name": "ADS",
			"field": "Local 1",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["name"]["secondName"])

	// delete a keys map
	delete(user2, "curso")
	fmt.Println(user2)

	// add new keys map
	user2["address"] = map[string]string {
			"street": "Street 1",
			"district": "District 1",
	}
	fmt.Println(user2)
	fmt.Println(user2["address"]["street"])
}