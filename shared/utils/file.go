package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CheckFileExist(link string) ([]string, error) {
	files, err := filepath.Glob(link)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ReadFileFromSource(source string) ([]byte, error) {
	file, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFileZip(rc io.Reader) ([]byte, error) {
	fileCtn := new(bytes.Buffer)
	_, err := fileCtn.ReadFrom(rc)
	if err != nil {
		return nil, err
	}
	return fileCtn.Bytes(), nil
}

func ScanFolder(link string) ([]string, error) {
	var files []string
	_ = filepath.Walk(link, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, nil
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}
			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			_, err = io.Copy(outFile, rc)
			outFile.Close()
			if err != nil {
				return filenames, err
			}
		}
	}
	return filenames, nil
}

func WriteFiletoDisk(filename string, fileDir string, fileData []byte, fmode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(fileDir), os.ModePerm); err != nil {
		return err
	}
	// fmt.Println("Writing", fileDir, filename)
	outFile, err := os.OpenFile(fileDir+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fmode)
	if err != nil {
		return err
	}
	_, err = io.Copy(outFile, bytes.NewBuffer(fileData))
	if err != nil {
		return err
	}
	return nil
}
