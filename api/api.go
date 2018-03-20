package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/2ndWatch/golibs/version"
)

type IndexContents struct {
	Version  string
	Service  string
	Hostname string
}

func GetIndexContents(pkg version.PackageInfo) IndexContents {
	h, _ := os.Hostname()
	return IndexContents{
		version.Pkg.Version,
		version.Pkg.Message,
		h,
	}

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	response, err := json.MarshalIndent(version.Pkg, "", "\t")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(response))
}

func Index(w http.ResponseWriter, r *http.Request) {
	content := GetIndexContents(version.Pkg)
	textTemplate := strings.Join(
		[]string{
			"<html><head><title>DemoApplication</title></head><body>",
			"<h1>ServiceName: {{.Service}}</h1></br>",
			"<ul>",
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
