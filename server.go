package main

import (
	"fmt"
	"template"
	"http"
	"compress/gzip"
	"mime"
	"path"
	"os"
	"strconv"
)

const errHtml = `
	<html>
		<head>
		<style>
		body{
			font-family: Arial;
		}
		div{
			padding: 15px 30px;
			border: 1px solid black;
			width: 640px;
			margin-top: 40px;
			margin-left: auto;
			margin-right: auto;
			background: white;
			color: #369;
		}
		</style>
		</head><body>
			<div>
				<h1>%v - %s</h1>
				<p>Kengal 0.9.1</p>
			</div>
		</body>
	</html>`

type ServerError struct {
	Code int
	Msg  string
}

func (se *ServerError) Write(w http.ResponseWriter) {
	w.WriteHeader(se.Code)
	errOut := fmt.Sprintf(errHtml, se.Code, se.Msg)
	w.Write([]byte(errOut))
	w.Flush()
}
func Dispatch(w http.ResponseWriter) os.Error {
	w.SetHeader("Content-Type", "text/html; charset=utf-8")
	w.SetHeader("Content-Encoding", "gzip")

	templ, err := template.Parse(View.Themes.Current().Index, nil)
	if err != nil {
		return err
	}
	gz, err := gzip.NewWriter(w)
	if err != nil {
		return err
	}
	err = templ.Execute(gz, View)
	if err != nil {
		return err
	}

	gz.Close()
	return nil
}

func ParseParameters(url, host string) os.Error {
	View.Host = host
	View.Imprint = false
	View.Index = 0
	View.Rubric = 0
	View.Article = 0
	
	fmt.Println(url)
	fmt.Println(host)
	
	dir, file := path.Split(url)
	if file == "" {
		// is Index Startpage
		View.Index = 1
		return nil
	}
	i, err := strconv.Atoi(file)
	if err == nil {
		//is Index non Startpage
		View.Index = i
		if ((i - 1) * PaginatorMax) > len(View.Articles.Index()) {
			View.Index = 0
			return nil
		}
		return nil
	}
	if file == "impressum" {
		// is Impressum
		View.Imprint = true
		return nil
	}
	nextdir := path.Clean(dir)
	dir, file = path.Split(nextdir)
	i, err = strconv.Atoi(file)
	if err != nil {
		return err
	}
	if dir == "/kategorie/" {
		//is Rubricpage
		View.Rubric = i
		if View.Rubrics.Index().Current() == nil {
			View.Rubric = 0
		}
		return nil
	}
	if dir == "/artikel/" {
		//is Articlepage
		View.Article = i
		if View.Articles.Index().Current() == nil {
			View.Article = 0
		}
		return nil
	}
	return os.ENOTDIR
}

func Controller(w http.ResponseWriter, r *http.Request) {
	ParseParameters(r.URL.Path, r.Host)

	w.SetHeader("Content-Type", "text/html; charset=utf-8")
	

	if View.Blogs.Current() == nil {
		se := &ServerError{403, "Forbidden"}
		se.Write(w)
		return
	}
	fmt.Println(View.Index)

	if View.Index != 0 {
		w.SetHeader("Content-Encoding", "gzip")
		View.HeadMeta = fmt.Sprintf(`<meta name="description" content="%s" />`, View.Blogs.Current().Description)
		View.HeadMeta += fmt.Sprintf(`<meta name="keywords" content="%s" />`, View.Blogs.Current().Keywords)
		Dispatch(w)
		return
	}
	if View.Article != 0 {
		w.SetHeader("Content-Encoding", "gzip")
		View.HeadMeta = fmt.Sprintf(`<meta name="description" content="%s" />`, View.Articles.Current().Description)
		View.HeadMeta += fmt.Sprintf(`<meta name="keywords" content="%s" />`, View.Articles.Current().Keywords)
		Dispatch(w)
		return
	}
	if View.Rubric != 0 {
		w.SetHeader("Content-Encoding", "gzip")
		View.HeadMeta = fmt.Sprintf(`<meta name="description" content="%s" />`, View.Rubrics.Current().Description)
		View.HeadMeta += fmt.Sprintf(`<meta name="keywords" content="%s" />`, View.Rubrics.Current().Keywords)
		Dispatch(w)
		return
	}
	if View.Imprint {
		w.SetHeader("Content-Encoding", "gzip")
		View.HeadMeta = fmt.Sprintf(`<meta name="description" content="%s" />`, "Impressum")
		View.HeadMeta += fmt.Sprintf(`<meta name="keywords" content="%s" />`, "Impressum")
		Dispatch(w)
		return
	}
	se := &ServerError{404, "Not Found"}
	se.Write(w)
	return
}
func Images(w http.ResponseWriter, r *http.Request) {
	imagePath := path.Base(r.URL.Path)
	mimeType := mime.TypeByExtension(path.Ext(imagePath))

	w.SetHeader("Content-Type", mimeType)
	w.SetHeader("Cache-Control", "max-age=31104000, public")
	current := View.Themes.Current()
	for _, v := range View.Resources {
		if v.Template == current.ID {
			if v.Name == imagePath {
				w.Write(v.Data)
				w.Flush()
			}
		}
	}
}
func GlobalController(w http.ResponseWriter, r *http.Request) {
	imagePath := path.Base(r.URL.Path)
	mimeType := mime.TypeByExtension(path.Ext(imagePath))
	w.SetHeader("Content-Type", mimeType)
	data := make([]byte, 0)
	for _, v := range View.Globals {
		if v.Name == imagePath {
			data = v.Data
		}
	}
	kind, _ := path.Split(mimeType)
	if kind == "image/" {
		w.Write(data)
		return
	}
	w.SetHeader("Content-Encoding", "gzip")
	gz, _ := gzip.NewWriter(w)
	gz.Write(data)
	gz.Close()
}
func Css(w http.ResponseWriter, r *http.Request) {
	w.SetHeader("Content-Encoding", "gzip")
	w.SetHeader("Content-Type", "text/css")

	gz, _ := gzip.NewWriter(w)
	gz.Write([]byte(View.Themes.Current().Style))
	gz.Close()
}
