package main

import (
	"fmt"
	"github.com/str1ngs/util/json"
	"os"
	"os/exec"
	"path/filepath"
)

type FileEntry struct {
	Path string
	Mode os.FileMode
}

func Add(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	wd, _ := os.Getwd()
	fmt.Println("entering ", wd)
	entries := []FileEntry{}
	walkFn := func(p string, info os.FileInfo, err error) error {
		if p == "." {
			return nil
		}
		entries = append(entries, FileEntry{p, info.Mode()})
		return err
	}
	filepath.Walk(".", walkFn)
	for _, e := range entries {
		fmt.Println(e.Path, e.Mode)
	}
	err = json.Write(entries, "unix_perm.json")
	if err != nil {
		return err
	}
	os.Chdir("..")
	cmd := exec.Command("ipfs", "add", "-n", "-r", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
