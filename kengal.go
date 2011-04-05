package main

import (
	"fmt"
	"os"
	"flag"
	"http"
	gobzip "github.com/shaban/kengal/gobzip"
)

func (a *Article)getBlog() *Blog{
	for k, v := range View.Blogs{
		if v.ID == a.Blog{
			return View.Blogs[k]
		}
	}
	return nil
}
func (r *Rubric)getBlog() *Blog{
	for k, v := range View.Blogs{
		if v.ID == r.Blog{
			return View.Blogs[k]
		}
	}
	return nil
}

func (t *Theme) Active() bool {
	if View.Blogs.Current().Template == t.ID{
		return true
	}
	return false
}

func (b *Blog) Active() bool {
	if b.ID == View.Blogs.Current().ID{
		return true
	}
	return false
}

func (r *Rubric) Active() bool {
	if r.ID == View.Rubrics.Current().ID{
		return true
	}
	return false
}

func (b Blogs) Current() *Blog {
	for k, v := range b {
		if v.Url == View.Host {
			return b[k]
		}
	}
	return nil
}

func (t Themes) Current() *Theme {
	current := View.Blogs.Current()
	if current == nil{
		return nil
	}
	for k, v := range t {
		if v.ID == current.Template {
			return t[k]
		}
	}
	return nil
}

func (a Articles) Latest() []*Article {
	l := len(a)
	if l < 5 {
		return a
	}
	return a[0:5]
}
func (a Articles) Index() Articles {
	b := View.Blogs.Current()
	s := make([]*Article, 0)
	for k, v := range a {
		if b.ID == v.Blog {

			s = append(s, a[k])
		}
	}
	return s
}

func (a Articles) Next() string {
	if View.Index == 0 {
		return ""
	}
	if len(a) > View.Index*PaginatorMax {
		return fmt.Sprintf("/index/%v", View.Index+1)
	}
	return ""
}

func (a Articles) Paginated() []*Article {
	if View.Index == 0 {
		return nil
	}
	l := len(a)
	if l < PaginatorMax {
		return a
	}
	if (View.Index-1)*PaginatorMax+PaginatorMax > l {
		return a[(View.Index-1)*PaginatorMax : l]
	}
	return a[(View.Index-1)*PaginatorMax : (View.Index-1)*PaginatorMax+PaginatorMax]
}
func (a Articles) Prev() string {
	if View.Index <= 1 {
		return ""
	}
	return fmt.Sprintf("/index/%v", View.Index-1)
}

func (a Articles) Current() *Article {
	if View.Article == 0 {
		return nil
	}
	for k, v := range a {
		if v.ID == View.Article {
			return a[k]
		}
	}
	return nil
}

func (a Articles) Rubric() []*Article {
	if View.Rubric == 0 {
		return nil
	}
	s := make([]*Article, 0)
	for k, v := range a {
		if v.Rubric == View.Rubric {
			s = append(s, a[k])
		}
	}
	if len(s) == 0 {
		return nil
	}
	return s
}
func (r Rubrics) Current() *Rubric {
	if View.Rubric == 0 {
		return nil
	}
	for k, v := range r {
		if v.ID == View.Rubric {
			return r[k]
		}
	}
	return nil
}
func (r Rubrics) Index() Rubrics {
	b := View.Blogs.Current()
	s := make([]*Rubric, 0)
	for k, v := range r {
		if v.Blog == b.ID {
			s = append(s, r[k])
		}
	}
	return s
}

var View = new(Page)
var PaginatorMax = 5
var client = gobzip.DefaultClient

func (a *Article) DateTime() string {
	return a.Date
}
func (a *Article) Path() string {
	return fmt.Sprintf("/artikel/%v/%s", a.ID, a.Url)
}
func (a *Article) RubricPath() string {
	for _, v := range View.Rubrics {
		if v.ID == a.Rubric {
			return v.Path()
		}
	}
	return ""
}
func (a *Article) RubricTitle() string {
	for _, v := range View.Rubrics {
		if v.ID == a.Rubric {
			return v.Title
		}
	}
	return ""
}
func (r *Rubric) Path() string {
	return fmt.Sprintf("/kategorie/%v/%s", r.ID, r.Url)
}

func main() {
	flag.StringVar(&View.Server, "s", "", "Geben Sie hier die IP des Servers an")
	flag.StringVar(&View.Master, "m", "", "Geben Sie hier die Adresse des MasterServers an")

	flag.Parse()

	if View.Server == "" {
		flag.Usage()
		os.Exit(0)
	}
	var err os.Error
	client := gobzip.DefaultClient
	client.Init(View, "db", "/admin/delete/", "/admin/replace/", "/admin/insert/", "/admin/audit/")
	client.Audit(View.Master,View.Server,&View.Articles)
	client.Audit(View.Master,View.Server,&View.Blogs)
	client.Audit(View.Master,View.Server,&View.Rubrics)
	client.Audit(View.Master,View.Server,&View.Globals)
	client.Audit(View.Master,View.Server,&View.Resources)
	client.Audit(View.Master,View.Server,&View.Themes)
	
	err = client.SaveKind(View.Articles)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SaveKind(View.Blogs)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SaveKind(View.Globals)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SaveKind(View.Rubrics)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SaveKind(View.Resources)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err = client.SaveKind(View.Themes)
	if err!= nil{
		fmt.Println(err)
		os.Exit(1)
	}
	client.HandleEvents()
	
	http.HandleFunc("/", Controller)

	http.HandleFunc("/global/", GlobalController)
	http.HandleFunc("/images/", Images)
	http.HandleFunc("/style.css", Css)
	http.HandleFunc("/favicon.ico", GlobalController)
	
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.String())
		os.Exit(1)
	}
}