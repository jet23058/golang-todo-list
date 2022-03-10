package responses

type Success struct {
	Data       interface{}
	Message    string `string:"success"`
	Status     string `string:"success"`
	StatusCode int    `int:"200"`
}
