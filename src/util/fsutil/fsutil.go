// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package fsutil

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Mkdir(path string) error {
	perm := os.FileMode(0700)
	err := os.MkdirAll(path, perm)
	if err != nil {
		return err
	}
	return nil
}

func FilesListing(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// https://stackoverflow.com/a/47028486
	sort.Slice(entries, func(i, j int) bool {
		e1, _ := entries[i].Info()
		e2, _ := entries[j].Info()
		// sorting descending by modified time
		return e1.ModTime().After(e2.ModTime())
	})

	list := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			list = append(list, entry.Name())
		}
	}
	return list, nil
}

func WriteStreamToFile(path, filename string, stream io.Reader) error {
	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}
	return nil
}

func Sha256sum(stream io.Reader) (string, error) {
	hash := sha256.New()
	_, err := io.Copy(hash, stream)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func ResolveFileConflict(path, file string) string {
	exists := func(file string) bool {
		_, err := os.Stat(file)
		return err == nil
	}

	// absolute path
	fpath := filepath.Join(path, file)
	if !exists(fpath) {
		return file
	}

	next := 1
	ext := filepath.Ext(file)
	name := strings.TrimSuffix(file, ext)
	for {
		fname := fmt.Sprintf("%s(%d)%s", name, next, ext)
		fpath = filepath.Join(path, fname)
		if !exists(fpath) {
			return fname
		}
		next++
	}
}
