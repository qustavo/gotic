package fs

import "io/ioutil"

type file struct {
	data []byte
}

var files map[string]*file

func init() {
	files = make(map[string]*file)
}

func ReadFile(filename string) ([]byte, error) {
	if file, ok := files[filename]; ok {
		return file.data, nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	files[filename] = &file{data: data}
	return data, nil
}

func Add(filename string, data []byte) {
	file := &file{data: data}
	files[filename] = file
}