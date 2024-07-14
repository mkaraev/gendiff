## Gendiff - JSON/YAML Diff Utility

### Overview
This command-line utility generates the difference between two JSON or YAML files. It helps you compare configurations, settings, or any structured data stored in JSON or YAML format, providing clear and concise output highlighting the differences.

This project is given to students as a learning project in the [Hexlet](https://ru.hexlet.io/) online platform

### Installation

```
go get github.com/mkaraev/gendiff-golang@v0.1.0
```


### Usage
```
Compares two configuration files and shows a difference.

Usage:
  gendiff [-h] [-f FORMAT] first_file second_file [flags]

Flags:
  -f, --format string   set format of output (default "stylish")
  -h, --help            help for gendiff
```

