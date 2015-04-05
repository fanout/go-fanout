//    jsonobjectformat.go
//    ~~~~~~~~~
//    This module implements the JsonObjectFormat struct.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

// The JSON object format used for publishing messages to Fanout.io. The
// Value field represents the data to be published.
type JsonObjectFormat struct {
    value string
}

// The name of the format.
func (format *JsonObjectFormat) Name() string {
    return "json-object"
}

// The method used to export the format data.
func (format *JsonObjectFormat) Export() interface{} {
    return format.value
}
