package hotdog

import (
	"github.com/gorilla/mux"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	thatsWhatHeSaid = []string{
		"My chiropractor wanted to massage my glutes with a large spoon",
		"Miley Cyrus is generally accepted as attractive",
		"Bieber is a decent looking guy",
		"I've got man hands",
		"You gotta get some scented candles bro",
		"That's actually very flattering, thank you",
		"Casinos and gay bars. Last night got weird.",
		"Hey, fuck egg shells. If one of those little fuckers falls in your scrambled eggs it's almost impossible to fish it out.",
		"My evenings are pretty booked up",
		"Follow my live tweets bro?",
		"I've got to make my kilt payment",
		"Oh, doing some lunges here? I'll do a couple.",
	}
)

func render(w http.ResponseWriter, i int) {
	tmpl, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, thatsWhatHeSaid[i])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saying(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if id > len(thatsWhatHeSaid) || err != nil {
		http.Redirect(w, r, "http://whatdidalexsaytoday.com", http.StatusFound)
		return
	}

	render(w, id)
}

func index(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	render(w, rand.Intn(len(thatsWhatHeSaid)))
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", saying)
	r.HandleFunc("/", index)
	http.Handle("/", r)
}
