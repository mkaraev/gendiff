## Gendiff - JSON/YAML Diff Utility

### Overview
This command-line utility generates the difference between two JSON or YAML files. It helps you compare configurations, settings, or any structured data stored in JSON or YAML format, providing clear and concise output highlighting the differences.

This project is given to students as a learning project in the [Hexlet](https://ru.hexlet.io/) online platform

### Installation

If you want to use it in your Go projects:
```
go get github.com/mkaraev/gendiff@v0.1.0
```

If you want to use it as command-line tool:
```
git clone git@github.com:mkaraev/gendiff.git
cd gendiff/
make build
./cmd/gendiff <path-to-file> <path-to-file>
```


### Usage
```
Compares two configuration files and shows the difference.

Usage:
  gendiff [-h] [-f FORMAT] first_file second_file [flags]

Flags:
  -f, --format string   set format of output (default "stylish")
  -h, --help            help for gendiff
```

