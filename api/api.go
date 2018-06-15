package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type IndexContents struct {
	Version  string
	Service  string
	Hostname string
	User     string
}

var pkg = &IndexContents{
	Version: "4.0.0",
}

func GetIndexContents() *IndexContents {
	u, err := ioutil.ReadFile("/config/user")
	if err != nil {
		pkg.User = os.Getenv("USER")
	} else {
		pkg.User = string(u)
	}

	pkg.Hostname, _ = os.Hostname()

	return pkg

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	response, err := json.MarshalIndent(pkg, "", "\t")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(response))
}

func Index(w http.ResponseWriter, r *http.Request) {
	content := GetIndexContents()
	textTemplate := strings.Join(
		[]string{
			"<html><head><title>DemoApplication</title></head><body>",
			"<h1>Hello, {{.User}}</h1></br>",
			"<ul>",
			"ServiceName: {{.Service}}</br>",
			"Hostname: {{.Hostname}}</br>",
			"Version: {{.Version}}",
			"</ul>",
			"</body></html>",
		}, "")
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	err = tmpl.Execute(w, content)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

}
