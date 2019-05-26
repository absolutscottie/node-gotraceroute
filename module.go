package main

import (
	"C"
)

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type Hop struct {
	Number   int    `json:"number"`
	Address  string `json:"address"`
	Rtt      int    `json:"rtt"`
	Complete bool   `json:"complete"`
	Error    string `json:"error"`
}

//export PingHost
func PingHost(hostname string, ttl, timeoutMillis, packetSize int) *C.char {
	result := Hop{
		Number:   ttl,
		Address:  "*",
		Rtt:      -1,
		Complete: false,
		Error:    "",
	}

	ips, err := net.LookupIP(hostname)
	if err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	var dst net.IPAddr
	for _, ip := range ips {
		if ip.To4() != nil {
			dst.IP = ip
			break
		}
	}
	if dst.IP == nil {
		result.Error = "IP not found."
		return EncodeResult(result)
	}

	c, err := net.ListenPacket("ip4:1", "0.0.0.0") // ICMP for IPv4
	if err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	defer c.Close()
	p := ipv4.NewPacketConn(c)

	if err := p.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	rb := make([]byte, packetSize)

	wm.Body.(*icmp.Echo).Seq = ttl
	wb, err := wm.Marshal(nil)
	if err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	if err := p.SetTTL(ttl); err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}

	// In the real world usually there are several
	// multiple traffic-engineered paths for each hop.
	// You may need to probe a few times to each hop.
	begin := time.Now()
	if _, err := p.WriteTo(wb, nil, &dst); err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	if err := p.SetReadDeadline(time.Now().Add(time.Duration(timeoutMillis) * time.Millisecond)); err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	n, _, peer, err := p.ReadFrom(rb)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return EncodeResult(result)
		}
		result.Error = err.Error()
		return EncodeResult(result)
	}
	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		result.Error = err.Error()
		return EncodeResult(result)
	}
	rtt := time.Since(begin)

	// In the real world you need to determine whether the
	// received message is yours using ControlMessage.Src,
	// ControlMessage.Dst, icmp.Echo.ID and icmp.Echo.Seq.
	switch rm.Type {
	case ipv4.ICMPTypeTimeExceeded:
		result.Address = peer.String()
		if names, _ := net.LookupAddr(result.Address); len(names) > 0 {
			result.Address = names[0]
		}
		result.Rtt = int(rtt.Seconds() * 1000.0)

	case ipv4.ICMPTypeEchoReply:
		result.Address = peer.String()
		if names, _ := net.LookupAddr(result.Address); len(names) > 0 {
			result.Address = names[0]
		}
		result.Rtt = int(rtt.Seconds() * 1000.0)
		result.Complete = true

	default:
		//not sure
	}

	return EncodeResult(result)
}

func EncodeResult(result Hop) *C.char {
	outputBytes, _ := json.Marshal(result)
	return C.CString(string(outputBytes))
}

func main() {
}
