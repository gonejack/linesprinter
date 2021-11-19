package linesprinter

import "io"

// LinesPrinter break string into fixed length lines.
// caller must call Close to flush remaining bytes and an extra ending sep after use.
type LinesPrinter struct {
	w   io.Writer
	n   int
	np  int
	len int
	cap int
	sep []byte
	dat []byte
}

func (x *LinesPrinter) Write(p []byte) (n int, err error) {
	for {
		if len(p) == 0 {
			break
		}
		d := x.n - x.len%x.np
		if len(p) < d {
			d = len(p)
		}
		x.len += copy(x.dat[x.len:], p[:d])
		p, n = p[d:], n+d
		if x.len%x.np == x.n {
			x.len += copy(x.dat[x.len:], x.sep)
		}
		if x.len == x.cap {
			_, e := x.w.Write(x.dat[:])
			if e != nil {
				return n, e
			}
			x.len = 0
		}
	}
	return
}

// Close the caller must call Close to flush any partially line, and ending sep.
func (x *LinesPrinter) Close() (e error) {
	if x.len > 0 {
		_, e = x.w.Write(x.dat[:x.len])
		_, e = x.w.Write(x.sep)
	}
	return
}

// NewLinesPrinter new a printer with line length limit and separator
func NewLinesPrinter(w io.Writer, lineLen int, sep []byte) *LinesPrinter {
	return NewLinesPrinterN(w, lineLen, 100, sep)
}

// NewLinesPrinterN new a printer with control of how many lines could stay in memory
// thus users can control internal buffer size.
func NewLinesPrinterN(w io.Writer, lineLen int, memLine int, sep []byte) *LinesPrinter {
	if lineLen < 1 {
		panic("lineLen < 1")
	}
	if memLine < 1 {
		panic("bufLen < 1")
	}
	p := &LinesPrinter{
		w:   w,
		n:   lineLen,
		np:  lineLen + len(sep),
		sep: sep,
	}
	p.cap = p.np * memLine
	p.dat = make([]byte, p.cap)
	return p
}
