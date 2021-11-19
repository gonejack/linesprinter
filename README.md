# linesprinter
golang library to print string into fixed length lines

### Usage

```shell
go get github.com/gonejack/linesprinter
```

Setting line length to 76 and sep as \r\n
```golang
func main() {
    p := linesprinter.NewLinesPrinter(os.Stdout, 76, []byte("\r\n"))
    defer p.Close()
    w := base64.NewEncoder(base64.StdEncoding, p)
    defer w.Close()
    
    t := "this is some random string this is some random string this is some random string"
    w.Write([]byte(t))
}
```

Output
```
dGhpcyBpcyBzb21lIHJhbmRvbSBzdHJpbmcgdGhpcyBpcyBzb21lIHJhbmRvbSBzdHJpbmcgdGhp
cyBpcyBzb21lIHJhbmRvbSBzdHJpbmc=
```
