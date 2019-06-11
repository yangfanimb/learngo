package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

func main() {
	var showProcess, force bool
	flag.BoolVar(&showProcess,"v", false,"show process")
	flag.BoolVar(&force,"f", false,"force to copy, ignore overwrite tips")
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		return
	}

	if err := copyFile(flag.Arg(0), flag.Arg(1), showProcess, force); err != nil {
		fmt.Printf("copy failed: %+v\n", err)
	}
}

func copyFile(src string, dst string, showProcess bool, forceWrite bool) error {
	if !fileExist(src) {
		return errors.Errorf("file %s is not exist", src)
	}

	if fileExist(dst) && !forceWrite {
		fmt.Println("over write it?y/n")
		reader := bufio.NewReader(os.Stdin)
		input, _, _ := reader.ReadLine()
		if strings.TrimSpace(string(input)) != "y" {
			return nil
		}
	}

	in, err := os.Open(src)
	if err != nil {
		return errors.WithMessagef(err, "open file %s failed", src)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return errors.WithMessagef(err, "open file %s failed", dst)
	}
	defer out.Close()


	_, err = io.Copy(out, in)
	if showProcess {
		fmt.Printf("copy %s --> %s\n", src, dst)
	}
	return err
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
