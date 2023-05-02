package awsip

import (
	"encoding/json"
	"io"
	"os"
)

// Ranges contains the set of IP prefixes for IPv4 and IPv6 for AWS.
type Ranges struct {
	// The publication time, in Unix epoch time format.
	SyncToken string `json:"syncToken"`
	// The publication date and time, in UTC YY-MM-DD-hh-mm-ss format.
	CreateDate string `json:"createDate"`
	// IPv4 ranges.
	Prefixes []Prefix `json:"prefixes"`
	// IPv6 ranges.
	Ipv6Prefixes []Prefix `json:"ipv6_prefixes"`
}

// Prefix is the IP prefix ranges
type Prefix struct {
	// The public IPv6 address range, in CIDR notation. Note that AWS may advertise a prefix in more specific ranges.
	Ipv6Prefix *string `json:"ipv6_prefix,omitempty"`
	// The AWS Region or GLOBAL for edge locations. The CLOUDFRONT and ROUTE53 ranges are GLOBAL.
	Region string `json:"region"`
	// The subset of IP address ranges. The addresses listed for API_GATEWAY are
	// egress only. Specify AMAZON to get all IP address ranges (meaning that
	// every subset is also in the AMAZON subset). However, some IP address
	// ranges are only in the AMAZON subset (meaning that they are not also
	// available in another subset).
	Service string `json:"service"`
	// The name of the network border group, which is a unique set of
	// Availability Zones or Local Zones from where AWS advertises IP addresses.
	NetworkBorderGroup string `json:"network_border_group"`
	// The public IPv4 address range, in CIDR notation. Note that AWS may
	// advertise a prefix in more specific ranges. For example, prefix
	// 96.127.0.0/17 in the file may be advertised as 96.127.0.0/21,
	// 96.127.8.0/21, 96.127.32.0/19, and 96.127.64.0/18.
	IPPrefix *string `json:"ip_prefix,omitempty"`
}

// Read ranges from io.Reader
func Read(reader io.Reader) (*Ranges, error) {
	decoder := json.NewDecoder(reader)
	var ranges Ranges
	err := decoder.Decode(&ranges)
	if err != nil {
		return nil, err
	}
	return &ranges, nil
}

// ReadFile reads ranges from file
func ReadFile(filename string) (*Ranges, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f)
}
