package haveibeenpwned

import (
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

const APIurl = "https://haveibeenpwned.com/api/v2/"


func DoRequest(request *http.Request) ([]byte, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	var response []byte

	request.Header.Set("User-Agent", "pwnWatch")

	res, getErr := httpClient.Do(request)

	if getErr != nil {
		return response,getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		return response,readErr
	}
	return body, nil
}

func GetAccountBreaches(account string) ([]Breach, error) {
	var breaches []Breach

	req, reqErr := http.NewRequest(http.MethodGet,APIurl+"breachedaccount/"+account,nil)

	if reqErr != nil{
		return breaches,reqErr
	}

	body, err := DoRequest(req)

	if err != nil{
		return breaches, err
	}

	jsonErr := json.Unmarshal(body,&breaches)
	if jsonErr != nil {
		return breaches,jsonErr
	}

	return breaches,nil
}

func GetBreach(breachName string) (Breach, error) {
	var breach Breach
	req, reqErr := http.NewRequest(http.MethodGet,APIurl+"breach/"+breachName,nil)

	if reqErr != nil{
		return breach,reqErr
	}

	body, err := DoRequest(req)

	if err != nil{
		return breach, err
	}

	jsonErr := json.Unmarshal(body,&breach)
	if jsonErr != nil {
		return breach,jsonErr
	}

	return breach,nil
}

func GetBreaches()  ([]Breach, error) {
	var breaches []Breach

	req, reqErr := http.NewRequest(http.MethodGet,APIurl+"breaches",nil)

	if reqErr != nil{
		return breaches,reqErr
	}

	body, err := DoRequest(req)

	if err != nil{
		return breaches, err
	}

	jsonErr := json.Unmarshal(body,&breaches)
	if jsonErr != nil {
		return breaches,jsonErr
	}

	return breaches,nil
}

func GetAccountPastes(account string) ([]Paste, error) {
	var pastes []Paste

	req, reqErr := http.NewRequest(http.MethodGet,APIurl+"pasteaccount/"+account,nil)

	if reqErr != nil{
		return pastes,reqErr
	}

	body, err := DoRequest(req)

	if err != nil{
		return pastes, err
	}

	jsonErr := json.Unmarshal(body,&pastes)
	if jsonErr != nil {
		return pastes,jsonErr
	}

	return pastes,nil
}

/*
var commandMap = make(map[string]Command,0)


type ExecFunc func(client *http.Client, flagSet *flag.FlagSet)

type Command struct {
	flagSet *flag.FlagSet
	exec ExecFunc
}

func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc)  {
	commandMap[name] = Command{flagSet, exec}
}

func InitFlags()  {

}

*/