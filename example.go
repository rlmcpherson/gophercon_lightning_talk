// +build OMIT
package main

import (
	"io"
	"log"
	"os"

	"github.com/rlmcpherson/s3gof3r"
)

const (
	bucket_name = "rm-dev-repos"
	file_name   = "profiling_optimization.slide"
)

func main() {

	s3gof3r.SetLogger(os.Stdout, "", log.LstdFlags, false)

	// STARTEXAMPLE OMIT
	presentation, err := os.Open(file_name) // open presentation file
	if err != nil {
		log.Fatal(err)
	}

	k, err := s3gof3r.EnvKeys() // get S3 keys from environment
	if err != nil {
		log.Fatal(err)
	}
	// Open bucket to put file into
	s3 := s3gof3r.New("", k)
	b := s3.Bucket(bucket_name)

	// Open a PutWriter for upload
	w, err := b.PutWriter(presentation.Name(), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	if _, err = io.Copy(w, presentation); err != nil { // Copy into S3
		log.Fatal(err)
	}
	log.Printf("%s uploaded to %s", file_name, bucket_name)
	// STOPEXAMPLE OMIT
}
