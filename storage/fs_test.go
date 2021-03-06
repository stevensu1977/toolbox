package storage

import (
	"testing"
)

func TestFS_IsExits(t *testing.T) {
	path := "../src"
	if IsExit(path) {
		t.Logf("path : [%s] exits", path)
	} else {
		t.Fatalf("path [%s] not exits", path)
	}
}

func TestFS_Abs(t *testing.T) {
	path := "../src"
	fullpath, err := Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fullpath)
}

func TestFS_MkdirAll(t *testing.T) {
	path := "./tmp/123/测试"
	err := MkdirAll(path)
	if err != nil {
		t.Fatal(err)
	}
}
