package main

import (
	"fmt"
	"maps"
	"unicode/utf8"
)

func main(){
	// basic structure
	// map[<key-type>]<value-type>



	local_map := map[string]int{
		"John Doe" : 20,
		"Jane Doe" : 30,
		// "John Doe" : 40, --> as we can't have duplicate keys in a map , it will throw error
	}


	//file-scope map
	var gl0bal_map = map[byte]int{
		'a' : 1,
		'b' : 2,
	}


	fmt.Println(local_map)
	fmt.Println(gl0bal_map)


	s:= "CR7️⃣" // C - 1byte , R - 1byte , 7 - 1byte , U+FE0F - 3bytes , U+20E3 → 3 bytes (😑 yeah , this 7️⃣ emoji here is actually made of 3 separate things -- by chatGPT)
	fmt.Println(len(s)) // as total len is 9bytes based on the UTF-8 encoding
	fmt.Println(utf8.RuneCountInString(s)) // 5 runes in the string

	for k , v:= range s{ // k gives the index which is by default an integer value and v gives the value stored in that index
		fmt.Println("Index : " , k  , "value : " , v , "value length : " , utf8.RuneLen(v))
	}

	profile := make(map[string]string)

	profile["Name"] = "John Doe"
	profile["Age"] = "20"
	profile["Location"] = "New York"


	prof2 := map[string]string {
		"Name" : "John Doe",
		"Age" : "20",
		"Location" : "New York",
	}

	if maps.Equal(profile , prof2){
		fmt.Println("Maps are equal")
	}

	delete(profile , "Age")

	if maps.Equal(profile , prof2){
		fmt.Println("Maps are equal")
	}else{
		fmt.Println("Maps are not equal")
	}


}