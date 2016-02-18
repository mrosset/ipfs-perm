package main

import (
	"fmt"
	"github.com/str1ngs/util/json"
	"log"
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
		mode := info.Mode()
		if mode == os.ModeSymlink {
			return nil
		}
		entries = append(entries, FileEntry{p, info.Mode()})
		return err
	}
	filepath.Walk(".", walkFn)
	err = json.Write(entries, "unix_perm.json")
	if err != nil {
		return err
	}
	os.Chdir("..")
	cmd := exec.Command("ipfs", "add", "-r", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Get(out string, hash string) error {
	cmd := exec.Command("ipfs", "get", hash)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	jf := filepath.Join(hash, "unix_perm.json")
	entries := []FileEntry{}

	err = json.Read(&entries, jf)
	if err != nil {
		return err
	}
	if err := os.Chdir(hash); err != nil {
		return err
	}
	for _, e := range entries {
		err = os.Chmod(e.Path, e.Mode)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
