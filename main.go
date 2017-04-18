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
	"flag"
	"fmt"
	loader "github.com/KamiQuasi/cheatsheeter/content-loader"
	parser "github.com/KamiQuasi/cheatsheeter/parser"
	storage "github.com/KamiQuasi/cheatsheeter/storage"
	"log"
)

var (
	file, out string
	csStorage = storage.NewHtmlStorage()
)

func Init() {
	flag.StringVar(&file, "file", "", "URL of cheatsheet XML")
	flag.StringVar(&out, "out", "", "Name of new HTML file (without extension)")
	flag.Parse()
}

func main() {
	Init()

	xmlFileBytes := loader.LoadContent(file)

	result, parseError := parser.Parse(xmlFileBytes)
	if parseError != nil {
		log.Fatalf("Failed to parse docs. Cause: '%s'\n", parseError.Error())
	}

	fmt.Println("Content successfully parsed.")
	if out == "" {
		fmt.Println(string(result))
	} else {
		stg := csStorage.Store(result, out)
		fmt.Printf("Result saved by path '%s'\n", stg.Path)
	}
	fmt.Println("Application completed task.")
}
