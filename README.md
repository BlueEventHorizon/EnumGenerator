# Localenum



## About this repository

This is a Git repository for commands that convert `Localizable.strings` files, which are multilingual files used in Xcode, into enum definitions. The enum definition is output in swift syntax.

This command aims to replace RSwift string definitions with simpler ones.



## Build 

```
$ go build localenum.go scanner.go
```

## Usage

```
$ localenum top_directory_for_scan enum_name > output_directory
```



## Cathage

