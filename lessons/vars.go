package lessons

import "fmt"

func Vars() {

	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32

	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year written", year, "page number", pageNumber, "graded", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "year written", year, "page number", pageNumber, "graded", grade)

	animal1 := "cat"
	animal2 := "dog"

	// Add your code below:
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

	if lessonLearned := true; lessonLearned {
		fmt.Println("Great job! You can continue on to the next exercise.")
	} else {
		fmt.Println("Practice makes perfect.")
	}
}
