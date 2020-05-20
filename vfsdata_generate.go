// +build ignore

package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/shurcooL/vfsgen"
)

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func untar(dst string, r io.Reader) error {

	gzr, err := gzip.NewReader(r)

	if err != nil {
		return err
	}

	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func main() {
	version, err := os.Open("./version")
	if err != nil {
		log.Fatal(err)
	}
	defer version.Close()
	ver, err := ioutil.ReadAll(version)
	v, err := semver.NewVersion(fmt.Sprintf("%s", ver))
	if err != nil {
		log.Fatal(err)
	}
	latestVersion := v.String()
	archiveDir := "disposable-email-domains-" + latestVersion
	latestArchive := latestVersion + ".tar.gz"
	latestReleaseURL := "https://github.com/ivolo/disposable-email-domains/archive/" + latestArchive
	os.Mkdir("./build", 0755)
	resp, err := http.Get(latestReleaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("./build/" + latestArchive)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	archive, err := os.Open("./build/" + latestArchive)
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()
	err = untar("./build", archive)
	if err != nil {
		log.Fatal(err)
	}
	err = copy("./build/"+archiveDir+"/index.json", "./list/index.json")
	if err != nil {
		log.Fatal(err)
	}
	err = vfsgen.Generate(
		http.Dir("./list"),
		vfsgen.Options{
			Filename:     "./vfsdata_disposable.go",
			PackageName:  "disposable",
			VariableName: "asset",
		})
	if err != nil {
		log.Fatalln(err)
	}
}
