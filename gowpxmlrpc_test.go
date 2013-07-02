package wordpress

import (
	"testing"
)

const (
	URL      = "http://blog.qortex.cn/xmlrpc.php"
	USERNAME = "name"
	PASSWORD = "pass"
)

func TestGetCategories(t *testing.T) {
	options := make(map[string]interface{})
	ba := &BlogAccount{
		Url:      URL,
		UserName: USERNAME,
		PassWord: PASSWORD,
		BlogId:   "0",
	}
	cats := GetCategories(ba, options)
	for _, cat := range cats {
		t.Log(cat.Id)
		t.Log(cat.Name)
	}
}

func TestGetAuthors(t *testing.T) {
	options := make(map[string]interface{})
	ba := &BlogAccount{
		Url:      URL,
		UserName: USERNAME,
		PassWord: PASSWORD,
		BlogId:   "0",
	}
	authors := GetAuthors(ba, options)
	for _, author := range authors {
		t.Log(author.Id)
		t.Log(author.Name)
	}
}

func TestGetTags(t *testing.T) {
	options := make(map[string]interface{})
	ba := &BlogAccount{
		Url:      URL,
		UserName: USERNAME,
		PassWord: PASSWORD,
		BlogId:   "0",
	}
	tags := GetTags(ba, options)
	for _, tag := range tags {
		t.Log(tag.Id)
		t.Log(tag.Name)
	}
}

func TestNewPost(t *testing.T) {
	options := make(map[string]interface{})
	options["post_title"] = "a啊啊aAAAAA"
	options["post_author"] = 2
	options["post_content"] = "post_cont en aaaaaa a aaaaaaaaa 啊啊t"
	options["post_type"] = "post"
	options["post_status"] = "publish"
	options["terms"] = map[string]interface{}{"category": map[string]interface{}{"0": "1", "1": "2"}}
	options["terms_names"] = map[string]interface{}{"post_tag": map[string]interface{}{"0": "tag1", "1": "tag2"}}

	ba := &BlogAccount{
		Url:      URL,
		UserName: USERNAME,
		PassWord: PASSWORD,
		BlogId:   "0",
	}

	id := NewPost(ba, options)
	t.Log(id)
}

func TestUploadFile(t *testing.T) {
	options := make(map[string]interface{})
	options["data"] = map[string]interface{}{"name": "Filename", "bits": ""}

	ba := &BlogAccount{
		Url:      URL,
		UserName: USERNAME,
		PassWord: PASSWORD,
		BlogId:   "0",
	}

	id := NewPost(ba, options)
	t.Log(id)
}
