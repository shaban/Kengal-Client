package main

import (
	"os"
	gobzip "github.com/shaban/kengal/gobzip"
	"time"
)

type Articles []*Article
type Rubrics []*Rubric
type Blogs []*Blog
type Themes []*Theme
type Resources []*Resource
type Globals []*Global

func (ser Articles)Len()int{
	return len(ser)
}
func (ser Articles)Less(i, j int) bool{
	it,_ := time.Parse("02.01.2006 15:04:05",ser[i].Date)
	jt,_ := time.Parse("02.01.2006 15:04:05",ser[j].Date)
	return jt.Seconds() < it.Seconds()
}
func (ser Articles)Swap(i, j int){
	cycle := make([]*Article,1)
	copy(cycle,ser[i:i+1])
	ser[i] = ser[j]
	ser[j]=cycle[0]
}
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