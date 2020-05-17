package api

import (
	"log"
	"net/http"

	//"github.com/gocolly/colly"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	/*
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	url := r.PostForm.Get("url")
	if url == "" {
		log.Println("missing url argument")
		return
	}

	header := http.Header{}
	for k, v := range r.PostForm {
		if k == "url" {
			continue
		}
		for _, vv := range v {
			header.Set(k, vv)
		}
	}
	log.Printf("url: %s, header: %#v", url, header)

	c := colly.NewCollector()
	rHeader := http.Header{}
	statusCode := 0
	var body []byte
	var rErr error

	c.OnResponse(func(r *colly.Response) {
		rHeader = *r.Headers
		statusCode = r.StatusCode
		body = r.Body
	})

	c.OnError(func(r *colly.Response, err error) {
		statusCode = 499
		rErr = err
	})

	c.Visit(url)

	for k, v := range rHeader {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	w.WriteHeader(statusCode)
	if rErr != nil {
		w.Write([]byte(rErr.Error()))
	} else {
		w.Write(body)
	}

	*/
}
