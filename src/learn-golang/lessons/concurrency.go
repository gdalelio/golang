package lessons

import "fmt"

func emit(c chan string) {
	fmt.Println("openning the channel!")
	words := []string{"The", "quick", "brown", "fox"}
	for _, word := range words {
		//transmitting on the channel using <- symbol
		c <- word
	}
	fmt.Println("\n closing the channel!")

	close(c)
}

//Channels is a function to demonstrate concurrency using channels to display words
func Channels() {
	fmt.Println("creating the channel!")
	//creating the channel to communicate on
	wordChannel := make(chan string)

	//starts the function emit as a separate process with a channel to receive the words
	go emit(wordChannel)
	//receive on wordChannel as we range over the words
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
