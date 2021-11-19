package linesprinter

import "io"

// LinesPrinter break string into fixed length lines.
// caller must call Close to flush remaining bytes and an extra ending sep after use.
type LinesPrinter struct {
	output io.Writer

	sep []byte

	n struct {
		chars int
		full  int
	}

	b struct {
		dat []byte
		len int
		cap int
	}
}

func (pp *LinesPrinter) Write(p []byte) (n int, err error) {
	for {
		if len(p) == 0 {
			break
		}

		ncp := pp.n.chars - pp.b.len%pp.n.full
		if len(p) < ncp {
			ncp = len(p)
		}

		pp.b.len += copy(pp.b.dat[pp.b.len:], p[:ncp])
		n += ncp
		p = p[ncp:]

		if pp.b.len%pp.n.full == pp.n.chars {
			pp.b.len += copy(pp.b.dat[pp.b.len:], pp.sep)
		}

		if pp.b.len == pp.b.cap {
			_, e := pp.output.Write(pp.b.dat[:])
			if e != nil {
				return n, e
			}
			pp.b.len = 0
		}
	}
	return
}

// Close the caller must call Close to flush any partially line, and ending sep.
func (pp *LinesPrinter) Close() (e error) {
	if pp.b.len > 0 {
		_, e = pp.output.Write(pp.b.dat[:pp.b.len])
		_, e = pp.output.Write(pp.sep)
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
		output: w,
		sep:    sep,
	}
	p.n.chars = lineLen
	p.n.full = lineLen + len(sep)
	p.b.cap = p.n.full * memLine
	p.b.dat = make([]byte, p.b.cap)
	return p
}
