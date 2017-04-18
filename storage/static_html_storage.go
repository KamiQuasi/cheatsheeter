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

package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
	"github.com/KamiQuasi/cheatsheeter/utils"
)

const (
	/** Docs folder name  */
	DOCS_FOLDER = "cheatsheets";
	/** Permission for static html file */
	FILE_PERMISSION = 0644;
)

type htmlStorage struct {
	Path string
}

/** Create new Html storage */
func NewHtmlStorage() htmlStorage {
	storage := htmlStorage{Path: filepath.Join(utils.GetAppPath(), DOCS_FOLDER)}
	fmt.Println("Create storage folder by path: ", storage.Path)

	if err := os.MkdirAll(storage.Path, os.ModePerm); err != nil {
		log.Fatalf("Failed to create storage for docs. Cause: '%s'\n", err.Error())
	}
	return storage
}

/** Save static html file to storage. */
func (storage *htmlStorage) Store(result []byte, fileName string) *htmlStorage {
	target := filepath.Join(storage.Path, fileName)
	fmt.Printf("Save file to storage by path: '%s'\n", target)

	if err := ioutil.WriteFile(target, result, FILE_PERMISSION); err != nil {
		log.Fatalf("Failed to store docs. Cause: '%s'\n", err.Error())
	}
	return storage
}
