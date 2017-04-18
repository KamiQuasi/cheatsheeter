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
package content_loader

import (
	"io/ioutil"
	"log"
	"net/http"
)

/* Load ".cheatsheet.xml" content. */
func LoadContent(resource string) []byte {
	if len(resource) == 0 {
		log.Fatal("Please enter file URL\n")
	}

	var response *http.Response; var err error
	if response, err = http.Get(resource); err != nil {
		log.Fatalf("Failed to get resource: '%s'. Cause: '%s'\n", resource, err.Error())
	}
	defer response.Body.Close()

	var xmlFileBytes []byte
	if xmlFileBytes, err = ioutil.ReadAll(response.Body); err != nil {
		panic("Failed to read response body: " + err.Error())
	}
	return xmlFileBytes
}
