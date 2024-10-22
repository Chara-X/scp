# scp

## Example

```go
func ExampleCopyFrom() {
	var f, _ = os.Open("./from.txt")
	defer f.Close()
	var c = &copy.Client{}
	c.Client, _ = ssh.Dial("tcp", addr, config)
	defer c.Close()
	c.CopyFrom(f, "/app/to.txt", "0655")
}
```

```go
func ExampleCopyTo() {
	var f, _ = os.Create("./from.txt")
	defer f.Close()
	var c = &copy.Client{}
	c.Client, _ = ssh.Dial("tcp", addr, config)
	defer c.Close()
	c.CopyTo(f, "/app/to.txt")
}
```
