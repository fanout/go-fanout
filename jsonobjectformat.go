//    jsonobjectformat.go
//    ~~~~~~~~~
//    This module implements the JsonObjectFormat struct.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

type JsonObjectFormat struct {
    value string
}

func (format JsonObjectFormat) Name() string {
    return "json-object"
}

func (format JsonObjectFormat) Export() interface{} {
    return format.value
}
