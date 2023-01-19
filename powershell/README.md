# Powershell Reverse Shell

This code creates a new TCP client object and connects to the attacker's IP address and port number specified. It then creates a stream to read and write data and creates an array of bytes for the data. The script then enters a while loop to continuously read data from the stream, execute it using the "iex" command, and send the output back to the attacker. The script also appends the current working directory to the output before sending it back. Finally, it closes the client connection.
