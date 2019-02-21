# Package ddir

Package ddir provides functions for managing the application's data directory.

```go
import "github.com/godump/ddir"
```

- [Example](#Example)
- [Auto](#Auto)

# Example

`Base` sets base data path.

```go
ddir.Base("/tmp/play")
```

Create `/tmp/play/foo` if it doesn't exists.

```go
ddir.Make("foo")
```

Get absolute path `/tmp/play/bar` by given elems:

```go
name = ddir.Join("bar")
...
```

# Auto

Auto is an automatic Base function call affected by the operating system. Most applications' data directories follow the operating system's specifications, for example, the data directory of vim is placed in ~/.vim.

```go
ddir.Auto("Play")
```

- Equals with `Base("~/AppData/Local/Play")` on windows
- Equals with `Base("~/.play")` on linux
