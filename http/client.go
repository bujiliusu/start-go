package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world1\n")
}

type myReader struct {
	body string
}

func (m *myReader) Read(p []byte) (n int, err error) {
	fmt.Printf("%s", n)
	return n, nil
}

func main(){
	http.HandleFunc("/hello", HelloServer)
	//log.Fatal(http.ListenAndServe(":12345", nil))
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Mobile/15E148 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect", request)
			return nil
		},
	}
	//resp, err := http.DefaultClient.Do(request)
    resp , err := client.Do(request)
	//resp, err := http.Get("http://www.baidu.com")

	fmt.Printf("%T", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", data)
}
