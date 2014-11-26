package dirtar_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dockpit/dirtar"
)

func TestTar(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	dir, err := ioutil.TempDir("", "dirtar_t_")
	if err != nil {
		t.Fatal(err)
	}

	//tar
	b := bytes.NewBuffer(nil)
	err = dirtar.Tar(wd, b)
	if err != nil {
		t.Fatal(err)
	}

	//untar into tmp dir
	err = dirtar.Untar(dir, b)
	if err != nil {
		t.Fatal(err)
	}

	//should now be able to stat this file
	fi, err := os.Stat(filepath.Join(dir, "tar_test.go"))
	if err != nil {
		t.Fatal(err)
	}

	if fi.Size() == 0 {
		t.Fatal("Untarred files should not be empty")
	}
}
