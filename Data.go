package main

import (
	"os"
	gobzip "github.com/shaban/kengal/gobzip"
	"strconv"
	"time"
	"fmt"
)

type Articles []*Article
type Rubrics []*Rubric
type Blogs []*Blog
type Themes []*Theme
type Resources []*Resource
type Globals []*Global

func (s *Article) Key() int {
	return s.ID
}
func (s *Blog) Key() int {
	return s.ID
}
func (s *Global) Key() int {
	return s.ID
}
func (s *Resource) Key() int {
	return s.ID
}
func (s *Rubric) Key() int {
	return s.ID
}
func (s *Theme) Key() int {
	return s.ID
}
func (s *Article) Log() string {
	return fmt.Sprintf("Typ: Article, ID: %v, Title: %s, Url: %s", s.ID, s.Title, s.Url)
}
func (s *Blog) Log() string {
	return fmt.Sprintf("Typ: Blog, ID: %v, Title: %s, Url: %s", s.ID, s.Title, s.Url)
}
func (s *Global) Log() string {
	return fmt.Sprintf("Typ: Global, ID: %v, Name: %s", s.ID, s.Name)
}
func (s *Resource) Log() string {
	return fmt.Sprintf("Typ: Resource, ID: %v, Name: %s", s.ID, s.Name)
}
func (s *Rubric) Log() string {
	return fmt.Sprintf("Typ: Rubric, ID: %v, Title: %s, Url: %s", s.ID, s.Title, s.Url)
}

func (s *Theme) Log() string {
	return fmt.Sprintf("Typ: Theme, ID: %v, Title: %s, FromUrl: %s", s.ID, s.Title, s.FromUrl)
}
func (s *Article) Kind() string {
	return "articles"
}
func (s *Blog) Kind() string {
	return "blogs"
}
func (s *Global) Kind() string {
	return "globals"
}
func (s *Resource) Kind() string {
	return "resources"
}
func (s *Rubric) Kind() string {
	return "rubrics"
}
func (s *Theme) Kind() string {
	return "themes"
}
func (ser Articles) Kind() string {
	return "articles"
}
func (ser Blogs) Kind() string {
	return "blogs"
}
func (ser Globals) Kind() string {
	return "globals"
}
func (ser Resources) Kind() string {
	return "resources"
}
func (ser Rubrics) Kind() string {
	return "rubrics"
}
func (ser Themes) Kind() string {
	return "themes"
}
func (ser Articles) New() gobzip.Serial {
	return new(Article)
}
func (ser Blogs) New() gobzip.Serial {
	return new(Blog)
}
func (ser Globals) New() gobzip.Serial {
	return new(Global)
}
func (ser Resources) New() gobzip.Serial {
	return new(Resource)
}
func (ser Rubrics) New() gobzip.Serial {
	return new(Rubric)
}
func (ser Themes) New() gobzip.Serial {
	return new(Theme)
}
func (ser Articles) All(ins gobzip.Serializer) {
	View.Articles = ins.(Articles)
}
func (ser Blogs) All(ins gobzip.Serializer) {
	View.Blogs = ins.(Blogs)
}
func (ser Globals) All(ins gobzip.Serializer) {
	View.Globals = ins.(Globals)
}
func (ser Resources) All(ins gobzip.Serializer) {
	View.Resources = ins.(Resources)
}
func (ser Rubrics) All(ins gobzip.Serializer) {
	View.Rubrics = ins.(Rubrics)
}
func (ser Themes) All(ins gobzip.Serializer) {
	View.Themes = ins.(Themes)
}
func (ser Articles) NewKey() int {
	id := 0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Blogs) NewKey() int {
	id := 0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Globals) NewKey() int {
	id := 0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Resources) NewKey() int {
	id := 0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Rubrics) NewKey() int {
	id :=0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Themes) NewKey() int {
	id := 0
	for _, v := range ser {
		if v.ID > id {
			id = v.ID
		}
	}
	return id + 1
}
func (ser Articles)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Blogs)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Globals)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Resources)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Rubrics)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Themes)At(key int)gobzip.Serial{
	for k, v := range ser{
		if v.ID == key{
			return ser[k]
		}
	}
	return nil
}
func (ser Articles) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Article))
	return ser
}
func (ser Blogs) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Blog))
	return ser
}
func (ser Globals) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Global))
	return ser
}
func (ser Resources) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Resource))
	return ser
}
func (ser Rubrics) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Rubric))
	return ser
}
func (ser Themes) Insert(s gobzip.Serial) gobzip.Serializer {
	ser = append(ser, s.(*Theme))
	return ser
}
func (ser Articles) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Article)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Blogs) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Blog)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Globals) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Global)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Resources) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Resource)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Rubrics) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Rubric)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Themes) Replace(s gobzip.Serial) os.Error {
	for k, v := range ser {
		if v.ID == s.Key() {
			ser[k] = s.(*Theme)
			return nil
		}
	}
	return os.ENOENT
}
func (ser Articles) Init()gobzip.Serializer{
	s := make([]*Article, 0)
	var o Articles = s
	return o
}
func (ser Blogs) Init()gobzip.Serializer{
	s := make([]*Blog, 0)
	var o Blogs = s
	return o
}
func (ser Globals) Init()gobzip.Serializer{
	s := make([]*Global, 0)
	var o Globals = s
	return o
}
func (ser Resources) Init()gobzip.Serializer{
	s := make([]*Resource, 0)
	var o Resources = s
	return o
}
func (ser Rubrics) Init()gobzip.Serializer{
	s := make([]*Rubric, 0)
	var o Rubrics = s
	return o
}
func (ser Themes) Init()gobzip.Serializer{
	s := make([]*Theme, 0)
	var o Themes = s
	return o
}
func (ser Articles) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}
func (ser Blogs) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}
func (ser Globals) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}
func (ser Resources) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}
func (ser Rubrics) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}
func (ser Themes) Keys()[]int{
	keys := make([]int,0)
	for _, v := range ser{
		keys = append(keys, v.ID)
	}
	return keys
}

func (ser Articles) NewFromForm(from map[string][]string) gobzip.Serial {
	a := new(Article)
	key :=View.KeyFromForm(from)
	if key == 0 {
		a.ID = ser.NewKey()
	} else {
		a.ID = key
	}
	a.Blog, _ = strconv.Atoi(from["Blog"][0])
	a.Date = time.LocalTime().Format("02.01.2006 15:04:05")
	a.Description = from["Description"][0]
	a.Keywords = from["Keywords"][0]
	a.Rubric, _ = strconv.Atoi(from["Rubric"][0])
	a.Teaser = from["Teaser"][0]
	a.Text = from["Text"][0]
	a.Title = from["Title"][0]
	a.Url = from["Url"][0]
	return a
}
func (ser Blogs) NewFromForm(from map[string][]string) gobzip.Serial {
	a := new(Blog)
	key :=View.KeyFromForm(from)
	if key == 0 {
		a.ID = ser.NewKey()
	} else {
		a.ID = key
	}
	a.Description = from["Description"][0]
	a.Keywords = from["Keywords"][0]
	a.Server, _ = strconv.Atoi(from["Server"][0])
	a.Slogan = from["Slogan"][0]
	a.Template, _ = strconv.Atoi(from["Template"][0])
	a.Title = from["Title"][0]
	a.Url = from["Url"][0]
	return a
}
func (ser Rubrics) NewFromForm(from map[string][]string) gobzip.Serial {
	a := new(Rubric)
	key :=View.KeyFromForm(from)
	if key == 0 {
		a.ID = ser.NewKey()
	} else {
		a.ID = key
	}
	a.Blog, _ = strconv.Atoi(from["Blog"][0])
	a.Description = from["Description"][0]
	a.Keywords = from["Keywords"][0]
	a.Title = from["Title"][0]
	a.Url = from["Url"][0]
	return a
}
func (ser Globals) NewFromForm(from map[string][]string) gobzip.Serial {
	return nil
}
func (ser Resources) NewFromForm(from map[string][]string) gobzip.Serial {
	return nil
}
func (ser Themes) NewFromForm(from map[string][]string) gobzip.Serial {
	return nil
}

func (p *Page) Delegate(kind string) gobzip.Serializer {
	switch kind {
	case "articles":
		return p.Articles
	case "blogs":
		return p.Blogs
	case "globals":
		return p.Globals
	case "resources":
		return p.Resources
	case "rubrics":
		return p.Rubrics
	case "themes":
		return p.Themes
	}
	return nil
}
func (p *Page) KeyFromForm(from map[string][]string) int {
	if from["ID"] !=nil{
		key, err := strconv.Atoi(from["ID"][0])
		if err != nil {
			return 0
		}
		return key
	}
	return 0
}

/*type Server struct {
	ID     int
	IP     string
	Vendor string
	Cpu    string
	Cache  string
	Memory string
}*/

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
	Url         string
	Keywords    string
	Description string
	Blog        int
}

type Article struct {
	ID          int
	Date        string
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
	ID       int
	Name     string
	Template int
	Data     []byte
}

type Global struct {
	ID   int
	Name string
	Data []byte
}

type Page struct {
	HeadMeta  string
	Rubrics   Rubrics
	Articles  Articles
	Blogs     Blogs
	Blog      int
	Themes    Themes
	Resources Resources
	Globals   Globals
	//Servers   Servers
	Index     int
	Rubric    int
	Article   int
	Server    string
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

func LoadAll() os.Error {
	/*err := master.LoadKind(&View.Articles)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Blogs)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Globals)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Resources)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Rubrics)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Servers)
	if err != nil {
		return err
	}
	err = master.LoadKind(&View.Themes)
	if err != nil {
		return err
	}*/
	return nil
}