### Update 02/01/2026

So this script utilizes a hostfile and host.cfg commands files. To configure and or issue exec commands on the devices. Its a pretty simple concept. You only call the hostfile.txt. It then uses the hosts names called to pull the .cfg files for the commands.

hostfile.txt
```
leaf1a
leaf1b
leaf2
spine1
spine2
```

The commands files are simply the hostname + .cfg

leaf1a.cfg
```
enable
configure
interface loopback69
 description leaf1a


show interfaces loopback 69
```
leaf1b.cfg
```
enable
configure
interface loopback69
 description leaf1b


show interfaces loopback 69
```
And so on and so forth.....


To run the script simply ***./run_cmds hostsfile.txt username password***

```
$ ./run_cmds hostfile.txt admin admin
hosts: [leaf1a leaf1b leaf2 spine1 spine2]
leaf1a.cfg
[enable configure interface loopback69  description leaf1a   show interfaces loopback 69]
{
  "jsonrpc": "2.0",
  "id": "1",
  "result": [
    {},
    {},
    {},
    {},
    {},
    {},
    {
      "interfaces": {
        "Loopback69": {
          "name": "Loopback69",
          "forwardingModel": "routed",
          "lineProtocolStatus": "up",
          "interfaceStatus": "connected",
          "hardware": "loopback",
          "interfaceAddress": [],
          "description": "leaf1a",
          "bandwidth": 0,
          "mtu": 65535,
          "l3MtuConfigured": false,
          "l2Mru": 0,
          "lastStatusChangeTimestamp": 1769908292.0489767
        }
      }
    }
  ]
}
leaf1b.cfg
[enable configure interface loopback69  description leaf1b   show interfaces loopback 69 ]
{
  "jsonrpc": "2.0",
  "id": "1",
  "result": [
    {},
    {},
    {},
    {},
    {},
    {},
    {
      "interfaces": {
        "Loopback69": {
          "name": "Loopback69",
          "forwardingModel": "routed",
          "lineProtocolStatus": "up",
          "interfaceStatus": "connected",
          "hardware": "loopback",
          "interfaceAddress": [],
          "description": "leaf1b",
          "bandwidth": 0,
          "mtu": 65535,
          "l3MtuConfigured": false,
          "l2Mru": 0,
          "lastStatusChangeTimestamp": 1769908293.0955615
        }
      }
    },
    {}
  ]
}
leaf2.cfg
[enable configure interface loopback69  description leaf2  show interfaces loopback 69 ]
{
  "jsonrpc": "2.0",
  "id": "1",
  "result": [
    {},
    {},
    {},
    {},
    {},
    {
      "interfaces": {
        "Loopback69": {
          "name": "Loopback69",
          "forwardingModel": "routed",
          "lineProtocolStatus": "up",
          "interfaceStatus": "connected",
          "hardware": "loopback",
          "interfaceAddress": [],
          "description": "leaf2",
          "bandwidth": 0,
          "mtu": 65535,
          "l3MtuConfigured": false,
          "l2Mru": 0,
          "lastStatusChangeTimestamp": 1769908293.816889
        }
      }
    },
    {}
  ]
}
spine1.cfg
[  enable configure interface loopback69  description spine1  show interfaces loopback 69]
{
  "jsonrpc": "2.0",
  "id": "1",
  "result": [
    {},
    {},
    {},
    {},
    {},
    {},
    {},
    {
      "interfaces": {
        "Loopback69": {
          "name": "Loopback69",
          "forwardingModel": "routed",
          "lineProtocolStatus": "up",
          "interfaceStatus": "connected",
          "hardware": "loopback",
          "interfaceAddress": [],
          "description": "spine1",
          "bandwidth": 0,
          "mtu": 65535,
          "l3MtuConfigured": false,
          "l2Mru": 0,
          "lastStatusChangeTimestamp": 1769908294.460361
        }
      }
    }
  ]
}
spine2.cfg
[enable configure interface loopback69  description spine2   show interfaces loopback 69  ]
{
  "jsonrpc": "2.0",
  "id": "1",
  "result": [
    {},
    {},
    {},
    {},
    {},
    {},
    {
      "interfaces": {
        "Loopback69": {
          "name": "Loopback69",
          "forwardingModel": "routed",
          "lineProtocolStatus": "up",
          "interfaceStatus": "connected",
          "hardware": "loopback",
          "interfaceAddress": [],
          "description": "spine2",
          "bandwidth": 0,
          "mtu": 65535,
          "l3MtuConfigured": false,
          "l2Mru": 0,
          "lastStatusChangeTimestamp": 1769908294.7915115
        }
      }
    },
    {},
    {}
  ]
}
```
