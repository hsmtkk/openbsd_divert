import socket

import scapy.utils
from scapy.all import IP, Ether

IPPROTO_DIVERT = 258
DIVERT_PORT = 700
PCAP_FILE = 'out.pcap'

with socket.socket(socket.AF_INET, socket.SOCK_RAW, IPPROTO_DIVERT) as sock:
    sock.bind(('0.0.0.0', DIVERT_PORT))
    with scapy.utils.PcapWriter(PCAP_FILE) as writer:
        while True:
            data, peer = sock.recvfrom(2048)
            print("received %d bytes" % len(data))
            eth = Ether()/IP(data)
            writer.write(eth)
            sock.sendto(data, peer)
