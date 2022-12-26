# wtft - What's the file type?

wtft is a simple command line tool that helps you figure out what the real file type is. It does this by looking at the file's magic number and comparing it to a list of known magic numbers. 

## Build and Run
pre-requisites: go 1.19 or higher

To build the cli: `make build-cli` or `go build -o wtft ./main.go`. [Here is a guide](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04) to output to other platforms

Example usage:

`wtft -file /path/to/png/file.png` will output the following json:
```
{
    "name": "PNG",
    "ext": "png",
    "mime_type": "image/png",
    "supported_extensions": [
        "png"
    ]
}
```

Example error: 
```
$ wtft -file /path/to/unknown/file
No file type could be found for file. File may be corrupt or signature does not exist in wtft codebase yet...
```
All errors will exit with code 1