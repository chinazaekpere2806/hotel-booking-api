package utils

import (
	"net/mail"
	//"fmt"
	"net"
	"strings"
)

func IsValidSyntax(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func HasMailServer(email string) bool {
	// Split the email at the "@" symbol to get the domain
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	// Look up the MX (Mail Exchange) records for the domain
	mxRecords, err := net.LookupMX(domain)

	// If there's no error and we found at least one record, it's valid
	return err == nil && len(mxRecords) > 0
}
