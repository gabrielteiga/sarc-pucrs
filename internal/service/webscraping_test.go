package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ExpectedResponse struct {
	ID       string
	Subject  string
	Group    string
	Year     int
	Semester int
	Building string
	Room     string
}

func TestGetScheduleByURL(t *testing.T) {
	tests := []struct {
		Url      string
		Expected ExpectedResponse
	}{
		{
			Url: "https://sarc.pucrs.br/Default/Export.aspx?id=c37c8651-287e-4b65-a0b8-a286a6eb69be&ano=2025&sem=1",
			Expected: ExpectedResponse{
				ID:       "c37c8651-287e-4b65-a0b8-a286a6eb69be",
				Subject:  "Linguagens de Programação",
				Group:    "31",
				Year:     2025,
				Semester: 1,
				Building: "32",
				Room:     "215",
			},
		},
		{
			Url: "https://sarc.pucrs.br/Default/Export.aspx?id=25b9fe09-df0b-47c0-8abf-d6a570871f6b&ano=2025&sem=1",
			Expected: ExpectedResponse{
				ID:       "25b9fe09-df0b-47c0-8abf-d6a570871f6b",
				Subject:  "Verificação e Validação de Software",
				Group:    "30",
				Year:     2025,
				Semester: 1,
				Building: "32",
				Room:     "414",
			},
		},
		{
			Url: "https://sarc.pucrs.br/Default/Export.aspx?id=e9448fa0-3507-41fe-be4c-f1d8b9a32d43&ano=2025&sem=1",
			Expected: ExpectedResponse{
				ID:       "e9448fa0-3507-41fe-be4c-f1d8b9a32d43",
				Subject:  "Fundamentos de Redes de Computadores",
				Group:    "30",
				Year:     2025,
				Semester: 1,
				Building: "32",
				Room:     "213",
			},
		},
	}

	for _, test := range tests {
		runTest(t, test.Url, test.Expected)
	}
}

func runTest(t *testing.T, url string, expected ExpectedResponse) {
	scheduleMap := GetScheduleByURL(url)

	for key, value := range scheduleMap {
		assert.Equal(t, expected.ID, key)
		assert.Equal(t, expected.Subject, value.Subject)
		assert.Equal(t, expected.Group, value.Group)
		assert.Equal(t, expected.Year, value.Year)
		assert.Equal(t, expected.Semester, value.Semester)
		assert.Equal(t, expected.Building, value.Building)
		assert.Equal(t, expected.Room, value.Room)
	}
}
