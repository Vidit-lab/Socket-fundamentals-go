import socket

SERVER_ADDRESS = ("127.0.0.1", 8080)

def handle_client(conn: socket.socket, address: tuple[str, int]) -> None:
    try:
        data = bytearray()
        while True:
            chunk = conn.recv(1024)
            if not chunk:
                break
            data.extend(chunk)

        if data:
            conn.sendall(data)
            print(f"Echoed {data.decode('utf-8')!r} to {address[0]}:{address[1]}")
    finally:
        conn.close()


def main() -> None:
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server:
        server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        server.bind(SERVER_ADDRESS)
        server.listen(1)
        print(f"Server listening on {SERVER_ADDRESS[0]}:{SERVER_ADDRESS[1]}")

        while True:
            conn, address = server.accept()
            handle_client(conn, address)


if __name__ == "__main__":
    main()