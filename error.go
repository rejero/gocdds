package gocdds

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lddsc
#include "ddsc/dds.h"
#include "ddsc/dds.h"
*/
import (
	"C"
	"math"
)

type CddsErrorType uint16

const DDS_XRETCODE_BASE = -50

const (
	Ok CddsErrorType = iota
	Error
	Unsupported
	BadParameter
	PreconditionNotMet
	OutOfResource
	NotEnabled
	ImmutablePolicy
	AlreadyDeleted
	TimeOut
	NoData
	IllegalOperation
	NotAllowedBySecurity
)

var retcodes = []string{
	"Success",
	"Non specific error",
	"Feature unsupported",
	"Bad parameter value",
	"Precondition for operation not met",
	"Out of resources",
	"Configurable feature is not enabled",
	"Attempt is made to modify an immutable policy",
	"Policy is used with inconsistent values",
	"Attempt is made to delete something more than once",
	"Timeout",
	"Expected data is not provided",
	"Function is called when it should not be",
	"Credentials are not enough to use the function",
}

var xretcodes = []string{
	"Unknown return code",
	"Operation in progress",
	"Try again",
	"Interrupted",
	"Not allowed",
	"Host not found",
	"Network not available",
	"Connection not available",
	"No space left",
	"Result too large",
	"Not found",
}

func (e CddsErrorType) Error() string {
	return retcodes[int(e)]
}

func (e CddsErrorType) XError() string {
	return xretcodes[int(e)]
}

func CddsStrretcode(rc uint32) string {
	nretcodes := uint32(len(retcodes))
	nxretcodes := uint32(len(xretcodes))

	if DDS_XRETCODE_BASE >= 0 {
		panic("DDS_XRETCODE_BASE must be negative")
	}

	if uint32(rc) == ^uint32(math.MaxInt32) {
		return xretcodes[0]
	}

	if rc < 0 {
		rc = -rc
	}

	if rc >= 0 && rc < nretcodes {
		return retcodes[rc]
	} else if rc >= uint32(-DDS_XRETCODE_BASE) && rc < uint32(-DDS_XRETCODE_BASE)+nxretcodes {
		return xretcodes[rc-uint32(-DDS_XRETCODE_BASE)]
	} else {
		return xretcodes[0]
	}
}
