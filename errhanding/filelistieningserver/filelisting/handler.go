package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

const prefix = "/list/"
func HandleFileList(writer http.ResponseWriter, request *http.Request) error{
	if strings.Index(request.URL.Path, prefix) != 0 {
		//fmt.Println(request.URL.Path)
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all)
	return nil
}
