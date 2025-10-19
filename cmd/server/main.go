package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rishabh21g/mailchimp/internal/handlers"
)

func main() {
	// var wg sync.WaitGroup
	// recipientChannel := make(chan types.Recipient) // unbuffered channel
	// go services.LoadRecipients("../../internal/data/email.csv", recipientChannel)

	// workers := 5
	// for i := 0; i <= workers; i++ {
	// 	wg.Add(1)
	// 	go services.EmailWorker(i, recipientChannel, &wg)
	// }
	// wg.Wait()
	// Gin server code started
	router := gin.Default()
	routerAuth := router.Group("/api/v1/auth" )
	routerAuth.GET("/login",  handlers.LoginHandler )
	routerAuth.GET("/signup",  handlers.RegisterHandler )
	router.Run(":3000")
}
