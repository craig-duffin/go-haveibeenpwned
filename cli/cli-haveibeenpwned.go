package main

import (
	"github.com/craig-duffin/go-haveibeenpwned"
	"os"
	"fmt"
)

func main()  {
	args := os.Args

	if !(len(args) == 1 || len(args) ==2){
		displayHelp()
	}


	switch args[1] {
	case "get-account-breaches":
		displayAccountBreaches(args[2])
	case "get-breach":
		displayBreach(args[2])
	case "get-all-breaches":
		displayAllBreaches()
	case "get-account-pastes":
		displayAccountPastes(args[2])
	case "get-data-classes":
	case "check-password":
	default:
		displayHelp()
	}
}

func displayAccountBreaches(account string)  {
	breaches := haveibeenpwned.GetAccountBreaches(account)
	rows := []map[string]interface{}{}
	headers := []string{"Title","Name","Domain","Data Classes","Breach Date"}
	for _, breach := range breaches{
		row := map[string]interface{}{}
		row["Title"] = breach.Title
		row["Name"] = breach.Name
		row["Domain"] = breach.Domain
		row["Data Classes"] = arrayToStringWithCommas(breach.DataClasses)
		row["Breach Date"] = breach.BreachDate
		rows = append(rows,row)
	}
	PrettyPrintMany(headers,rows)
}

func displayBreach(breachName string)  {
	breach := haveibeenpwned.GetBreach(breachName)
	row := map[string]interface{}{}
	row["Title"] = breach.Title
	row["Name"] = breach.Name
	row["Domain"] = breach.Domain
	row["Description"] = breach.Description
	row["Data Classes"] = arrayToStringWithCommas(breach.DataClasses)
	row["Breach Date"] = breach.BreachDate
	row["Added Date"] = breach.AddedDate
	row["Modified Date"] = breach.ModifiedDate
	row["Pwn Count"] = breach.PwnCount
	row["Is Verified?"] = breach.IsVerified
	row["Is Sensitive?"] = breach.IsSensitive
	row["Is Active?"] = breach.IsActive
	row["Is Retired?"] = breach.IsRetired
	row["Is Spam List?"] = breach.IsSpamList
	row["Logo Type"] = breach.LogoType

	PrettyPrintSingle(row)
}

func displayAllBreaches()  {
	breaches := haveibeenpwned.GetBreaches()
	rows := []map[string]interface{}{}
	headers := []string{"Title","Name","Domain","Data Classes","Breach Date"}
	for _, breach := range breaches{
		row := map[string]interface{}{}
		row["Title"] = breach.Title
		row["Name"] = breach.Name
		row["Domain"] = breach.Domain
		row["Data Classes"] = arrayToStringWithCommas(breach.DataClasses)
		row["Breach Date"] = breach.BreachDate
		rows = append(rows,row)
	}
	PrettyPrintMany(headers,rows)
}

func displayAccountPastes(account string)  {
	pastes := haveibeenpwned.GetAccountPastes(account)
	rows := []map[string]interface{}{}
	headers := []string{"Source","ID","Title","Date","Email Count"}
	for _, paste := range pastes{
		row := map[string]interface{}{}
		row["Source"] = paste.Source
		row["ID"] = paste.Id
		row["Title"] = paste.Title
		row["Date"] = paste.Date
		row["Email Count"] = paste.EmailCount
	}
	PrettyPrintMany(headers,rows)
}

func arrayToStringWithCommas(stringArray []string) (stringWithComma string) {
	for _, element := range stringArray{
		stringWithComma += element+","
	}
	return
}

func displayHelp()  {
	fmt.Println("Please enter a correct command:")
	fmt.Println("get-account-breaches <account>            --Retrieves information about all breaches assocaited with account--")
	fmt.Println("get-account-pastes <account>              --Retrieves information about all pastes assocaited with account--")
	fmt.Println("get-breach <breach name>                  --Retrieves information about a specefic breach--")
	fmt.Println("check-password <password>                 --Checks if a password is included in a breach/paste--")
	fmt.Println("get-all-breaches                          --Retrieves information about all breaches--")
	fmt.Println("get-data-classes                          --Retrieves statistics of different data classes--")
}
