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
	expected := &document{
		Client: "tsuru",
		Metric: "response_time",
		Count:  1,
		App:    "myapp_cloud_tsuru_com",
		Value:  44,
	}
	doc, err := handle(input)
	c.Assert(err, check.IsNil)
	c.Assert(doc, check.DeepEquals, expected)
}
