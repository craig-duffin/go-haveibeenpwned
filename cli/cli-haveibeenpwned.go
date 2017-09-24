package main

import (
	"github.com/craig-duffin/go-haveibeenpwned"
	"os"
	"fmt"
	"github.com/crackcomm/go-clitable"
)

func main()  {
	args := os.Args

	if (len(args) <2 || len(args) >3){
		displayHelp()
		os.Exit(1)
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
		displayDataClasses()
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
	prettyPrintMany(headers,rows)
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

	prettyPrintSingle(row)
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
	prettyPrintMany(headers,rows)
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
	prettyPrintMany(headers,rows)
}

func displayDataClasses()  {
	dataClasses := haveibeenpwned.GetDataClasses()
	fmt.Printf("%s",arrayToStringWithCommas(dataClasses))
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

func arrayToStringWithCommas(stringArray []string) (stringWithComma string) {
	for _, element := range stringArray{
		stringWithComma += element+", "
	}
	return
}



// PrettyPrintMany - Pretty prints maps as tables
func prettyPrintMany(headers []string, rows []map[string]interface{}) {
	table := clitable.New(headers)
	for _, row := range rows {
		table.AddRow(row)
	}
	table.Print()
}

// PrettyPrintSingle - Pretty prints map as tables
func prettyPrintSingle(row map[string]interface{}) {
	clitable.PrintHorizontal(row)
}
