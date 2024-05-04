from socket import *

s = socket(AF_INET, SOCK_STREAM)

s.connect(("localhost", 12345)) # connect to server

msg = "Hello World"  # compose a message
s.send(msg.encode()) # send the message

data = s.recv(1024) # receive the response
print(data.decode()) # print the result
s.close() # close the connection
