#gorackhdRedfish

### Go bindings for RackHD Redfish API

## Description
gorackhdRedfish represents API bindings for Go that allow you to manage the [RackHD](https://github.com/RackHD/RackHD) Redfish API. This specific API is different than the [gorackhd bindings](https://github.com/emccode/gorackhd) as these are related to IPMI functions for individual machines. The intended functions are a direct implementation of what's available within the RackHD Redfish API. 

## Compatibility
gorackhdRedfish is created using [go-swagger](https://github.com/go-swagger/go-swagger) client generator. 

The client relies on a swagger spec file that is taken from RackHD's [on-http](https://github.com/RackHD/on-http) library and packaged at `/swagger-spec/redfish.yml`. As the swagger spec changes, this client will update as needed. 

## Installation & Creation of the Client
The client is found within the `/client` folder which is generated by go-swagger. To generate/install the client:
```
$ go get github.com/emccode/gorackhd
$ cd $GOPATH/src/github.com/emccode/gorackhd
$ make
```

The Makefile utilizes [glide](https://github.com/Masterminds/glide) to reference git commit `6b712512cbe` of [go-swagger](https://github.com/go-swagger/go-swagger) that has been tested and known to be working. The Makefile will generate a go-swagger directory in `/vendor/github.com/` and then create the client. 

## Environment Variables
| Name        | Description           |
| ------------- |:-------------:|
| `GORACKHD_ENDPOINT`      | the API endpoint. `localhost:9090` is used by default if not set             |


## Using the client

This sample script will retrieve (GET) all Roles. Take a look at `gorackhdredfish_test.go` for more

```
package main

import (
    "fmt"
    "log"
    "os"

    httptransport "github.com/go-swagger/go-swagger/httpkit/client"
    "github.com/go-swagger/go-swagger/strfmt"

    apiclientRedfish "github.com/emccode/gorackhdRedfish/client"
    "github.com/emccode/gorackhdRedfish/client/redfish_v1"
)

func main() {

    // create the transport
    transport := httptransport.New("localhost:9090", "/redfish/v1", []string{"http"})

    // configure the host. include port with environment variable. For instance the vagrant image would be localhost:9090
    if os.Getenv("GORACKHD_ENDPOINT") != "" {
        transport.Host = os.Getenv("GORACKHD_ENDPOINT")
    }

    // create the API client, with the transport
    client := apiclientRedfish.New(transport, strfmt.Default)

    //use any function to do REST operations
    resp, err := client.RedfishV1.GetRoles(nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%#v\n", resp.Payload)
}

```

### Use the client to POST
Import the `redfish_v1` library for all the functions, then also import the `models` for the pre-defined structs needed for payloads.:
```
import (
    ...
    "github.com/emccode/gorackhdRedfish/client/redfish_v1"
    "github.com/emccode/gorackhdRedfish/models"
    ...
)
```

### Use the client to fetch (GET) individual items

Import the library for what you are trying to retrieve:
```
import (
    ...
    "github.com/emccode/gorackhdRedfish/client/redfish_v1"
    ...
)
```

Lookup the params that are required in a struct:
```
resp, err := client.RedfishV1.GetSystem(&redfish_v1.GetSystemParams{Identifier: "57154fe9d67951e70958c213"})
if err != nil {
      log.Fatal(err)
}
    
fmt.Printf("%#v\n", resp.Payload.PowerState)

```

## Contribution
Create a fork of the project into your own repository. 

Run the tests provided in `gorackhdRedfish_test.go` to verify GET/POST still function (change the NodeID accordingly):
```
env DEBUG=1 go test -run TestRedfishGetRolesOperation -v
env DEBUG=1 go test -run TestRedfishGetPowerStatus -v
env DEBUG=1 go test -run TestRedfishShutdown -v
```

If tests do not pass, please create an issue so it can be addressed or fix and create a PR.

If all tests pass, make changes and create a pull request with a description on what was added or removed and details explaining the changes in lines of code. If approved, project owners will merge it.

## Licensing
gorackhd is freely distributed under the [MIT License](http://emccode.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.

##Support
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [EMC {code} Community slack team](http://community.emccode.com/) and join the **#rackhd** channel. Lastly, for questions asked on [Stackoverflow.com](https://stackoverflow.com) please tag them with **EMC**. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.






