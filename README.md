# 🪞Echo
[Echo](https://www.rfc-editor.org/rfc/rfc862) protocol, implemented on top of TCP/UDP server.

### Usage 
```bash

$ ./echo -help
Usage of echo:
  -port string
    	port number to listen on (default "7")
  -proto string
    	protocol to use: tcp or udp (default "tcp")


$ ./echo.out -port 2280 -proto udp
2025/11/20 20:16:05 🚀 Echo udp server. on port 2280...
```
