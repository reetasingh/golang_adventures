package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
	"unsafe"
)

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	addr := uintptr(unsafe.Pointer(b))
	// Write the address into the buffer itself
	b.Reset()
	b.WriteString(fmt.Sprintf("%v", addr))

	// Replace this with time.Now() in a real logger.
	// b.WriteString(timeNow().UTC().Format(time.RFC3339))
	// b.WriteByte(' ')
	// b.WriteString(key)
	// b.WriteByte('=')
	// b.WriteString(val)
	b.WriteString("\n")
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
	Log(os.Stdout, "path", "/search?q=tree")
	Log(os.Stdout, "path", "/search?q=abc")
	Log(os.Stdout, "path", "/search?q=pqr")
}
