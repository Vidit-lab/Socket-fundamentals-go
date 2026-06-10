import socket

SERVER_ADDRESS = ("127.0.0.1", 8080)

def main() -> None:
    message = input("Enter Message:\n")

    with socket.create_connection(SERVER_ADDRESS) as conn:
        conn.sendall(message.encode("utf-8"))
        conn.shutdown(socket.SHUT_WR)

        response = bytearray()
        while True:
            chunk = conn.recv(1024)
            if not chunk:
                break
            response.extend(chunk)

    print("Received:")
    print(response.decode("utf-8"))


if __name__ == "__main__":
    main()