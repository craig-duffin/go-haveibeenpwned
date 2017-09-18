package haveibeenpwned

import (
	"net/http"
	"encoding/json"
	"log"
	"crypto/sha1"
	"fmt"
)

const APIurl = "https://haveibeenpwned.com/api/v2/"

func GetAccountBreaches(account string) (breaches []Breach) {
	res, err := http.Get(APIurl+"breachedaccount/"+account)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&breaches)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetBreach(breachName string) (breach Breach) {
	res, err := http.Get(APIurl+"breachedaccount/"+breachName)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&breach)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetBreaches()  (breaches []Breach) {
	res, err := http.Get(APIurl+"breaches")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&breaches)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetAccountPastes(account string) (pastes []Paste) {
	res, err := http.Get(APIurl+"pasteaccount/"+account)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pastes)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetDataClasses() (dataClasses []string){
	res, err := http.Get(APIurl+"pasteaccount/"+"dataclasses")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&dataClasses)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CheckPassword(password string) bool {
	hash := sha1.New()
	hash.Write([]byte(password))
	sha := hash.Sum(nil)
	res, err := http.Get(fmt.Sprintf("%spwnedpassword/%x?originalPasswordIsAHash=true",APIurl,sha))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false
}