package main

import (
	"fmt"
	"os"
	"flag"
	"http"
)

type Servers []*Server
type Articles []*Article
type Rubrics []*Rubric
type Blogs []*Blog
type Themes []*Theme
type Resources []*Resource
type Globals []*Global

func (a Articles)New()interface{}{
	art := new(Article)
	return art
}
func (a Articles)Insert(insert interface{})interface{}{
	a = append(a,insert.(*Article))
	return a
}
func (b Blogs)New()interface{}{
	blg := new(Blog)
	return blg
}
func (b Blogs)Insert(insert interface{})interface{}{
	b = append(b,insert.(*Blog))
	return b
}
func (g Globals)New()interface{}{
	glb := new(Global)
	return glb
}
func (g Globals)Insert(insert interface{})interface{}{
	g = append(g,insert.(*Global))
	return g
}
func (r Resources)New()interface{}{
	rsrc := new(Resource)
	return rsrc
}
func (r Resources)Insert(insert interface{})interface{}{
	r = append(r,insert.(*Resource))
	return r
}
func (r Rubrics)New()interface{}{
	rb := new(Rubric)
	return rb
}
func (r Rubrics)Insert(insert interface{})interface{}{
	r = append(r,insert.(*Rubric))
	return r
}
func (s Servers)New()interface{}{
	srv := new(Server)
	return srv
}
func (s Servers)Insert(insert interface{})interface{}{
	s= append(s, insert.(*Server))
	return s
}
func (t Themes)New()interface{}{
	thm := new(Theme)
	return thm
}
func (t Themes)Insert(insert interface{})interface{}{
	t = append(t,insert.(*Theme))
	return t
}

func (s Servers) Current() *Server {
	for k, v := range s {
		if v.ID == View.Server {
			return s[k]
		}
	}
	return nil
}

func (srv *Server) Active() bool {
	if View.Server == 0 || srv.ID != View.Server{
		return false
	}
	return true
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
func (b Blogs) Replace(blg *Blog)bool{
	for k, v := range b{
		if v.ID == blg.ID{
			b[k] = blg
			return false
		}
	}
	return true
}
func (r Rubrics) Replace(rb *Rubric)bool{
	for k, v := range r{
		if v.ID == rb.ID{
			r[k] = rb
			return false
		}
	}
	return true
}

func (a Articles) Replace(art *Article)bool{
	for k, v := range a{
		if v.ID == art.ID{
			a[k] = art
			return false
		}
	}
	return true
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

type Server struct {
	ID     int
	IP     string
	Vendor string
	Type   string
	Item   int
}

type Blog struct {
	ID          int
	Title       string
	Url         string
	Template    int
	Keywords    string
	Description string
	Slogan      string
	Server      int
}

type Rubric struct {
	ID          int
	Title       string
	Url    string
	Keywords    string
	Description string
	Blog        int
}

type Article struct {
	ID          int
	Date       	string
	Title       string
	Keywords    string
	Description string
	Text        string
	Teaser      string
	Blog        int
	Rubric      int
	Url         string
}

type Resource struct {
	ID	int
	Name     string
	Template int
	Data     []byte
}

type Global struct {
	ID	int
	Name     string
	Data     []byte
}

/*type BlogError struct {
	Code int
	Msg  string
}*/

type Page struct {
	HeadMeta  string
	Rubrics   Rubrics
	Articles  Articles
	Blogs     Blogs
	Blog      int
	Themes    Themes
	Resources Resources
	Globals   Globals
	Servers   Servers
	Index     int
	Rubric    int
	Article   int
	Server    int
	Imprint   bool
	Host      string
	Master	string
}
type Theme struct {
	ID      int
	Index   string
	Style   string
	Title   string
	FromUrl string
}

var View = new(Page)
var PaginatorMax = 5

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
	flag.IntVar(&View.Server, "s", 0, "Geben Sie hier die ID des Servers an")

	flag.Parse()

	if View.Server == 0 {
		flag.Usage()
		os.Exit(0)
	}

	err := audit()
	if err != nil {
		fmt.Println(err.String())
		os.Exit(1)
	}
	
	for _,v := range View.Articles{
		fmt.Println("artikel")
		saveItem("articles",v,v.ID)
	}
	for _,v := range View.Blogs{
		fmt.Println("blogs")
		saveItem("blogs",v,v.ID)
	}
	for _,v := range View.Globals{
		saveItem("globals",v,v.ID)
	}
	for _,v := range View.Resources{
		saveItem("resources",v,v.ID)
	}
	for _,v := range View.Rubrics{
		saveItem("rubrics",v,v.ID)
	}
	for _,v := range View.Servers{
		saveItem("servers",v,v.ID)
	}
	for _,v := range View.Themes{
		saveItem("themes",v,v.ID)
	}
	
	http.HandleFunc("/", Controller)

	http.HandleFunc("/admin/save/", ClientSave)
	http.HandleFunc("/admin/new/", ClientNew)
	http.HandleFunc("/admin/delete/", ClientDelete)

	http.HandleFunc("/global/", GlobalController)
	http.HandleFunc("/images/", Images)
	http.HandleFunc("/style.css", Css)
	http.HandleFunc("/favicon.ico", GlobalController)
	
	http.ListenAndServe(":80", nil)
	os.Exit(0)
}