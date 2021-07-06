package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// 页面结构体
type Page struct {
	Title string
	Body  []byte
}

// 数据根目录，页面模板根目录
var dataPath = "./data/"
var tmplPath = "./tmpl/"

// 写文件到本地
func (p *Page) save() error {
	filename := dataPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 从本地读文件
func loadPage(title string) (*Page, error) {
	filename := dataPath + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// 模板缓存
// The function `template.Must` is a convenience wrapper
// that panics when passed a non-nil error value,
// and otherwise returns the `*Template` unaltered.
// A panic is appropriate here;
// if the templates can't be loaded
// the only sensible thing to do is exit the program.
var templates = template.Must(
	template.ParseFiles(
		tmplPath+"edit.html",
		tmplPath+"view.html",
	),
)

// 验证
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 验证 path，提取 title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

// 渲染 html 模板
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl + ".html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// - - -
	err := templates.ExecuteTemplate(w, tmpl+".html", p) // 这里的模板名字不能加tmplPath，否则找不到
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'.
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Printf("m = %v\n", m)
		if m == nil {
			// If the title is invalid,
			// an error will be written to the ResponseWriter
			// using the `http.NotFound` function.
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// 处理 Path 以 `/view/` 开的头的请求
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)  // 此段逻辑放到 闭包 makeHandler 中实现，提高代码的可复用性。
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // http.StatusFound=302
		return
	}
	renderTemplate(w, "view", p)
}

// 处理 Path 以 `/edit/` 开头的请求
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// 保存文件到本地，并重定向client的请求到 以`/view/`开头的 Path
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	// The value returned by `FormValue` is of type string.
	// We must convert that value to []byte
	// before it will fit into the Page struct.
	// We use `[]byte(body)` to perform the conversion.
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler) // Add a handler to make the web root redirect to /view/FrontPage.
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Printf("监听地址：localhost:8080\n")
	fmt.Printf("停止服务器请按：ctrl + c\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
