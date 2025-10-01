package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type TodoList struct {
	Title    string
	TodoList []List
}

type List struct {
	Done bool
	Name string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, from url %s", r.URL.RequestURI())
}

/*
arr - массив с файлами
TemplateDir - их общий путь
*/
func pathJoinTemplate(arr *[]string, TemplateDir string) {
	for i := range *arr {
		(*arr)[i] = filepath.Join(TemplateDir, (*arr)[i])
	}
}

func TodoPage(w http.ResponseWriter, r *http.Request) {
	data := TodoList{
		Title: "Todo List",
		TodoList: []List{
			{Done: false, Name: "OneTodo"},
			{Done: true, Name: "TwoTodo"},
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/todo.html"))

	tmpl.Execute(w, data)
}

func main() {
	const Port string = ":8080"

	// создать файловый сервер
	fs := http.FileServer(http.Dir("assest/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/todo", TodoPage)

	// r - указательный тип request
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		paths := []string{
			"layout.html",
			"index.html",
		}

		pathJoinTemplate(&paths, "templates")

		// ParseFiles вставляют несколько html
		tmpl, err := template.ParseFiles(paths...)

		if err != nil {
			panic(err)
		}

		// tmpl := template.Must(template.ParseFiles(paths...)) короткая запись без проверки

		// в первый параметр сообщаем куда вставить шаблон, второй данные для шаблона
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		paths := []string{
			"templates/layout.html",
			"templates/set.html",
		}

		tmpl := template.Must(template.ParseFiles(paths...))

		tmpl.Execute(w, nil)
	})

	fmt.Printf("Server: http://localhost%s/", Port)

	// создать сервер
	http.ListenAndServe(Port, nil)
}
