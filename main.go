package main

import(
  "fmt"
  "log"
  "syscall"
)

const(
  IPPROTO_DIVERT = 258
  DIVERT_PORT = 700
)

func main(){
  fd,err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, IPPROTO_DIVERT)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("fd=%d\n", fd)

  sockAddr := syscall.SockaddrInet4 {
    Port: DIVERT_PORT,
  }
  if err:=syscall.Bind(fd, &sockAddr); err != nil {
    log.Fatal(err)
  }
  for {
    var buf []byte
    recved, peer, err := syscall.Recvfrom(fd, buf, 0)
    if err != nil {
      log.Fatal(err)
    }
    if recved == 0 {
      continue
    }
    fmt.Printf("received %d bytes from %v\n", recved, peer)
    
    if err := syscall.Sendto(fd, buf, 0, peer); err != nil {
      log.Fatal(err)
    }
    fmt.Printf("sent %d bytes to %v\n", len(buf), peer)
  }
}

/*
#define DIVERT_PORT 700

int
main(int argc, char *argv[])
{
	int fd, s;
	struct sockaddr_in sin;
	socklen_t sin_len;

	fd = socket(AF_INET, SOCK_RAW, IPPROTO_DIVERT);
	if (fd == -1)
		err(1, "socket");

	memset(&sin, 0, sizeof(sin));
	sin.sin_family = AF_INET;
	sin.sin_port = htons(DIVERT_PORT);
	sin.sin_addr.s_addr = 0;

	sin_len = sizeof(struct sockaddr_in);

	s = bind(fd, (struct sockaddr *) &sin, sin_len);
	if (s == -1)
		err(1, "bind");

	for (;;) {
		ssize_t n;
		char packet[IP_MAXPACKET];
		struct ip *ip;
		struct tcphdr *th;
		int hlen;
		char src[48], dst[48];

		memset(packet, 0, sizeof(packet));
		n = recvfrom(fd, packet, sizeof(packet), 0,
		    (struct sockaddr *) &sin, &sin_len);
		if (n == -1) {
			warn("recvfrom");
			continue;
		}
		if (n < sizeof(struct ip)) {
			warnx("packet is too short");
			continue;
		}

		ip = (struct ip *) packet;
		hlen = ip->ip_hl << 2;
		if (hlen < sizeof(struct ip) || ntohs(ip->ip_len) < hlen ||
		    n < ntohs(ip->ip_len)) {
			warnx("invalid IPv4 packet");
			continue;
		}

		th = (struct tcphdr *) (packet + hlen);

		if (inet_ntop(AF_INET, &ip->ip_src, src,
		    sizeof(src)) == NULL)
			(void)strlcpy(src, "?", sizeof(src));

		if (inet_ntop(AF_INET, &ip->ip_dst, dst,
		    sizeof(dst)) == NULL)
			(void)strlcpy(dst, "?", sizeof(dst));

		printf("%s:%u -> %s:%u\n",
		    src,
		    ntohs(th->th_sport),
		    dst,
		    ntohs(th->th_dport)
		);

		n = sendto(fd, packet, n, 0, (struct sockaddr *) &sin,
		    sin_len);
		if (n == -1)
			warn("sendto");
	}

	return 0;
}
*/
