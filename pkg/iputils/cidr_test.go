// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package iputils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeCommon/pkg/iputils"
)

func TestParseCIDR(t *testing.T) {
	for _, cidrString := range []string{
		"192.168.2.100/24",
		"2607:5300:203:afac::/125",
	} {
		cidr, err := iputils.ParseCIDR(cidrString)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(cidr, "=> [", cidr.From(), "-", cidr.To(), "]")
	}
}
