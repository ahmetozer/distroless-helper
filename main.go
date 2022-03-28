package main

import (
	"debug/elf"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/u-root/u-root/pkg/ldd"
)

func check(e error) {
	if e != nil {
		log.Fatalf("%s", e)
	}
}

var (
	ELFFile     *elf.File
	ELFFilePath string
)

func main() {

	var err error
	// check os args
	if len(os.Args) != 3 {
		err = fmt.Errorf("please use like as " + os.Args[0] + " /bin/bash /target")
	}
	check(err)

	ELFFilePath = os.Args[1]
	// Open binnary
	f, err := os.Open(ELFFilePath)
	check(err)

	// ELF file underlying reader
	ELFFile, err = elf.NewFile(f)
	check(err)

	// Read and decode ELF identifier
	var ident [16]uint8
	f.ReadAt(ident[0:], 0)
	check(err)

	if ident[0] != '\x7f' || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' {
		err = fmt.Errorf("Bad magic number at %d\n", ident[0:4])
	}
	check(err)

	log.Printf("Class: %s\n", ELFFile.Class.String())
	log.Printf("Machine: %s\n", ELFFile.Machine.String())

	list, err := ldd.Ldd([]string{os.Args[1]})

	for i := range list {
		targetFile := os.Args[2] + list[i].FullName
		targetFile = strings.Replace(targetFile, "//", "/", -1)

		log.Printf("file %s is copied to %s\n", list[i].FullName, targetFile)

		check(os.MkdirAll(filepath.Dir(targetFile), os.ModePerm))

		check(cp(list[i].FullName, targetFile))

	}

}

func cp(src, dst string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := dstFile.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, srcStat.Mode())
	if err != nil {
		return err
	}

	return
}
