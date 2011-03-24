package main

import (
	"os"
	"fmt"
	"json"
	"io/ioutil"
	"io"
	"compress/gzip"
	"http"
)

type Serializer interface {
	New() interface{}
	Insert(insert interface{}) interface{}
}
const DB_ROOT = "db"

func saveItem(kind string, item interface{}, key int) os.Error {
	data, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%v.json", DB_ROOT,kind, key), data, 0666)
	if err != nil {
		return err
	}
	return nil
}
func deleteItem(kind string, id int) os.Error {
	err := os.Remove(fmt.Sprintf("%s/%s/%v.json", DB_ROOT,kind, id))
	return err
}
func updateBlog(b *Blog) os.Error {
	err := saveItem("blogs", b, b.ID)
	if err != nil {
		return err
	}
	View.Blogs.Replace(b)
	return nil
}

func updateRubric(rb *Rubric) os.Error {
	err := saveItem("rubrics", rb,rb.ID)
	if err != nil {
		return err
	}
	View.Rubrics.Replace(rb)
	return nil
}

func updateArticle(a *Article) os.Error {
	err := saveItem("articles", a, a.ID)
	if err != nil {
		return err
	}
	View.Articles.Replace(a)
	return nil
}

func insertBlog(b *Blog) os.Error {
	err := saveItem("blogs", b, b.ID)
	if err != nil {
		return err
	}
	View.Blogs = append(View.Blogs, b)
	return nil
}

func insertRubric(rb *Rubric) os.Error {
	err := saveItem("rubrics", rb, rb.ID)
	if err != nil {
		return err
	}
	View.Rubrics = append(View.Rubrics, rb)
	return nil
}

func insertArticle(a *Article) os.Error {
	err := saveItem("articles", a, a.ID)
	if err != nil {
		return err
	}
	View.Articles = append(View.Articles, a)
	return nil
}
func deleteRubric(id int) os.Error {
	s := make([]*Rubric, 0)
	for k, v := range View.Rubrics {
		if v.ID == id {
			err := deleteItem("rubrics", v.ID)
			if err != nil {
				return err
			}

		} else {
			s = append(s, View.Rubrics[k])
		}
	}
	View.Rubrics = s
	return nil
}

func deleteArticle(id int) os.Error {
	s := make([]*Article, 0)
	for k, v := range View.Articles {
		if v.ID == id {
			err := deleteItem("articles", v.ID)
			if err != nil {
				return err
			}
		} else {
			s = append(s, View.Articles[k])
		}
	}
	View.Articles = s
	return nil
}
func objFromGz(r io.ReadCloser, obj interface{})(os.Error){
	defer r.Close()
	gz, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(gz)
	err = decoder.Decode(&obj)
	if err != nil {
		return err
	}
	gz.Close()
	return nil
}
func objFromMaster(url string, scheme interface{})os.Error{
	r, _, err := http.Get(url)
	if err!= nil{
		return err
	}
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(gz)
	err = decoder.Decode(scheme)
	if err != nil {
		return err
	}
	gz.Close()
	return nil
}
func audit()(os.Error){
	err := objFromMaster("http://k-dany.de/admin/audit/articles",&View.Articles)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/blogs",&View.Blogs)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/globals",&View.Globals)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/resources",&View.Resources)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/servers",&View.Servers)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/themes",&View.Themes)
	if err != nil {
		return err
	}
	err = objFromMaster("http://k-dany.de/admin/audit/rubrics",&View.Rubrics)
	if err != nil {
		return err
	}
	return nil
}
