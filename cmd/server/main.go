package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/rishabh21g/mailchimp/internal/database"
	"github.com/rishabh21g/mailchimp/internal/handlers"
)

func main() {
	_ , err:= database.Init_DB()
	if err != nil {
		slog.Error("Could not connect to Database" , "error" , err)
	}
	
	// Gin server code started
	router := gin.Default()
	routerAuth := router.Group("/api/v1/auth" )
	routerAuth.POST("/signup",  handlers.RegisterHandler )
	router.Run(":3000")
}

// var wg sync.WaitGroup
// recipientChannel := make(chan types.Recipient) // unbuffered channel
// go services.LoadRecipients("../../internal/data/email.csv", recipientChannel)

// workers := 5
// for i := 0; i <= workers; i++ {
// 	wg.Add(1)
// 	go services.EmailWorker(i, recipientChannel, &wg)
// }
// wg.Wait()
// connect to DB