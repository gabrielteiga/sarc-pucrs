package entities

type Class struct {
	ID          string `json:"id"`
	Day         string `json:"day"`
	Date        string `json:"date"`
	Hour        string `json:"hour"`
	Description string `json:"description"`
	Activity    string `json:"activity"`
	Resource    string `json:"resource"`
}

func NewClass(id, day, date, hour, description, activity, resource string) *Class {
	return &Class{
		ID:          id,
		Day:         day,
		Date:        date,
		Hour:        hour,
		Description: description,
		Activity:    activity,
		Resource:    resource,
	}
}
