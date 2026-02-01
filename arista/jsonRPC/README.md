So this script utilizes hostfile and host.cfg commands files. To configure and or issue exec commands on the devices. Its pretty simple the config files use the format of host.cfg.

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
          "lastStatusChangeTimestamp": 1769908292.048976
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
          "lastStatusChangeTimestamp": 1769908293.0955617
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
          "lastStatusChangeTimestamp": 1769908293.8168893
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
          "lastStatusChangeTimestamp": 1769908294.4603612
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
          "lastStatusChangeTimestamp": 1769908294.7915118
        }
      }
    },
    {},
    {}
  ]
}
```
