from socket import *

s = socket(AF_INET, SOCK_STREAM)

s.bind(("localhost", 12345))
s.listen(1)
conn, addr = s.accept() # returns new socket

while True: # forever
    data = conn.recv(1024) # receive data from client
    if not data: break  # stop if client stopped
    msg = data.decode()+"*" # process the incoming data into a response
    print(msg)
    conn.send(msg.encode()) # return response
conn.close() # close the connection



