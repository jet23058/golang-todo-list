package responses

type Fail struct {
	Data       interface{} `json:"[]"`
	Message    string      `string:"failed"`
	Status     string      `string:"failed"`
	StatusCode int         `int:"422"`
}
