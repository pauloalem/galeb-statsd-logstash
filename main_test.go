package main

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(S{})

func (S) TestHandle(c *check.C) {
	input := []byte("galeb.myapp_cloud_tsuru_com.10_236_99_181_32772.requestTime:44|ms")
	expected := map[string]string{
		"client": "tsuru",
		"metric": "response_time",
		"count":  "1",
		"app":    "myapp_cloud_tsuru_com",
		"value":  "44",
	}
	c.Assert(handle(input), check.DeepEquals, expected)
}
