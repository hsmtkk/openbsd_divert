import socket

IPPROTO_DIVERT = 258

with socket.socket(socket.AF_INET, socket.SOCK_RAW, IPPROTO_DIVERT) as sock:
    sock.bind(('0.0.0.0', 700))
    while True:
        data, peer = sock.recvfrom(2048)
        print("received %d bytes" % len(data))
        sock.sendto(data, peer)
