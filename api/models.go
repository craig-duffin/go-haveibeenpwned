package api

type Breach struct {
	Title string `json:"Title"`
	Name string `json:"Name"`
	Domain string `json:"Domain"`
	BreachDate string `json:"BreachDate"`
	AddedDate string `json:"AddedDate"`
	ModifiedDate string `json:"ModifiedDate"`
	PwnCount int `json:"PwnCount"`
	Description string `json:"Description"`
	DataClasses []string `json:"DataClasses"`
	IsVerified bool `json:"IsVerified"`
	IsSensitive bool `json:"IsSensitive"`
	IsActive bool `json:"IsActive"`
	IsRetired bool `json:"IsRetired"`
	IsSpamList bool `json:"IsSpamList"`
	LogoType string `json:"LogoType"`
}

type Paste struct {
	Source string `json:"Source"`
	Id string `json:"Id"`
	Title string `json:"Title"`
	Date string `json:"Date"`
	EmailCount int `json:"EmailCount"`
}


