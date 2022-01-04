package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Pkg struct {
	ImportPath string   `json:"ImportPath"`
	Name       string   `json:"Name"`
	Deps       []string `json:"Deps"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: main PKG_NAME")
		return
	}

	pkgs, err := execGoList(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}

	all, err := execGoList("...")
	if err != nil {
		log.Fatal(err)
	}

	var paths = []string{}
	for _, p := range pkgs {
		paths = append(paths, p.ImportPath)
	}

	for _, v := range all {
		if contains(v.Deps, paths) {
			fmt.Println(v.ImportPath)
		}
	}
}

func execGoList(pkgName ...string) ([]Pkg, error) {
	args := []string{"list", "-json"}
	args = append(args, pkgName...)

	var stderr bytes.Buffer
	cmd := exec.Command("go", args...)
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error: %s", stderr.String())
	}

	var pkgs []Pkg
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var p Pkg
		if err := dec.Decode(&p); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, p)
	}

	return pkgs, nil
}

func contains(list []string, values []string) bool {
	for _, l := range list {
		for _, v := range values {
			if l == v {
				return true
			}
		}
	}
	return false
}
