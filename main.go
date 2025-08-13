package main

import (
	"log"
	"mailSender/pkg/server"
)

//500 >=
//50k <=

func main() {
	err := server.SendMail()
	if err != nil {
		log.Println(err)
		return
	}
}

//	if category == "Therapist" || category == "Enterpanuer"{
//		subject = "SELF MADE WOMEN FEATURE"
//	   }  else if category == "Coaches" || category == "Bussiness consultant" ||  category == "Gym trainers" || categor{
//	       subject := "MSN TOP 10 COACHES FEATURE"
//	   }
