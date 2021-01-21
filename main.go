package main

import (
	"fmt"
	"strings"
	"strconv"
	"log"
	"net/http"
)

/*

Access via localhost:8080/palindrome with POST Method
set body to x-www-form-urlencoded and create an input field (
	key = number
	value = Enter non-zero positive number (FirstNumber  (space) SecondNumber)
)

*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message" : "Hello there"}`))
	case "POST":
		r.ParseForm()
		fmt.Println(r.FormValue("number"))

		if r.FormValue("number") == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message" : "Nothing to process"}`))
		}else{
			tempData := doPalindrome(r.FormValue("number"))
			w.WriteHeader(http.StatusOK)
			w.Write(tempData)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message" : "Not Found"}`))
	}
}

func main() {
	http.HandleFunc("/palindrome", home)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doPalindrome(s string) []byte {

	process := processPalindrome(s)

	if process[0] == "err" {
		return []byte(`{"message" : "Invalid inputted number"}`)
	}else{

		if process[0] == "err-1" {
			return []byte(`{"message" : "Inputted number can't be less than 1 and can't be more than 1.000.000.000"}`)
		}else{

			if process[0] == "err-2" {
				return []byte(`{"message" : "Inputted number can't be less than 1 and can't be more than 1.000.000.000. Also can't be less than first inputted number."}`)
			}else{
				return []byte(`{
					"message" : "Success",
					"primaryNumber" : `+ process[0] +`,
					"secondaryNumber" : `+ process[1] +`,
					"countedPalindrome" : ` + process[2] + `
				}`)
			}

		}

	}

}


func processPalindrome(s string) [3]string {

	var temporaryFirst, temporarySecond, countedPalindrome int
	var retArray [3]string

	words := strings.Split(s, " ") // Split the string into two with space delimeter between two numbers
	parseOne, err := strconv.Atoi(words[0]) // Take the first array from splitted string and convert it's type data to Integer
	parseTwo, err := strconv.Atoi(words[1]) // Take the second array from splitted string and convert it's type data to Integer

	if(err != nil){
		retArray[0] = "err"
		return retArray
	}

	if parseOne < 1 || parseOne > 1000000000 { // Condition if number less than 1 or more than 1 billion
		retArray[0] = "err-1"
		return retArray
	}else{

		if parseTwo < 1 || parseTwo > 1000000000 || parseTwo < parseOne { // Condition if number less than 1 or more than 1 billion and if secondary number is less than primary number
			retArray[0] = "err-2"
			return retArray
		}else{
			temporaryFirst = parseOne
			temporarySecond = parseTwo
		}
		
	}

	// Begin looping logic

	for i := temporaryFirst; i <= temporarySecond; i++ {

		if(isPalindrome(i)){ // If returned value is true add the value of countedPalindrome
			countedPalindrome++
		}

	}

	// End of looping logic

	// Assign to array, and then return it

	retArray[0] = strconv.Itoa(temporaryFirst) // Returned data as String
	retArray[1] = strconv.Itoa(temporarySecond) // Returned data as String
	retArray[2] = strconv.Itoa(countedPalindrome) // Returned data as String

	return retArray

	// End


}

// Logic function, determine whether the thrown number is palindrome number or not
func isPalindrome(x int) bool {

	if x < 0 || (x % 10 == 0 && x != 0) {
		return false // If x is 0 or null, return false already
	}

	var rev int
	for x > rev {
		rev = rev * 10 + x % 10
		x /= 10
	}

	return x == rev || x == rev / 10 // Consider it like odd and even cases.

}
// End of logic