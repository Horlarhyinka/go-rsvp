package main
import (
	"html/template"
	"net/http"
	"fmt"
)

type attendance struct{
	Name, Email, Tel string
	willAttend bool
} 

func renderTemplate(w http.ResponseWriter, name string, d any){
	temp, err := template.ParseFiles("templates/rsvp.html","templates/" + name + ".html")
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, d)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Execution failed", http.StatusInternalServerError)
		return
	}
}

func renderWelcome(w http.ResponseWriter, request *http.Request){
	fmt.Println("rendering home...")
	renderTemplate(w, "welcome", nil)
}

func renderForm(w http.ResponseWriter, request *http.Request){
	renderTemplate(w, "form", nil)
}

func renderSorry(w http.ResponseWriter, request *http.Request){
	renderTemplate(w, "sorry", nil)
}

func renderThanks(w http.ResponseWriter, request *http.Request){
	renderTemplate(w, "thanks", nil)
}

func renderAttendees(w http.ResponseWriter, request *http.Request){
	attendees := []attendance{
		{Name: "Ola", Email: "testing@dev.co", Tel: "+2349065445036", willAttend: true},
	}
	renderTemplate(w, "attendees", attendees)
}

func main(){
	http.HandleFunc("/attendees", renderAttendees)
	http.HandleFunc("/thanks", renderThanks)
	http.HandleFunc("/sorry", renderSorry)
	http.HandleFunc("/form", renderForm)
	http.HandleFunc("/", renderWelcome)

	http.ListenAndServe(":5000", nil)	
}