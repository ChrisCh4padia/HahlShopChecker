package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

var daytime = "Guten Morgen"
var pause = "               "
var beforpause = pause
var neofetch bool = false

func main() {

	intro()
}

func intro() {

	fmt.Println("~~~~~~~~~~~~~~~~DER GROßE HAHLCOM ONLINESHOPCHECKER~~~~~~~~~~~~~~~~")
	fmt.Println()
	for i := 0; i < 13; i++ {
		fmt.Printf(" # # #")
		time.Sleep(150 * time.Millisecond)
	}
	println(" ")
	sideway()

	for i := 0; i < 13; i++ {
		fmt.Printf(" # # #")

		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println(" ")
	fmt.Println()
	hahlcom()
}

func sideway() {

	today := time.Now().Weekday()
	dateday := time.Now().Day()
	datemonth := time.Now().Month()
	dateyear := time.Now().Local().Year()

	hour := time.Now().Hour()
	Minute := time.Now().UTC().Minute()

	/*
		datedaychar := 0
		datemonthchar := 0
		dateyearchar := 0
			switch {
			case time.Now().Weekday() >= 10:
				todaychar = 1
			}
	*/

	dateday2string := strconv.Itoa(dateday)
	dateyear2string := strconv.Itoa(dateyear)
	todaylength := len(dateday2string)
	monthlength := len(datemonth.String())
	daylength := len(today.String())
	yearlength := len(dateyear2string)
	distance := daylength + yearlength + monthlength + todaylength

	switch {
	case distance == 9:
		pause = "                                                           "
	case distance == 10:
		pause = "                                                          "
	case distance == 11:
		pause = "                                                         "
	case distance == 12:
		pause = "                                                        "
	case distance == 13:
		pause = "                                                       "
	case distance == 14:
		pause = "                                                      "
	case distance == 15:
		pause = "                                                     "
	case distance == 16:
		pause = "                                                    "
	case distance == 17:
		pause = "                                                   "
	case distance == 18:
		pause = "                                                  "
	case distance == 19:
		pause = "                                                 "
	case distance == 20:
		pause = "                                                "
	case distance == 21:
		pause = "                                               "
	case distance == 22:
		pause = "                                              "
	case distance == 23:
		pause = "                                             "
	case distance == 24:
		pause = "                                            "
	case distance == 25:
		pause = "                                           "
	case distance == 26:
		pause = "                                          "
	case distance == 27:
		pause = "                                         "
	case distance == 28:
		pause = "                                        "
	case distance == 29:
		pause = "                                       "
	case distance == 30:
		pause = "                                      "
	case distance == 31:
		pause = "                                     "

	}

	hour2string := strconv.Itoa(hour)
	minute2string := strconv.Itoa(Minute)
	hourlength := len(hour2string)
	minutelength := len(minute2string)

	//minutelength = 1

	timelength := hourlength + minutelength

	timepause := " "
	switch {
	case timelength == 2:
		timepause = "                                                    "
	case timelength == 3:
		timepause = "                                                   "
	case timelength == 4:
		timepause = "                                                  "

	}
	//fmt.Println(timelength)

	fmt.Println(" #                                                                           #")
	fmt.Println(" # ", today, dateday, datemonth, dateyear, pause, "#")
	fmt.Println(" #  Current Time:", hour, ":", Minute, timepause, "#")

	for i := 0; i < 2; i++ {
		fmt.Printf(" # ")
		fmt.Printf("                                                                          #")
		fmt.Println(" ")
		time.Sleep(150 * time.Millisecond)
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	länge := len(username)

	t := time.Now()
	switch {
	case t.Hour() < 10:
		daytime = "Guten Morgen,"
		pause = "                        "
		switch {
		case länge == 9:
			pause = "                       "
		case länge == 10:
			pause = "                      "
		case länge == 11:
			pause = "                     "
		case länge == 12:
			pause = "                    "
		case länge == 13:
			pause = "                   "
		case länge == 14:
			pause = "                  "
		case länge == 15:
			pause = "                 "
		case länge == 16:
			pause = "                "
		case länge == 17:
			pause = "               "
		case länge == 18:
			pause = "              "
		case länge == 19:
			pause = "             "
		case länge == 20:
			pause = "            "
		case länge == 21:
			pause = "           "
		case länge == 22:
			pause = "          "
		case länge == 23:
			pause = "         "
		case länge == 24:
			pause = "        "
		case länge == 25:
			pause = "       "
		case länge == 26:
			pause = "      "
		case länge == 27:
			pause = "     "
		case länge == 28:
			pause = "    "
		case länge == 29:
			pause = "   "
		case länge == 30:
			pause = "  "
		case länge == 31:
			pause = " "

		}
	case t.Hour() > 16:
		daytime = "Guten Abend,"
		pause = "                        "
		switch {
		case länge == 9:
			pause = "                        "
		case länge == 10:
			pause = "                       "
		case länge == 11:
			pause = "                      "
		case länge == 12:
			pause = "                     "
		case länge == 13:
			pause = "                    "
		case länge == 14:
			pause = "                   "
		case länge == 15:
			pause = "                  "
		case länge == 16:
			pause = "                 "
		case länge == 17:
			pause = "                "
		case länge == 18:
			pause = "               "
		case länge == 19:
			pause = "              "
		case länge == 20:
			pause = "             "
		case länge == 21:
			pause = "            "
		case länge == 22:
			pause = "           "
		case länge == 23:
			pause = "          "
		case länge == 24:
			pause = "         "
		case länge == 25:
			pause = "        "
		case länge == 26:
			pause = "       "
		case länge == 27:
			pause = "      "
		case länge == 28:
			pause = "     "
		case länge == 29:
			pause = "    "
		case länge == 30:
			pause = "   "
		case länge == 31:
			pause = "  "
		case länge == 32:
			pause = " "
		}
	default:
		daytime = "Guten Tag,"
		pause = "                          "
		switch {
		case länge == 8:
			pause = "                           "
		case länge == 9:
			pause = "                          "
		case länge == 10:
			pause = "                         "
		case länge == 11:
			pause = "                        "
		case länge == 12:
			pause = "                       "
		case länge == 13:
			pause = "                      "
		case länge == 14:
			pause = "                     "
		case länge == 15:
			pause = "                    "
		case länge == 16:
			pause = "                   "
		case länge == 17:
			pause = "                  "
		case länge == 18:
			pause = "                 "
		case länge == 19:
			pause = "                "
		case länge == 20:
			pause = "               "
		case länge == 21:
			pause = "              "
		case länge == 22:
			pause = "             "

		}
	}

	fmt.Println(" #", "                         ", daytime, username, pause, "#")
	for i := 0; i < 4; i++ {
		fmt.Printf(" # ")
		fmt.Printf("                                                                          #")
		fmt.Println(" ")
		time.Sleep(150 * time.Millisecond)
	}

}

func exit() {

	os.Exit(1)
}

func hahlcom() {
	url := "https://www.hahlcom.de/index.php/shop"
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	string3 := "<title>Fehler: 404 Seite nicht gefunden</title>"
	string2 := "/></a><strong>aktuell ist der Shop noch in der Entwicklung</strong><a"
	string := string(html)

	compare := (strings.Contains(string, string2))
	compare2 := (strings.Contains(string, string3))

	date := time.Now()
	//fmt.Println(date)

	format := "2006-01-02 15:04:05"
	then, _ := time.Parse(format, "2017-01-17 00:00:00")
	//fmt.Println(then)

	diff := date.Sub(then)

	//func Since(t Time) Duration
	//Since returns the time elapsed since t.
	//It is shorthand for time.Now().Sub(t).

	hours := (diff.Hours())             // number of Hours
	nanoseconds := (diff.Nanoseconds()) // number of Nanoseconds
	minutes := (diff.Minutes())         // number of Minutes
	seconds := (diff.Seconds())         // number of Seconds

	days := (int(diff.Hours() / 24)) // number of days

	if compare2 == true {
		fmt.Println("Die Seite kann nicht gefunden werden. Bitte suchen sie https://hahlcom.eu/ auf.")
	}

	if compare == true {

		fmt.Println("Der Hahlcom.de Shop ist seit der Anmeldung der Domain (17.01.2017) in Arbeit. Das sind", days, "Tage.", "\n", "Das sind:", "\n", hours, "Stunden.", "\n", minutes, "Minuten.", "\n", seconds, "Sekunden.", "\n", nanoseconds, "Nanosekunden.")
	} else {
		fmt.Println("Der Aufbau der Seite unter https://www.hahlcom.de/index.php/shop hat sich verändert. KÖNNTE DIES DER TAG SEIN?!?!?")
	}

	fmt.Println()
	fmt.Println("Vielen Dank für ihre Aufmerksamkeit.")
	fmt.Println("Auf Wiedersehen.")
	time.Sleep(20 * time.Second)
}
