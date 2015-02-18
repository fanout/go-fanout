//    fanout.go
//    ~~~~~~~~~
//    This module implements the Fanout struct and features.
//    :authors: Konstantin Bokarius.
//    :copyright: (c) 2015 by Fanout, Inc.
//    :license: MIT, see LICENSE for more details.

package fanout

import "github.com/fanout/go-pubcontrol"

type Fanout struct {
    pubControl *pubcontrol.PubControl
}

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

func (fanout *Fanout) Publish(channel string, data, id, prev_id string) error {
    format := &JsonObjectFormat{value: data} 
    item := pubcontrol.NewItem([]pubcontrol.Formatter{format}, id, prev_id)
    return fanout.pubControl.Publish(channel, item)
}
