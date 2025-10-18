package main

func main() {
	recipientChannel := make(chan Recipient) // unbuffered channel
	LoadRecipients("./email.csv", recipientChannel)
	defer close(recipientChannel) // Close the channel after loading all recipients
}
