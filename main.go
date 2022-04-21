package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var label2 widget.Label
var label3 widget.Label
var label4 widget.Label
var labelhour widget.Label
var labelmin widget.Label
var labelsek widget.Label
var labelnanosek widget.Label

func main() {

	app := app.New()
	w := app.NewWindow("DER GROßE HAHLCOM ONLINESHOPCHECKER    ")
	label1 := widget.NewLabel("Diese Applikation checked, wie lange der Hahlcom.de Shop schon im Rohbau ist.")
	label2.SetText("")
	btn01 := widget.NewButton("KALKULIEREN!", func() { hahlcom() })
	contenttab1 := widget.NewVBox(&label2, label1, &label2, btn01, &label2, &label3, &label4, &labelhour, &labelmin, &labelsek, &labelnanosek, &label2)
	w.SetContent(contenttab1)
	w.Resize(fyne.Size{600, 350})
	w.SetFixedSize(true)
	w.ShowAndRun()
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

	if compare2 {

		label3.SetText("Die Seite kann nicht gefunden werden. Bitte suchen sie https://hahlcom.eu/ auf.")
	}

	if compare {
		daysstr := strconv.Itoa(days)
		hoursstr := fmt.Sprintf("%f", hours)
		minstr := fmt.Sprintf("%f", minutes)
		secstr := fmt.Sprintf("%f", seconds)
		nanosstr := strconv.Itoa(int(nanoseconds))
		label3.SetText("Der Hahlcom.de Shop ist seit der Anmeldung der Domain (17.01.2017) in Arbeit. ")
		label4.SetText("Das sind " + daysstr + " Tage.")
		labelhour.SetText("Das sind: " + hoursstr + " Stunden.")
		labelmin.SetText("Das sind: " + minstr + " Minuten.")
		labelsek.SetText("Das sind: " + secstr + " Sekunden.")
		labelnanosek.SetText("Das sind: " + nanosstr + " Nanosekunden.")
	} else {
		label3.SetText("Der Aufbau der Seite unter https://www.hahlcom.de/index.php/shop hat sich verändert. KÖNNTE DIES DER TAG SEIN?!?!?")
	}

}
