// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type cheatsheet struct {
	Title string `xml:"title,attr"`
	Intro intro  `xml:"intro"`
	Items []item `xml:"item"`
}

type intro struct {
	Description string `xml:"description"`
}

type item struct {
	Skip        bool    `xml:"skip,attr"`
	Label       string  `xml:"label,attr"`
	Title       string  `xml:"title,attr"`
	Description string  `xml:"description"`
	Action      action  `xml:"action"`
	Command     command `xml:"command"`
	Subitems    []item  `xml:"subitem"`
}

type command struct {
	Required      bool   `xml:"required,attr"`
	Returns       string `xml:"returns,attr"`
	Serialization string `xml:"serialization,attr"`
}

type action struct {
	PluginID string `xml:"pluginId,attr"`
	Class    string `xml:"class,attr"`
	Param1   string `xml:"param1,attr"`
	Param2   string `xml:"param2,attr"`
}

func main() {
	var file, out string
	flag.StringVar(&file, "file", "", "URL of cheatsheet XML")
	flag.StringVar(&out, "out", "", "Name of new HTML file (without extension)")
	flag.Parse()

	if len(file) > 0 {
		response, err := http.Get(file)
		if err != nil {
			log.Fatal(err)
		} else {
			defer response.Body.Close()
			xmlFile, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err.Error())
			}
			var cs cheatsheet
			xml.Unmarshal(xmlFile, &cs)
			fmt.Println(cs.Title)
			t, _ := template.ParseFiles("cheatsheet.html")
			var tpl bytes.Buffer
			if err := t.Execute(&tpl, cs); err != nil {
				log.Fatal(err)
			}

			cspath := filepath.Join(".", "cheatsheets")
			os.MkdirAll(cspath, os.ModePerm)

			ioutil.WriteFile(fmt.Sprintf("cheatsheets/%s", out), tpl.Bytes(), 0644)
		}
	} else {
		log.Fatal("Please enter file URL")
	}

}
