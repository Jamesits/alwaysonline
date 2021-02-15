# AlwaysOnline

Bypass NCSI and portal detection on a network level.

[![Build Status](https://dev.azure.com/nekomimiswitch/General/_apis/build/status/alwaysonline?branchName=master)](https://dev.azure.com/nekomimiswitch/General/_build/latest?definitionId=89&branchName=master)

## Usage

Ports required: tcp/80, tcp/53, udp/53.

Start the server:

```shell script
alwaysonline [--ipv4 192.168.1.2] [--ipv6 fd00::2]
```

(The IP addresses are the server IP addresses on the user-facing interface. If the server is behind destination NAT, use the public IP. The server will always listen on all the IPs available; these IP hints are just for faking DNS results.)

Hijack the following domains on your DNS server to the alwaysonline server:

```
// Windows
www.msftncsi.com
www.msftconnecttest.com

// iOS, macOS
captive.apple.com

// Android
connectivitycheck.gstatic.com
connectivitycheck.android.com
connect.rom.miui.com
network-test.debian.org
```
