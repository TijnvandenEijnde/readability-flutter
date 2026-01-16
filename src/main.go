package main

/*
#include <stdlib.h>

typedef struct {
	char* title;
	char* author;
	int length;
	char* excerpt;
	char* site_name;
	char* image_url;
	char* favicon_url;
	char* content; // HTML content
	char* text_content; // text content
	char* language;
	char* published_time;
	char* err;
	int success;
} CArticle;
*/
import "C"
import (
	"bytes"
	"strings"
	"time"
	"unsafe"

    "codeberg.org/readeck/go-readability/v2"
)

//export Parse
func Parse(url *C.char) (result C.CArticle) {
	goURL := C.GoString(url)
	article, err := readability.FromURL(goURL, 30*time.Second)
	if err != nil {
		return C.CArticle{
			err: C.CString(err.Error()),
			success: 0,
		}
	}

    var htmlBuf bytes.Buffer
    if err := article.RenderHTML(&htmlBuf); err != nil {
        return C.CArticle{
           err: C.CString(err.Error()),
           success: 0,
        }
    }
    htmlContent := htmlBuf.String()

    var textBuf bytes.Buffer
    if err := article.RenderText(&textBuf); err != nil {
        return C.CArticle{
           err: C.CString(err.Error()),
           success: 0,
        }
    }
    textContent := textBuf.String()

    wordCount := len(strings.Fields(textContent))

    publishedTime := ""
    if pt, err := article.PublishedTime(); err == nil {
        publishedTime = pt.Format(time.RFC3339)
    }

    return C.CArticle{
       title:          C.CString(article.Title()),
       author:         C.CString(article.Byline()),
       length:         C.int(wordCount),
       excerpt:        C.CString(article.Excerpt()),
       site_name:      C.CString(article.SiteName()),
       image_url:      C.CString(article.ImageURL()),
       favicon_url:    C.CString(article.Favicon()),
       content:        C.CString(htmlContent),
       text_content:   C.CString(textContent),
       language:       C.CString(article.Language()),
       published_time: C.CString(publishedTime),
       err:            nil,
       success:        1,
    }
}

//export FreeArticle
func FreeArticle(article C.CArticle) {
    if article.title != nil {
        C.free(unsafe.Pointer(article.title))
    }
    if article.author != nil {
        C.free(unsafe.Pointer(article.author))
    }
    if article.excerpt != nil {
        C.free(unsafe.Pointer(article.excerpt))
    }
    if article.site_name != nil {
        C.free(unsafe.Pointer(article.site_name))
    }
    if article.image_url != nil {
        C.free(unsafe.Pointer(article.image_url))
    }
    if article.favicon_url != nil {
        C.free(unsafe.Pointer(article.favicon_url))
    }
    if article.content != nil {
        C.free(unsafe.Pointer(article.content))
    }
    if article.text_content != nil {
        C.free(unsafe.Pointer(article.text_content))
    }
    if article.language != nil {
        C.free(unsafe.Pointer(article.language))
    }
    if article.published_time != nil {
        C.free(unsafe.Pointer(article.published_time))
    }
}

//export enforce_binding
func enforce_binding() {}

func main() {

}