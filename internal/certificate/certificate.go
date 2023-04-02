package certificate

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/vvrnv/gossl/internal/log"
)

type Certificate struct {
	DnsResolvedTo string
	Subject       pkix.Name
	IssuerName    pkix.Name
	DNSNames      []string
	StartDate     string
	ExpireDate    string
	PubAlgorithm  string
	SigAlgorithm  string
}

type CertificateInfo struct {
	Subject            string
	Issuer             string
	NotBefore          time.Time
	NotAfter           time.Time
	SignatureAlgorithm string
	PublicKeyAlgorithm string
}

type Connection struct {
	transport http.Transport
}

func (c Certificate) getSubject() pkix.Name {
	return c.Subject
}

func (c Certificate) getIssuerName() pkix.Name {
	return c.IssuerName
}

func (c Certificate) getDNSNames() []string {
	return c.DNSNames
}

func (c Certificate) getStartDate() string {
	return c.StartDate
}

func (c Certificate) getExpireDate() string {
	return c.ExpireDate
}

func (c Certificate) getPubAlgorithm() string {
	return c.PubAlgorithm
}

func (c Certificate) getSigAlgorithm() string {
	return c.SigAlgorithm
}

func SetTransport(host, ip string, timeout, port int) *Connection {

	transport := http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Duration(timeout) * time.Second,
	}

	dialer := &net.Dialer{
		Timeout:   time.Duration(timeout) * time.Second,
		KeepAlive: time.Duration(timeout) * time.Second,
		DualStack: true,
	}
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		if addr == fmt.Sprintf("%s:%d", host, port) {
			addr = fmt.Sprintf("%s:%d", ip, port)
		}
		return dialer.DialContext(ctx, network, addr)
	}

	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS13,
	}

	return &Connection{
		transport: transport,
	}
}

func GetCertificateInfo(ip string, host string, timeout int, port int, output string) error {

	c := SetTransport(host, ip, timeout, port)
	transport := c.transport

	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Duration(timeout)*time.Second)
	if err != nil {
		return log.Error(err)
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), transport.TLSClientConfig)
	if err != nil {
		return log.Error(err)
	}
	defer conn.Close()

	certificateField(conn.ConnectionState().PeerCertificates, ip, output)

	return nil
}

func expireDateCountToColor(expireDate string) string {
	nowFormat, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	expireFormat, _ := time.Parse("2006-01-02", expireDate)

	days := int32(expireFormat.Sub(nowFormat).Hours() / 24)
	if days < 0 {
		return color.HiRedString(fmt.Sprintf("(%v days from today)", days))
	} else if days <= 30 {
		return color.HiYellowString(fmt.Sprintf("(%v days from today)", days))
	}
	return color.HiGreenString(fmt.Sprintf("(%v days from today)", days))
}

func certificateField(peerCertificates []*x509.Certificate, ip string, output string) {
	for _, cert := range peerCertificates {
		if len(cert.DNSNames) > 0 {
			formatDate := "2006-01-02"
			x509C := &Certificate{
				DnsResolvedTo: ip,
				Subject:       cert.Subject,
				IssuerName:    cert.Issuer,
				DNSNames:      cert.DNSNames,
				StartDate:     cert.NotBefore.Format(formatDate),
				ExpireDate:    cert.NotAfter.Format(formatDate),
				PubAlgorithm:  cert.PublicKeyAlgorithm.String(),
				SigAlgorithm:  cert.SignatureAlgorithm.String(),
			}

			switch output {
			case "json":
				json, err := json.Marshal(x509C)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(json))
				fmt.Println("")
			case "":
				altNames := x509C.getDNSNames()

				fmt.Printf("\n%s %s\n", "DNS resolved to:", ip)
				fmt.Println("Subject Alternative Names:", strings.Join(altNames[:], ","))
				fmt.Println("Subject:", x509C.getSubject().String())
				fmt.Println("Issuer Name:", x509C.getIssuerName().String())
				fmt.Println("Start Date:", x509C.getStartDate())

				colorDays := expireDateCountToColor(x509C.getExpireDate())
				fmt.Println("Expire Date:", fmt.Sprintf("%s %s", x509C.getExpireDate(), colorDays))
				fmt.Println("Public Key Algorithm:", x509C.getPubAlgorithm())
				fmt.Println("Signature Algorithm:", x509C.getSigAlgorithm())
			default:
				log.Fatal("Unsupported output format")
			}
		}
	}
}
