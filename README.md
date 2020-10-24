# AlwaysOnline

Bypass NCSI and portal detection on a network level.

Experimental; not tested.

[![Build Status](https://dev.azure.com/nekomimiswitch/General/_apis/build/status/alwaysonline?branchName=master)](https://dev.azure.com/nekomimiswitch/General/_build/latest?definitionId=89&branchName=master)

## Usage

Start the server:

```shell script
alwaysonline [--ipv4 192.168.1.2] [--ipv6 fd00::2]
```

(The IP addresses are the server IP addresses on the user-facing interface.)

Hijack the following domains on your DNS server to the alwaysonline server:

```
www.msftncsi.com
www.msftconnecttest.com
captive.apple.com
connectivitycheck.gstatic.com
connectivitycheck.android.com
connect.rom.miui.com
network-test.debian.org
```
