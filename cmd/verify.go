package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vvrnv/gossl/internal/certificate"
	"github.com/vvrnv/gossl/internal/ip"
	"github.com/vvrnv/gossl/internal/log"
)

var (
	server  string
	timeout int
	port    int
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify SSL certificate",
	Long: `verify SSL certificate with domain name or ip address.

For example:
  gossl verify -s domain.com
  gossl verify --server 8.8.8.8`,
	Run: func(_ *cobra.Command, args []string) {

		ips, err := ip.GetIPV4(server)
		if err != nil {
			log.Fatal(err)
		}

		for _, ip := range ips {
			err = certificate.GetCertificateInfo(ip, server, timeout, port)
			if err != nil {
				log.Error(err)
			}
		}
	},
}

func init() {
	verifyCmd.Flags().StringVarP(&server, "server", "s", "", "enter domain name or ip address (required)")
	verifyCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "enter timeout in seconds for verify command. (optional)")
	verifyCmd.Flags().IntVarP(&port, "port", "p", 443, "enter SSL port for verify command. (optional)")
	verifyCmd.MarkFlagRequired("server")
	rootCmd.AddCommand(verifyCmd)
}
