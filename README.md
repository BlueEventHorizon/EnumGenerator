# EnumGenerator



## About this repository

This is a Git repository for commands that convert `Localizable.strings` files, which are multilingual files used in Xcode, into enum definitions. The enum definition is output in swift syntax.

This command aims to replace R.Swift string definitions with simpler ones.

## Import

```
$ go get -u cloud.google.com/go/translate
$ go get -u golang.org/x/text/language
$ go get -u google.golang.org/api/option
```

## Credential

If you'd like to use Google Translate API, 
Prelase set the environment variable GOOGLE_APPLICATION_CREDENTIALS to the file path of the JSON file that contains the Google service account key.

## Build 

```
$ go build enumgenerator.go scanner.go
```

## Usage

```
$ enumgenerator "top directory for scan" "enum name" > "output faile name"
```

if you don't add any option as below. senumer will automatically choose current directory as top directory for scan, and "LocalizableStrings" as enum name.

```
$ enumgenerator > "output_directory"
```




