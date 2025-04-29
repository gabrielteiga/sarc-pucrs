package service

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gabrielteiga/sarc-pucrs/internal/entities"
	"github.com/gocolly/colly"
)

// Identifica o padrão do título da página
// Exemplo: 46514-4 Linguagens de Programação (31) - 32/215
// – ^(\S+)     captura o ID (uma sequência sem espaços)
// – \s+(.+?)   captura o assunto
// – \s+\((\d+)\) captura a turma entre parênteses
// – \s*-\s*    separador hífen opcionalmente com espaços
// – (\d+)/(\d+)$ captura a paginação “número/número”
var titleRe = regexp.MustCompile(`^(\S+)\s+(.+?)\s+\((\d+)\)\s*-\s*(\d+)/(\d+)$`)

func parseTitle(title string) (id, subject, group, building, room string, err error) {
	m := titleRe.FindStringSubmatch(title)
	if m == nil {
		return "", "", "", "", "", fmt.Errorf("título em formato inesperado: %q", title)
	}
	return m[1], m[2], m[3], m[4], m[5], nil
}

func GetScheduleByURL(url string) map[string]*entities.Schedule {
	c := colly.NewCollector()

	scheduleMap := make(map[string]*entities.Schedule, 0)
	classes := make([]*entities.Class, 0)

	c.OnHTML("span[id=lblTitulo]", func(e *colly.HTMLElement) {
		if e.Response.StatusCode != 200 {
			log.Fatalln("Error during the url visit: ", e.Response.StatusCode)
			return
		}

		code, subject, group, building, room, err := parseTitle(e.Text)
		if err != nil {
			println("It was not possible to parse the subject title: ", err)
			return
		}

		year, err := strconv.Atoi(e.Request.URL.Query().Get("ano"))
		if err != nil {
			log.Fatalln("It's impossible to convert the query param to int: ", err)
		}

		semester, err := strconv.Atoi(e.Request.URL.Query().Get("sem"))
		if err != nil {
			log.Fatalln("It's impossible to convert the query param to int: ", err)
		}

		hashKey := e.Request.URL.Query().Get("id")
		schedule := entities.NewSchedule(code, subject, group, building, room, year, semester, nil)
		scheduleMap[hashKey] = schedule
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		class := entities.NewClass(
			e.ChildText("td:nth-child(1)"),
			e.ChildText("td:nth-child(2)"),
			e.ChildText("td:nth-child(3)"),
			e.ChildText("td:nth-child(4)"),
			e.ChildText("td:nth-child(5)"),
			e.ChildText("td:nth-child(6)"),
			e.ChildText("td:nth-child(7)"),
		)

		hashKey := e.Request.URL.Query().Get("id")
		classes = append(classes, class)
		scheduleMap[hashKey].Classes = append(scheduleMap[hashKey].Classes, class)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal("Error during URL visit: ", err)
	}

	return scheduleMap
}
