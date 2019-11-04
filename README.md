# EnumGenerator



## About this repository

This is a Git repository for commands that convert `Localizable.strings` files, which are multilingual files used in Xcode, into enum definitions. The enum definition is output in swift syntax.

This command aims to replace RSwift string definitions with simpler ones.



## Build 

```
$ go build generator.go scanner.go
```

## Usage

```
$ generator "top directory for scan" "enum name" > "output faile name"
```

if you don't add any option as below. senumer will automatically choose current directory as top directory for scan, and "LocalizableStrings" as enum name.

```
$ generator > "output_directory"
```




