package main

import (
	"os"

	"github.com/remind101/emp/Godeps/_workspace/src/github.com/remind101/empire/pkg/heroku"
)

var cmdCertAttach = &Command{
	Run:      runCertAttach,
	Usage:    "cert-attach",
	NeedsApp: true,
	Category: "certs",
	Short:    "attach a certificate to an app",
	Long: `
Attaches an SSL certificate to an applications web process. When using the ECS backend, this will attach an IAM server certificate to the applications ELB.

Before running this command, you should upload your SSL certificate and key to IAM using the AWS CLI.

Examples:

    $ aws iam upload-server-certificate --server-certificate-name myServerCertificate --certificate-body file://public_key_cert_file.pem --private-key file://my_private_key.pem --certificate-chain file://my_certificate_chain_file.pem
    $ emp cert-attach myServerCertificate -a
`,
}

func runCertAttach(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.PrintUsage()
		os.Exit(2)
	}

	cert := args[0]

	_, err := client.AppUpdate(mustApp(), &heroku.AppUpdateOpts{
		Cert: &cert,
	})
	must(err)
}
