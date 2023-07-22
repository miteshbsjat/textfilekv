# textfilekv - Key-Value Store on Plain Text File

textfilekv is a pure Go based key-value store on plain text file as `=` delimited lines of text.
Each line represent a new record. This file can be even edited by hand when not in use to edit/update.


## Table of Contents

- [Getting Started](#getting-started)
  - [Installing](#installing)
  - [Usage](#usage)
  - [Testing](#testing)
- [Features](#features)

## Getting Started

### Installing

To start using textfilekv, install Go and run `go install`:

```sh
$ go install github.com/miteshbsjat/textfilekv@latest
```

### Usage

It is just required to provide path of text file to store the key-value(s).
This file will be created if not present.

```go
package main

import (
    "fmt"
    tkv "github.com/miteshbsjat/textfilekv"
)

func main() {
	filePath := "/tmp/data.txt"
	kvs := tkv.NewKeyValueStore(filePath)

	kvs.Set("name", "Mitesh")
	kvs.Set("age", "41")
	kvs.Set("data", "{\"height\": 167}")

	name, exists := kvs.Get("name")
	if exists {
		fmt.Printf("Name: %s\n", name)
	}

	age, exists := kvs.Get("age")
	if exists {
		fmt.Printf("Age: %s\n", age)
	}

	data, exists := kvs.Get("data")
	if exists {
		fmt.Printf("Data: %s\n", data)
	}
	demokey, exists := kvs.Get("demokey")
	if exists {
		fmt.Printf("Demokey: %s\n", demokey)
	}
}
```

Output

```bash
Name: Mitesh
Age: 41
Data: {"height": 167}
```

Now adding the key `demokey` by running following command after first run of this program.
```bash
echo 'demokey=demovalue' >> /tmp/data.txt
```
Output
```bash
Name: Mitesh
Age: 41
Data: {"height": 167}
Demokey: demovalue
```

### Testing

Testing can be done by running `go test`

```bash
go test
```

## Features

`textfilekv` has the following features for now.
* Simple to comprehend and update using *nix editors like. nano, vim, emacs, even filter programs
* Same format as `.env` file supported by config tools or libraries like `viper`
* Useful for programs which need only limited set of key-values stored in a file

This is quite simple text file based key value store written in go, there are quite a few areas for improvements and optimizations. 
