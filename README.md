go-fanout
===========

Author: Konstantin Bokarius <kon@fanout.io>

A Go convenience library for publishing messages to Fanout.io using the EPCP protocol.

License
-------

go-fanout is offered under the MIT license. See the LICENSE file.

Installation
------------

```sh
go get github.com/fanout/go-fanout
```

go-fanout requires jwt-go 2.2.0 and go-pubcontrol 1.0.0. To ensure that the correct version of both of these dependencies are installed use godeps:

```sh
go get github.com/tools/godep
cd $GOPATH/src/github.com/fanout/go-fanout
$GOPATH/bin/godep restore
```

Usage
-----

```go
package main

import "github.com/fanout/go-fanout"
import "encoding/base64"

func main() {
    // Decode the base64 encoded key:
    decodedKey, err := base64.StdEncoding.DecodeString("<key>")
    if err != nil {
        panic("Failed to base64 decode the key")
    }

    // Pass true as the 3rd parameter for SSL or false for non-SSL:
    fo := fanout.NewFanout("<realm>", decodedKey, true)

    // Publish to the Fanout.io realm:
    err = fo.Publish("<channel>", "Test publish!", "", "")
    if err != nil {
        panic("Publish failed with: " + err.Error())
    }
}
```
