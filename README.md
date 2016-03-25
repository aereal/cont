# cont(1)

cont(1) is a stream bookmark.

cont reads the standard input, write it to the standard output, and records read position.

Next time cont run, it skips the input to last unread position and writes to the output.

## Synopsis

```
cat very-huge-file | cont | grep ...
```

## Install

```sh
go get github.com/aereal/cont
cd $GOPATH/src/github.com/aereal/cont
go build ./...
```

## Author

aereal

## License

MIT
