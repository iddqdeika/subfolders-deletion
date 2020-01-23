package main

import (
	"bufio"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args)<2{
		out(fmt.Errorf("dir not specified"))
	}
	sc := bufio.NewScanner(os.Stdin)
	defer sc.Scan()
	f, err := os.Open(os.Args[1])
	if err != nil{
		out(err)
	}
	fs, err := f.Readdir(0)
	if err != nil{
		panic(err)
	}
	if len(fs) == 0 {
		fmt.Printf("no files in dir: %v", os.Args[0])
		return
	}
		fmt.Printf("found %v files/dirs", len(fs))
	p := pb.New(len(fs))
	p.Start()
	for _, file := range fs {
		err := os.RemoveAll(filepath.Join(os.Args[1], file.Name()))
		if err != nil{
			continue
		}
		p.Add(1)
		p.Update()
	}
	p.Finish()
	fmt.Printf("finished")
}

func out(err error){
	fmt.Printf("error occured: %v", err)
}