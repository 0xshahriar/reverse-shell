# Golang Reverse Shell

This code creates a connection to the attacker's IP and port, starts a new command (in this case, "/bin/sh"), and then creates a new scanner to read the attacker's input. The scanner reads the attacker's input and writes it to the command's stdin. The code then reads the command's stdout and stderr and sends them back to the attacker.
