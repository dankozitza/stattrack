package stattrack

import (
	"github.com/dankozitza/seestack"
	"github.com/dankozitza/statdist"
	"strconv"
)

// this is the Stat struct from statdist
//type Stat struct {
//	Id int
//	Status string
//	ShortStack string
//	Message string
//	Stack string
//}

type StatTrack statdist.Stat

var pkgstat StatTrack = StatTrack{
	statdist.GetId(),
	"INIT",
	seestack.Short(),
	"package initialized",
	""}

func init() {
	statdist.Handle(statdist.Stat(pkgstat), true)
	return
}

func New(msg string) StatTrack {

	myst := StatTrack{
		statdist.GetId(),
		"INIT",
		seestack.ShortExclude(1),
		msg,
		""}

	statdist.Handle(myst.todist(), true)

	return myst
}

func (s StatTrack) Pass(m string) StatTrack {

	quiet := false
	if m == "" {
		m = "generic pass message"
		quiet = true
	}

	s.Status = "PASS"
	s.Message = m
	s.ShortStack = seestack.ShortExclude(1)
	statdist.Handle(s.todist(), quiet)
	return s
}

func (s StatTrack) Warn(m string) StatTrack {
	s.Status = "WARN"
	s.Message = m
	s.ShortStack = seestack.ShortExclude(1)
	statdist.Handle(s.todist(), false)
	return s
}

// ErrStatTrack
//
// This error object does not mention stattrack. It automatically sets the
// error string to display the attributes of the stattrack object. this
// way when s.Err("msg") is called the message does not need to mention
// it's package name.
//
type ErrStatTrack StatTrack

func (s ErrStatTrack) Error() string {

	return "\n[" + s.ShortStack + "][" + s.Status + "][" +
		strconv.Itoa(s.Id) + "] " + s.Message
}

func (s StatTrack) Err(m string) ErrStatTrack {
	s.Status = "ERROR"
	s.Message = m
	s.ShortStack = seestack.ShortExclude(1)
	s.Stack = seestack.Full()
	statdist.Handle(s.todist(), false)
	return ErrStatTrack(s)
}

func (s StatTrack) Panic(m string) {
	s.Status = "PANIC"
	s.Message = m
	s.ShortStack = seestack.ShortExclude(1)
	s.Stack = seestack.Full()
	statdist.Handle(s.todist(), false)
	panic(ErrStatTrack(s))
}

func (s StatTrack) PanicErr(m string, e error) {
	s.Status = "PANIC"
	s.Message = m + ": " + e.Error()
	s.ShortStack = seestack.ShortExclude(1)
	s.Stack = seestack.Full()
	statdist.Handle(s.todist(), false)
	panic(ErrStatTrack(s))
}

// converts StatTrack back to statdist.Stat
func (s StatTrack) todist() statdist.Stat {
	return statdist.Stat(s)
}
