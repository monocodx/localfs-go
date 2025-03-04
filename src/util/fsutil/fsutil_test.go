// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package fsutil_test

import (
	"bytes"
	"localfs/util/fsutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMkdir(t *testing.T) {
	// initialize testcases
	tcs := []struct {
		data     map[string]string
		expected string
	}{
		{
			data: map[string]string{
				"path": t.TempDir(),
				"name": "localfs",
			},
			expected: "localfs",
		},
		{
			data: map[string]string{
				"path": "/",
				"name": "localfs",
			},
			expected: "",
		},
	}

	t.Run("Make Directory With Access Permission", func(t *testing.T) {
		tdata := tcs[0].data
		expected := tcs[0].expected
		path := filepath.Join(tdata["path"], tdata["name"])

		err := fsutil.Mkdir(path)
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}

		actual, err := os.Stat(path)
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}

		if !actual.IsDir() {
			t.Errorf("\nExpected 'is a directory', but path is not a directory.")
		}

		if actual.Name() != expected {
			t.Errorf("\nTest Data: (Dir: %s)\nExpected: %s\nActual: %s",
				path, expected, actual.Name())
		}
	})

	t.Run("Make Directory Without Access Permission", func(t *testing.T) {
		tdata := tcs[1].data
		path := filepath.Join(tdata["path"], tdata["name"])

		err := fsutil.Mkdir(path)
		if err == nil {
			t.Errorf("\nExpected 'Permission denied' error, but no error was thrown.")
			t.FailNow()
		}
	})
}

func TestFilesListing(t *testing.T) {
	// initialize testcases
	tcs := []struct {
		data     string
		expected []string
	}{
		{
			data:     t.TempDir(),
			expected: []string{"test_file_2", "test_file_1"},
		},
	}

	t.Run("Files Listing Sort By Latest File Creation Timestamp", func(t *testing.T) {
		// setup test data
		tdata := tcs[0].data
		expected := tcs[0].expected
		func() {
			files := make([]string, 2)
			copy(files, expected)
			sort.Strings(files)

			for _, file := range files {
				// create test data
				f, err := os.Create(filepath.Join(tdata, file))
				if err != nil {
					t.Errorf("\nError: %s", err)
					t.FailNow()
				}
				f.Close()
				// delay creation
				time.Sleep(100 * time.Millisecond)
			}
		}()

		actual, err := fsutil.FilesListing(tdata)
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}

		if len(actual) != 2 ||
			actual[0] != expected[0] ||
			actual[1] != expected[1] {
			t.Errorf("\nTest Data: (Path: %s)\nExpected: %v\nActual: %v",
				tdata, expected, actual)
		}
	})
}

func TestWriteStreamToFile(t *testing.T) {
	// t.TempDir returns a temporary directory for the test to use.
	// The directory is automatically removed when the test and
	// all its subtests complete.
	tempDir := t.TempDir()
	// initialize testcases
	tcs := []struct {
		data     []byte
		expected string
	}{
		{
			data:     []byte("Fuiyoh!!"),
			expected: "Fuiyoh!!",
		},
	}

	t.Run("Text File", func(t *testing.T) {
		// setup test data
		tdata := tcs[0].data
		expected := tcs[0].expected

		err := fsutil.WriteStreamToFile(tempDir, "tempfile", bytes.NewReader(tdata))
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}

		inbyte, err := os.ReadFile(filepath.Join(tempDir, "tempfile"))
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}

		actual := string(inbyte)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%v)\nExpected: %s\nActual: %s",
				tdata, expected, actual)
		}
	})
}

func TestSha256sum(t *testing.T) {
	// t.TempDir returns a temporary directory for the test to use.
	// The directory is automatically removed when the test and
	// all its subtests complete.
	tempDir := t.TempDir()
	// initialize testcases
	tcs := []struct {
		data     []byte
		expected string
	}{
		{
			data:     []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			expected: "2a3dfe6fbe56133c3254e7c3db3f70e3f706e8e9030ef82d416d77a18c904633",
		},
		{
			data:     []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ\n"),
			expected: "91f4cd16425ef1ff5f222ba266d0285775f4ec250290c394642ff2eedd1aafea",
		},
	}

	t.Run("Hash From Stream", func(t *testing.T) {
		tdata := tcs[0].data
		expected := tcs[0].expected

		stream := bytes.NewReader(tdata)
		actual, _ := fsutil.Sha256sum(stream)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%v)\nExpected: %s\nActual: %s",
				tdata, expected, actual)
		}
	})

	t.Run("Hash From File", func(t *testing.T) {
		// setup test data
		tempFile := filepath.Join(tempDir, "tempfile")
		tdata := tcs[1].data
		expected := tcs[1].expected
		func() {
			err := os.WriteFile(tempFile, tcs[1].data, 0755)
			if err != nil {
				t.Errorf("\nError: %s", err)
				t.FailNow()
			}
		}()

		file, err := os.Open(tempFile)
		if err != nil {
			t.Errorf("\nError: %s", err)
			t.FailNow()
		}
		defer file.Close()

		actual, _ := fsutil.Sha256sum(file)

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%v)\nExpected: %s\nActual: %s",
				tdata, expected, actual)
		}
	})
}

func TestResolveFileConflict(t *testing.T) {
	// t.TempDir returns a temporary directory for the test to use.
	// The directory is automatically removed when the test and
	// all its subtests complete.
	tempDir := t.TempDir()
	// initialize testcases
	tcs := []struct {
		data     []string
		expected string
	}{
		{
			data:     []string{"test_file"},
			expected: "test_file(1)",
		},
		{
			data:     []string{"test_file.pdf"},
			expected: "test_file(1).pdf",
		},
		{
			data:     []string{"test_file.pdf", "test_file(1).pdf", "test_file(2).pdf"},
			expected: "test_file(3).pdf",
		},
	}

	t.Run("File Without File Extension", func(t *testing.T) {
		// setup test data
		tdata := tcs[0].data
		expected := tcs[0].expected
		func() {
			for _, data := range tdata {
				file, err := os.Create(filepath.Join(tempDir, data))
				if err != nil {
					t.Errorf("\nError: %s", err)
					t.FailNow()
				}
				file.Close()
			}
		}()

		actual := fsutil.ResolveFileConflict(tempDir, tdata[0])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%s)\nExpected: %v\nActual: %v",
				tdata, filepath.Join(tempDir, expected), actual)
		}
	})

	t.Run("File With File Extension", func(t *testing.T) {
		// setup test data
		tdata := tcs[1].data
		expected := tcs[1].expected
		func() {
			for _, data := range tdata {
				file, err := os.Create(filepath.Join(tempDir, data))
				if err != nil {
					t.Errorf("\nError: %s", err)
					t.FailNow()
				}
				file.Close()
			}
		}()

		actual := fsutil.ResolveFileConflict(tempDir, tdata[0])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%s)\nExpected: %v\nActual: %v",
				tdata, filepath.Join(tempDir, expected), actual)
		}
	})

	t.Run("File With More Than One Occurrences", func(t *testing.T) {
		// setup test data
		tdata := tcs[2].data
		expected := tcs[2].expected
		func() {
			for _, data := range tdata {
				file, err := os.Create(filepath.Join(tempDir, data))
				if err != nil {
					t.Errorf("\nError: %s", err)
					t.FailNow()
				}
				file.Close()
			}
		}()

		actual := fsutil.ResolveFileConflict(tempDir, tdata[0])
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\nTest Data: (%s)\nExpected: %v\nActual: %v",
				tdata, filepath.Join(tempDir, expected), actual)
		}
	})
}
