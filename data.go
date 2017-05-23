package technicolor

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type ANSISequence []byte

type ANSISequences []ANSISequence

// "Flatten" the ANSI Sequences into a single byte stream
func (a ANSISequences) Sequence() []byte {
	out := []byte{}
	for _, seq := range a {
		out = append(out, seq...)
	}
	return out
}

func NewANSISequence(arguments []string, ansiType byte) ANSISequence {
	return ANSISequence(append(append([]byte{0x1b, 0x5b},
		[]byte(strings.Join(arguments, ";"))...,
	), ansiType))
}

func NewTerminalWriter(file *os.File) Writer {
	output := NewWriter(file)
	if IsFileTerminal(file) {
		output = output.EnableColor()
	} else {
		output = output.DisableColor()
	}
	return output
}

func NewWriter(o io.Writer) Writer {
	return Writer{
		aNSISequences: ANSISequences{},
		output:        o,
		enableColor:   false,
	}
}

type Writer struct {
	aNSISequences ANSISequences
	output        io.Writer
	enableColor   bool
}

func (w Writer) DisableColor() Writer {
	return Writer{
		aNSISequences: w.aNSISequences,
		output:        w.output,
		enableColor:   false,
	}
}

func (w Writer) EnableColor() Writer {
	return Writer{
		aNSISequences: w.aNSISequences,
		output:        w.output,
		enableColor:   true,
	}
}

func (w Writer) Bold() Writer {
	return w.Add(NewANSISequence([]string{"1"}, 'm'))
}

func (w Writer) Add(seqs ...ANSISequence) Writer {
	return Writer{
		aNSISequences: append(w.aNSISequences, seqs...),
		output:        w.output,
		enableColor:   w.enableColor,
	}
}

func (w Writer) Reset(seqs ...ANSISequence) Writer {
	return Writer{
		aNSISequences: ANSISequences{},
		output:        w.output,
		enableColor:   w.enableColor,
	}
}

func (w Writer) Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, args...)
}

func (w Writer) Write(out []byte) (int, error) {
	if w.enableColor {
		return w.output.Write(append(w.aNSISequences.Sequence(), out...))
	}
	return w.output.Write(out)
}

// This is some yanked nasty shaz that will check to see if a file descriptor
// we have is a terminal or something else. This is handy for color stuff,
// since we want to be able to pipe things without seeing colors in the
// processing pipeline, but also make things sane for a user.
func IsTerminal(fd uintptr) bool {
	var termios syscall.Termios
	var ioctlReadTermios uintptr = 0x5401 // syscall.TCGETS
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		fd,
		ioctlReadTermios,
		uintptr(unsafe.Pointer(&termios)),
		0, 0, 0,
	)
	return err == 0
}

// Check to see if a given os.File is a Terminal. Notibly, `os.Stdout` is
// one of these, as is `os.Stderr`.
func IsFileTerminal(file *os.File) bool {
	return IsTerminal(file.Fd())
}
