# reverse-shell

Be aware that you will need netcat installed on the target machine in order for the reverse shell to work.

You will need to replace IP_ADDRESS and PORT with the IP address and port number that you want the reverse shell to connect to.

On the machine you want to connect to you need to have netcat listening on the specified port.

```
nc -nlvp PORT
```

Once you run the reverse shell file on the target machine, it will connect back to your machine and give you a command prompt.
