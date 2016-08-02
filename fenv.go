// fenv package provides safe wrapping functions of fenv.h.
// When imported, fenv automatically turns on pragma for STDC FENV_ACCESS into ON.
// Thus you can safely control the functionality
package fenv

/*
#cgo LDFLAGS: -lm
#include <fenv.h>
#pragma STDC FENV_ACCESS ON

fenv_t* wrap_FE_DFL_ENV()
{
	return FE_DFL_ENV;
}

*/
import "C"

// Env is a type of fenv_t in C. Implementation in detail depends on architecture.
type FEnv C.fenv_t

// Type representing exception flags. For now, support only for the
// basic abstraction of flags that are either set or clear. fexcept_t
// could be  structure that holds more info about the fp environment.
type FExcept C.fexcept_t

// FPU status word exception flags
const (
	FE_INVALID    = C.FE_INVALID
	FE_DIVBYZERO  = C.FE_DIVBYZERO
	FE_OVERFLOW   = C.FE_OVERFLOW
	FE_UNDERFLOW  = C.FE_UNDERFLOW
	FE_INEXACT    = C.FE_INEXACT
	FE_ALL_EXCEPT = FE_INEXACT | FE_DIVBYZERO | FE_OVERFLOW | FE_UNDERFLOW | FE_INEXACT
)

// FPU Status word rounding flags
const (
	FE_TONEAREST  = C.FE_TONEAREST
	FE_UPWARD     = C.FE_UPWARD
	FE_DOWNARD    = C.FE_DOWNWARD
	FE_TOWARDZERO = C.FE_TOWARDZERO
)

var FE_DFT_ENV = (*FEnv)(C.wrap_FE_DFL_ENV())

/* Floating-point exceptions */

func ClearExcept(except int) int {
	ret := C.feclearexcept(C.int(except))
	return int(ret)
}

func RaiseExcept(except int) int {
	ret := C.feraiseexcept(C.int(except))
	return int(ret)
}

func ExceptFlag(flag *FExcept, except int) int {
	ret := C.fegetexceptflag((*C.fexcept_t)(flag), C.int(except))
	return int(ret)
}

func SetExceptFlag(flag *FExcept, except int) int {
	ret := C.fesetexceptflag((*C.fexcept_t)(flag), C.int(except))
	return int(ret)
}

/* Rounding direction */

func Round() int {
	return int(C.fegetround())
}

func SetRound(rdir int) int {
	ret := C.fesetround(C.int(rdir))
	return int(ret)
}

/* Entire environment */

func Env(env *FEnv) int {
	ret := C.fegetenv((*C.fenv_t)(env))
	return int(ret)
}

func SetEnv(env *FEnv) int {
	ret := C.fesetenv((*C.fenv_t)(env))
	return int(ret)
}

func HoldExcept(env *FEnv) int {
	ret := C.feholdexcept((*C.fenv_t)(env))
	return int(ret)
}

func UpdateEnv(env *FEnv) int {
	ret := C.feupdateenv((*C.fenv_t)(env))
	return int(ret)
}

/* Other */

func TestExcept(except int) int {
	return int(C.fetestexcept(C.int(except)))
}
