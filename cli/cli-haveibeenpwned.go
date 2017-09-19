package main

import (
	"github.com/craig-duffin/go-haveibeenpwned"
	"os"
	"fmt"
)

func main()  {
	args := os.Args
	fmt.Println(args)
	switch args[1] {
	case "get-account-breaches":
		displayAccountBreaches(args[2])
	case "get-breach":
		displayBreach(args[2])
	case "get-all-breaches":
	case "get-account-pastes":
	case "get-data-classes":
	case "check-password":
	default:
		
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

func arrayToStringWithCommas(stringArray []string) (stringWithComma string) {
	for _, element := range stringArray{
		stringWithComma += element+","
	}
	return
}