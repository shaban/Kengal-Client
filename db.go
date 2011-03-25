package main

import (
	"os"
	"fmt"
	//"io"
	"compress/gzip"
	"http"
	"gob"
	"bytes"
)

type Serializer interface {
	New() interface{}
	Insert(insert interface{}) interface{}
}
const DB_ROOT = "db"

func saveItem(kind string, item interface{}, key int) os.Error {
	f, err := os.Open(fmt.Sprintf("%s/%s/%v.bin.gz", DB_ROOT,kind, key), os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bytes.NewBufferString("")
	
	genc := gob.NewEncoder(buf)
	err = genc.Encode(item)
	if err != nil {
		return err
	}
	gz, err := gzip.NewWriter(f)
	if err != nil {
		return err
	}
	_, err = gz.Write(buf.Bytes())
	if err != nil {
		return err
	}
	gz.Close()
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
func objFromStream(r *http.Request, scheme interface{})(os.Error){
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		return err
	}
	defer gz.Close()
	defer r.Body.Close()
	decoder := gob.NewDecoder(gz)
	err = decoder.Decode(scheme)
	if err != nil {
		return err
	}
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
	defer gz.Close()
	defer r.Body.Close()
	decoder := gob.NewDecoder(gz)
	err = decoder.Decode(scheme)
	if err != nil {
		return err
	}
	return nil
}
func audit()(os.Error){
	err := objFromMaster(fmt.Sprintf("http://%s/admin/audit/articles",View.Master),&View.Articles)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/blogs",View.Master),&View.Blogs)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/globals",View.Master),&View.Globals)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/resources",View.Master),&View.Resources)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/servers",View.Master),&View.Servers)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/themes",View.Master),&View.Themes)
	if err != nil {
		return err
	}
	err = objFromMaster(fmt.Sprintf("http://%s/admin/audit/rubrics",View.Master),&View.Rubrics)
	if err != nil {
		return err
	}
	return nil
}
