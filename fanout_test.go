//    fanout_test.go
//    ~~~~~~~~~
//    This module implements the Fanout tests.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

import ("testing"
        "github.com/stretchr/testify/assert")

func TestInitialize(t *testing.T) {
    fo := NewFanout("realm", []byte("key"), true)
    assert.NotNil(t, fo.pubControl)
    fo = NewFanout("realm", []byte("key"), false)    
    assert.NotNil(t, fo.pubControl)
}

func TestFormatUri(t *testing.T) {
    assert.Equal(t, formatUri("realm", true),
            "https://api.fanout.io/realm/realm")
    assert.Equal(t, formatUri("realm", false),
            "http://api.fanout.io/realm/realm")    
}

func TestPublish(t *testing.T) {
    fo := NewFanout("realm", []byte("key"), false)
    err := fo.Publish("chan", "data", "", "")
    assert.NotNil(t, err)
}

func TestGetItem(t *testing.T) {
    item := getItem("data", "id", "prev-id")
    export, _ := item.Export()
    assert.Equal(t, export["id"], "id");
    assert.Equal(t, export["prev-id"], "prev-id");
    assert.Equal(t, export["json-object"], "data");
}
