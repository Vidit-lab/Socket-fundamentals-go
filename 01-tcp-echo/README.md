# TCP Echo Learning Notes

## Goal
Learn what a connection actually is by sending one message from a client and receiving the same message back from a server.

## Task
The client sends `Hello` so the server can prove it received the message and returned the same bytes.

## Expected Output
```text
$ go run client.go

Enter Message:
Hello

Received:
Hello
```

## Self Check
### What is a socket?
A socket is one endpoint of network communication, so a program can send and receive data through the operating system.

### What is an IP?
An IP address identifies a machine on a network, so the client knows where to reach the server.

### What is a port?
A port identifies the service on that machine, so the OS can route traffic to the correct program.

### Why must the server start first?
The server must listen before the client connects, because a connection can only be established if something is already accepting it.

### What does Accept() actually do?
Accept waits for a client connection, then returns a live communication channel once the handshake is complete.

## Learning
### TCP
TCP gives us ordered, reliable delivery, so the same message can travel across the network without being lost or rearranged.

### Connection
A connection is a negotiated channel between two endpoints, so both sides know they are talking to each other and can exchange bytes.

### Client-Server
The client initiates the conversation and the server waits for it, which is the basic pattern behind most networked systems.

### Ports
Ports let one machine run many networked programs at the same time, because each program can listen on a different number.

### Streams
TCP behaves like a stream of bytes, so the application decides how to frame a message and read it back correctly.

## Observations
The client opens a TCP connection, sends the text once, and closes the write side so the server knows the message is complete.

The server reads the full request, writes the exact bytes back, and then closes the connection so the client can finish reading.

This shows that the message is not magically shared between programs; it is copied across a real network channel managed by TCP.

## Results
The round trip confirms that `Hello` arrives intact because the server returned the same text the client sent.

## How To Verify
1. Start the server in one terminal from this folder with `go run server.go`.
2. Start the client in another terminal with `go run client.go`.
3. Type `Hello` and press Enter.
4. Confirm the client prints `Received:` followed by `Hello`.
5. Stop the server with `Ctrl+C` after the test.

## Python Version
The same example is also available in Python with the same behavior and the same result.

### Expected Output
```text
$ python3 client.py

Enter Message:
Hello

Received:
Hello
```

### How To Verify In Python
1. Start the server in one terminal from this folder with `python3 server.py`.
2. Start the client in another terminal with `python3 client.py`.
3. Type `Hello` and press Enter.
4. Confirm the client prints `Received:` followed by `Hello`.
5. Stop the server with `Ctrl+C` after the test.