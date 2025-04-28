package entities

type Schedule struct {
	Code     string   `json:"id"`
	Subject  string   `json:"subject"`
	Group    string   `json:"group"`
	Building string   `json:"building"`
	Room     string   `json:"room"`
	Year     int      `json:"year"`
	Semester int      `json:"semester"`
	Classes  []*Class `json:"classes"`
}

func NewSchedule(code, subject, group, building, room string, year, semester int, classes []*Class) *Schedule {
	if classes == nil {
		classes = []*Class{}
	}

	return &Schedule{
		Code:     code,
		Subject:  subject,
		Group:    group,
		Building: building,
		Room:     room,
		Year:     year,
		Semester: semester,
		Classes:  classes,
	}
}

func (s *Schedule) AddClass(c *Class) {
	s.Classes = append(s.Classes, c)
}

func (s *Schedule) SetClasses(classes []*Class) {
	s.Classes = classes
}
