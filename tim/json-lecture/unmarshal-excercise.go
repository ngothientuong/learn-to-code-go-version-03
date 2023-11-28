package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	type Person struct {
		First   string   `json:"First"`
		Last    string   `json:"Last"`
		Age     int      `json:"Age"`
		Sayings []string `json:"Sayings"`
	}
	bs := []byte(s)
	var persons []Person
	err := json.Unmarshal(bs, &persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", persons)
	fmt.Println("\n======================\n")
	for index, person := range persons {
		fmt.Printf("Person %v is with \n First name: %s\n Last name: %s\nAge: %d\nWith saying: %s\n\n ", index, person.First, person.Last, person.Age, person.Sayings[0])
	}
}
