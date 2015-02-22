# gotic
Converts files files into go code.

# Concept
Gotic is both, a library and an utility.

The library `gotic/fs`wraps `ioutil.ReadFile()`.

The utility, `gotic`, takes one or more files and genaraates go code representing
them (as one or more `[]byte{ ... }`)

# Install
As usual: `go get github.com/gchaincl/gotic`

# Usage

Let's say you have a program (`main.go`) that reads af file:

```go
package main

import "github.com/gchaincl/gotic/fs"

func main() {
  f, err := fs.ReadFile("some/file")
  if err != nil {
    panic(err)
  }
  
  println(string(f))
}

```

Now, to embed `some/file` into your code, just run:
```bash
$ gotic some/file > main_gotic.go
```

Generated `main_gotic.go` will look like:

```go
package main

import "github.com/gchaincl/gotic/fs"

func init() {
	fs.Add("some/file", []byte{ 
		0x70,0x72,0x6F,0x63,0x65, ...
	})
}
```
