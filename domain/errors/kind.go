package errors

import "fmt"

// Kind categorizes the error.
type Kind uint

//go:generate stringer -type=Kind -linecomment
// List of error kinds.
const (
	Unknown            Kind = iota // unknown
	Internal                       // internal
	BadInput                       // bad_input
	NotFound                       // not_found
	AlreadyExists                  // already_exists
	FailedPrecondition             // failed_precondition
	Unauthorized                   // unauthorized
	Forbidden                      // forbidden
	kindEnd
)

func KindFromStr(s string) Kind {
	for i := 0; i < int(kindEnd); i++ {
		if s == Kind(i).String() {
			return Kind(i)
		}
	}
	panic(fmt.Sprintf("kind not found: %s", s))
}
