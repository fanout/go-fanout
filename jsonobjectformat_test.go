//    jsonobjectformat_test.go
//    ~~~~~~~~~
//    This module implements the JsonObjectFormat tests.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

import ("testing"
        "github.com/stretchr/testify/assert")

func TestJsonObjectFormatName(t *testing.T) {
    fmt := &JsonObjectFormat{value:"value"}
    assert.Equal(t, fmt.Name(), "json-object")
}

func TestJsonObjectFormatExport(t *testing.T) {
    fmt := &JsonObjectFormat{value:"value"}
    assert.Equal(t, fmt.Export(), "value")
}
