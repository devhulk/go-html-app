package main


import (
	"html/template"
	"net/http"
	"log"
)



type Task struct {
	Title string
	Done  bool
}



type TaskPageData struct {
	PageTitle string
	Tasks     []Task
}

func tasks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/templates/tasks.html"))
	data := TaskPageData{
		PageTitle: "Todays Tasks",
		Tasks: []Task{
			{Title: "Make tiktok video about SSR and Go", Done: false},
			{Title: "Integrate golang SSR and Tailwinds CSS", Done: true},
			{Title: "Carry whittington on fortnite", Done: true},
		},
	}
	tmpl.Execute(w, data)
}


func main() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs)) 		

	http.HandleFunc("/tasks", tasks)

	log.Print("Server is listening on port 3000...")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
