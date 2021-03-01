package rest

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// Day is
type Day struct {
	Weekday time.Weekday
	Chores  []string
}

// ThisWeekPage is
type ThisWeekPage struct {
	Title   string
	DaysMap map[int]Day
}

func thisWeekHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		current := time.Now()

		daysMap := make(map[int]Day)

		emptyArr := make([]string, 20)

		emptyArr[6] = "kur"

		daysMap[0] = Day{current.Weekday(), emptyArr}
		for i := 1; i <= 7; i++ {
			daysMap[i] = Day{current.AddDate(0, 0, i).Weekday(), emptyArr}
		}

		p := ThisWeekPage{Title: "This week", DaysMap: daysMap}

		path := "./pkg/http/rest/templates/thisWeek.html"

		tpl, err := template.New(path).Funcs(template.FuncMap{
			"inc": func(n int) int {
				return n + 1
			},
		}).ParseFiles(path)

		if err != nil {
			fmt.Println("Couldn't parse html for index file! \n\n > Error:\n", err)
			os.Exit(http.StatusNotFound)
		} else {
			err := tpl.Execute(w, p)
			if err != nil {
				fmt.Println("fdsfg Couldn't execute template! \n\n > Error:\n", err)
				os.Exit(http.StatusBadRequest)
			}
		}
	}
}

// SimplePage is
type SimplePage struct {
	Title  string
	Author string
}

func indexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p := SimplePage{Title: "Why the almighty xrexy is the best", Author: "xrexy"}

		tpl, err := template.ParseFiles("./pkg/http/rest/templates/index.html")
		if err != nil {
			fmt.Println("Couldn't parse html for index file! \n\n > Error:\n", err)
			os.Exit(http.StatusNotFound)
		} else {
			err := tpl.Execute(w, p)
			if err != nil {
				fmt.Println("Couldn't execute template! \n\n > Error:\n", err)
				os.Exit(http.StatusBadRequest)
			}
		}

	}
}
