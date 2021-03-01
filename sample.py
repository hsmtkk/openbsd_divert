import socket

import scapy.utils

IPPROTO_DIVERT = 258
DIVERT_PORT = 700
PCAP_FILE = 'out.pcap'

with socket.socket(socket.AF_INET, socket.SOCK_RAW, IPPROTO_DIVERT) as sock:
    sock.bind(('0.0.0.0', DIVERT_PORT))
    with scapy.utils.PcapWriter(PCAP_FILE) as writer:
        while True:
            data, peer = sock.recvfrom(2048)
            print("received %d bytes from %s" % (len(data), str(peer)))
            writer.write(data)
            sock.sendto(data, peer)
