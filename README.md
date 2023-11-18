# Adapters
We will need many adapters, the first is SMS adapter using SIM800/SIM900a chips as adapter.
To get SMS and parse into instructions, invoke the SDK of sim800 with instructions.

Basic instructions:
1. Check if there is any new SMS, if yes, parse it and send to the server.(use `AT+CMGL="ALL"` to get all SMS, sim800c will return a list of SMS, each SMS is a string, parse it and get the instruction)
2. Transfer into instruction and send to the Gateway.
3. Get response from the Gateway and send back to the sender.

OK !

```
Use Go over 1.20 version!


└─[0] go run main.go
[]string{"/dev/cu.BLTH", "/dev/cu.Bluetooth-Incoming-Port", "/dev/cu.HUAWEIFreeBudsPro", "/dev/cu.SLAB_USBtoUART", "/dev/cu.sBoseQC35II", "/dev/cu.usbserial-0001", "/dev/tty.BLTH", "/dev/tty.Bluetooth-Incoming-Port", "/dev/tty.HUAWEIFreeBudsPro", "/dev/tty.SLAB_USBtoUART", "/dev/tty.sBoseQC35II", "/dev/tty.usbserial-0001"}Find serial port: /dev/cu.BLTH
Find serial port: /dev/cu.Bluetooth-Incoming-Port
Find serial port: /dev/cu.HUAWEIFreeBudsPro
Find serial port: /dev/cu.SLAB_USBtoUART  // set this if you use CP2102 chips USBTTL
Find serial port: /dev/cu.sBoseQC35II
Find serial port: /dev/cu.usbserial-0001
Find serial port: /dev/tty.BLTH
Find serial port: /dev/tty.Bluetooth-Incoming-Port
Find serial port: /dev/tty.HUAWEIFreeBudsPro
Find serial port: /dev/tty.SLAB_USBtoUART
Find serial port: /dev/tty.sBoseQC35II
Find serial port: /dev/tty.usbserial-0001
INFO[2023-11-18T20:51:32+03:00] Send Bytes                                    bytes="AT+CMGF=1\r\n" length=11
INFO[2023-11-18T20:51:33+03:00] Send Bytes                                    bytes="AT+CSCS=\"GSM\"\r\n" length=15
INFO[2023-11-18T20:51:34+03:00] Send Bytes                                    bytes="AT+CNMI=2,1\r\n" length=13
```
