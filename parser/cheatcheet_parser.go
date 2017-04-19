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

package cheatsheeter

import (
	"bytes"
	"encoding/xml"
	"fmt"
        "html/template"
	"path"
	"github.com/KamiQuasi/cheatsheeter/utils"
	"os"
	"log"
)

const (
	TEMPLATE_NAME = "cheatsheet.html";
)

type cheatsheet struct {
	Title string `xml:"title,attr"`
	Intro intro  `xml:"intro"`
	Items []item `xml:"item"`
}

type intro struct {
	Description htmlValue `xml:"description"`
}

//workaround to disable skipping html tags.
type htmlValue struct {
	Value template.HTML `xml:",innerxml"`
}

type item struct {
	Skip        bool      `xml:"skip,attr"`
	Label       string    `xml:"label,attr"`
	Title       string    `xml:"title,attr"`
	Description htmlValue `xml:"description"`
	Action      action    `xml:"action"`
	Command     command   `xml:"command"`
	Subitems    []item    `xml:"subitem"`
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

/**
 * Parse bytes from ".cheatsheet.xml" to bytes for create static html. Static html structure defined by html template file.
 * Html template should be located in the same directory with executable binary file.
 */
func Parse(xmlFileBytes []byte) ([]byte, error) {
	var cs cheatsheet;
	if marshalErr := xml.Unmarshal(xmlFileBytes, &cs); marshalErr != nil {
		return nil, marshalErr
	}
	fmt.Println(cs.Title)

	t, parseErr := template.ParseFiles(getTemplatePath());
	if parseErr != nil {
		return nil, parseErr
	}

	var tpl bytes.Buffer
	if executeErr := t.Execute(&tpl, cs); executeErr != nil {
		return nil, executeErr
	}

	return tpl.Bytes(), nil
}

/** Get template path. Notice: For this application we should store html template with binary file in the same folder. */
func getTemplatePath() string  {
	templatePath := path.Join(utils.GetAppPath(), TEMPLATE_NAME)
	if _, statErr := os.Stat(templatePath); statErr != nil  {
		// Calculate path for "go run" command
		if os.IsNotExist(statErr) {
			rootPath, wdErr := os.Getwd();
			if  wdErr != nil {
				log.Fatalf("Failed to get templatePath. Cause: '%s'\n", wdErr.Error())
			}
			templatePath = path.Join(rootPath, TEMPLATE_NAME)
		}
	}
	return templatePath
}
