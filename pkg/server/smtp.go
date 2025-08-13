package server

import (
	"bufio"
	"fmt"
	"log"
	"mailSender/pkg/utils"
	"net/smtp"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func fetchMailAndName() error {
	var filePath string
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		fmt.Println("Enter the path of the file")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			filePath = strings.TrimSpace(scanner.Text())
		}
		if scanner.Err() != nil {
			return fmt.Errorf("Error reading input:", scanner.Err())
		}
	}
	if filePath == "" {
		return fmt.Errorf("No file path provided")
	}
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("Error opening excel file %s: %v", filePath, err)
	}
	defer func() {

		if err := file.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}()
	sheetName := "sheet1"
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("Error reading rows from sheet %s: %v", sheetName, err)
	}
	if len(rows) == 0 {
		return fmt.Errorf("no data found in sheet %s", sheetName)
	}
	return nil
}
func SendMail() error {
	from := "tailq039@gmail.com"
	password := "vbqo fdos lule xsbv"
	name := "Laura Mary Davidson"
	to := []string{"tahamir042@gmail.com"}

	gender, err := utils.FetchGender(name)
	if err != nil {
		fmt.Println("Error fetching gender", err)
		return err
	}
	if gender == "" {
		return fmt.Errorf("gender for %s mail :%s  could not be determined", name, to)
	}
	log.Println(gender)

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Go"
	body := "This is a message from go program"
	message := []byte(subject + "/n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println("error sending mail", err)
		return err
	}
	fmt.Println("Email sent successfully")
	return nil
}
