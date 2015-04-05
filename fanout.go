//    fanout.go
//    ~~~~~~~~~
//    This module implements the Fanout struct and features.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

import "github.com/fanout/go-pubcontrol"

// The Fanout struct is used for publishing messages to Fanout.io and is
// configured with a Fanout.io realm and associated key. SSL can either
// be enabled or disabled.
type Fanout struct {
    pubControl *pubcontrol.PubControl
}

// Initialize with a specified realm, key, and a boolean indicating wther
// SSL should be enabled or disabled.
func NewFanout(realm string, key []byte, ssl bool) *Fanout {
    fanout := new(Fanout)
    protocol := "https://"
    if (!ssl) {
        protocol = "http://"
    }
    fanout.pubControl = pubcontrol.NewPubControl([]map[string]interface{} {
            map[string]interface{} {
            "uri": protocol + "api.fanout.io/realm/" + realm,
            "iss": realm, 
            "key": key}})
    return fanout
}

// Publish the specified data to the specified channel for the configured
// Fanout.io realm. Optionally provide an ID and previous ID to send along
// with the message.
func (fanout *Fanout) Publish(channel string, data, id, prev_id string) error {
    format := &JsonObjectFormat{value: data} 
    item := pubcontrol.NewItem([]pubcontrol.Formatter{format}, id, prev_id)
    return fanout.pubControl.Publish(channel, item)
}
