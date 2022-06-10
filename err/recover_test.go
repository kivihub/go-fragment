package err

import (
	"testing"
)

var logger *testing.T

func TestRecover(t *testing.T) {
	logger = t

	ret := f0()
	t.Log("[TestRecover] f0 result is", ret)

	ret = f1()
	t.Log("[TestRecover] f1 result is", ret)

	ret = f2()
	t.Log("[TestRecover] f2 result is", ret)
}

func f0() (ret bool) {
	defer func() {
		if e := recover(); e != nil {
			logger.Log("[f0] recover", e)
		}
	}()
	logger.Log("[f0] start")
	panic("panic in f0")
	logger.Log("[f0] end")
	return true
}

func f1() (ret bool) {
	defer func() {
		if e := recover(); e != nil {
			logger.Log("[f1] recover", e)
			ret = false
		}
	}()
	logger.Log("[f1] start")
	panic("panic in f1")
	logger.Log("[f1] end")
	return true
}

func f2() (ret bool) {
	defer func() {
		if e := recover(); e != nil {
			logger.Log("[f2] recover", e)
			ret = true
		}
	}()
	logger.Log("[f2] start")
	panic("panic in f2")
	logger.Log("[f2] end")
	return true
}
