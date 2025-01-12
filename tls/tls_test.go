package tls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTLSConfigBuilder_AppendX509KeyPair(t *testing.T) {
	var (
		fakeCertPath = "/tmp/localhost.crt"
		fakeCert     = `
-----BEGIN CERTIFICATE-----
MIIDhDCCAmygAwIBAgIhANG4j9rHsCqX42dYqniv7bUWIakPRii44fVDlAFslA1N
MA0GCSqGSIb3DQEBBQUAMFYxCTAHBgNVBAYTADEJMAcGA1UECgwAMQkwBwYDVQQL
DAAxEDAOBgNVBAMMB296b24ucnUxDzANBgkqhkiG9w0BCQEWADEQMA4GA1UEAwwH
b3pvbi5ydTAeFw0yMjA2MDYxNjAyMDlaFw0zMjA2MDYxNjAyMDlaMEQxCTAHBgNV
BAYTADEJMAcGA1UECgwAMQkwBwYDVQQLDAAxEDAOBgNVBAMMB296b24ucnUxDzAN
BgkqhkiG9w0BCQEWADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMGr
dnRoQe4Dn+8GDx2jviswx0KFo7EGz9T+RtFQblTpgGcSWSx0irRk5ce/kDveDqYh
4ZgNBn4XGnDcogWJAGcq4XymCQq5jBHJ0q7n2GxZ+Tx5vlmyaDNG5zv1lLjgSl0P
22ClM8NM7hCnE3uZz90VVyQbZ9FFKFW2WX+3uJHJnwMtW+w/n8DnQOBlIAjyCRko
sRci9/dgaWFmx+Yscs1UPtG7pgEOO+6hiyFNgZEVGbgZv82jR4g8zr6SxcIT8ZLT
A4oOj+7GFSf8q6yNgp8bDNxbBS3rS2Vv6IrIjXsVfbilWGM3WkAVaOVD6OFJ65wc
AiggDm9oTFz0H2OOAZ0CAwEAAaNPME0wHQYDVR0OBBYEFHP9uHY/St1je+k0cZ/U
0ZO2KWJ8MB8GA1UdIwQYMBaAFHP9uHY/St1je+k0cZ/U0ZO2KWJ8MAsGA1UdEQQE
MAKCADANBgkqhkiG9w0BAQUFAAOCAQEANr3j2VVGCa87fHgJzs3k7IBq0VnXx1C1
tuWsJ96z0MPHcOD2g8OU2ZANBSGn/PhS71fqvRefr4wUKDLt947/Ycriai84lMv8
tWO5KLrHDGG8OuMy5orHVt2okcPFVsVmzJBxIv6fRtrzr152FvvRZegrsgnYloZL
rJdkHHl+9kHS02Nmf3+F6ANAg/jgnr0huJEt1xwB0YtKHTj346ksBP3OOyjolze9
J6xWsrb/fzFceDYNmIW5oyDm9j6ZxmgRa0turlo/pCp4gEojlyTGhg0HzxoPJmky
BQQbQaVBlv7XzP9xayJxOOtKiSXjv6OMFP1XFHZPzvx5/3VbVPrKsQ==
-----END CERTIFICATE-----`
		fakePrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAwat2dGhB7gOf7wYPHaO+KzDHQoWjsQbP1P5G0VBuVOmAZxJZ
LHSKtGTlx7+QO94OpiHhmA0GfhcacNyiBYkAZyrhfKYJCrmMEcnSrufYbFn5PHm+
WbJoM0bnO/WUuOBKXQ/bYKUzw0zuEKcTe5nP3RVXJBtn0UUoVbZZf7e4kcmfAy1b
7D+fwOdA4GUgCPIJGSixFyL392BpYWbH5ixyzVQ+0bumAQ477qGLIU2BkRUZuBm/
zaNHiDzOvpLFwhPxktMDig6P7sYVJ/yrrI2CnxsM3FsFLetLZW/oisiNexV9uKVY
YzdaQBVo5UPo4UnrnBwCKCAOb2hMXPQfY44BnQIDAQABAoIBAB62CnA2MjJEp+F2
9AGuvBLMRuTRHV6OpnlvoxpBJ2XWq1Js8fAfQPmPQHLW+U/NDESi/cunUR7AmiRI
kHbRPVuezKgAClIFj2UovHyY7lRsRh4lbh56MJTlCvkSnfVdN92fDJo8SEORmPTS
EcBZaUF1R17F68YfMeC5rGKY5y9S432LrsbMapfZhBOBLlSec4Xb1VRAy0OfLzka
2Gt9B4P4sHzzO6w+ODL/cedfZt687iCfaCzW3QwxSxnelFEb2GOD0KGhuIgN1Nkh
N+DsAMmjLk5cYwlRKbwv7zuohcSi7MAu/ziKcaKPTGi4kB1cWhVNIuVmz0Ow6lVM
b1gGfAECgYEA4+E8IEbMmgIpfMmgmYhY1DNcSxp7dqJBwx9i/lbF+OMWo7HFf9zS
3AlyO3uPXZwgY7DtN4wCYZNFaZxZ+ZiUbMUd/tFKJiFSgx/H1RzKZnP9L60s00t5
Zx/bIkKv4SnOkBEmYr0x3y6cIvdCRJpK3gDRNdTsiQ2wROp6EhmuV50CgYEA2ZGG
thkfL6oFqN5pu070qoBAdxSIerVQPWMugChP03G0omNggwvJQ1DBFIhJSYGzyy/E
vtZczX6ra1AF07H5Xd/R3Dnl2UkQYAo/0wX0dRczl4InB3YUKUkGjpiagQIs3obg
I8LGzB+gigcPG5cBbn9Q2mCVlN0W7X2qeiffMgECgYA9eNHuZwrkdLbaWc2//P7x
Z1V7UQ6DJywURdA/igrno+LEj70WS+x5vPaawy8ecnJuUhj7xgISblboMWw4H4fS
e64mwjB9brnCc31FqPmUf6+J/+46pX5/aiGD8XdehxeY13eCZUUhzoICEK9IsFKH
5rKJEgIoyo+FV7dMUK8uJQKBgQCL/eh5GTcI21Si8OGFhWtGnhlcxdh/ZFUJ+tx7
1/RtMDezWVBLYfURnE9wX7UQWCWQ8t6ckJ/MFdpExYvKSDUVIyQMmTB9HFcuBMpG
hQljohcvQK7OTTrxyawvap/XrMekM5LbT8PMqfkJdztQXFyudbtXmFgHHi9Xhsl/
qlIiAQKBgA9DhReh6kSqhP0ldvw9f5z2PrVJqGOaCgkRQ/dUuDhbvlgyXI0gS/Rj
cG6ys3uVh5sYM7eVdZDKrKg45cQddUkbZm07cymaJ0WSSVBvWcUPPZ3aBlMWY4Vj
rO3s7sbqs3cMwC5g9LjPsfieoZt9AppSrJeY+l74yNa/afqTscun
-----END RSA PRIVATE KEY-----`
	)

	b := NewConfigBuilder()
	b.readFile = func(name string) ([]byte, error) {
		if name == fakeCertPath {
			return []byte(fakeCert), nil
		}
		panic("invalid file name")
	}

	err := b.AppendX509KeyPair(fakeCert, fakePrivateKey)
	require.NoError(t, err)

	tlsConfig := b.Build()

	require.NotNil(t, tlsConfig)
	require.Len(t, tlsConfig.Certificates, 1)
}

func TestTLSConfigBuilder_AppendCARoot(t *testing.T) {
	var (
		fakeCertPath = "/tmp/localhost.crt"
		fakeCerts    = `
-----BEGIN CERTIFICATE-----
MIIDhDCCAmygAwIBAgIhANG4j9rHsCqX42dYqniv7bUWIakPRii44fVDlAFslA1N
MA0GCSqGSIb3DQEBBQUAMFYxCTAHBgNVBAYTADEJMAcGA1UECgwAMQkwBwYDVQQL
DAAxEDAOBgNVBAMMB296b24ucnUxDzANBgkqhkiG9w0BCQEWADEQMA4GA1UEAwwH
b3pvbi5ydTAeFw0yMjA2MDYxNjAyMDlaFw0zMjA2MDYxNjAyMDlaMEQxCTAHBgNV
BAYTADEJMAcGA1UECgwAMQkwBwYDVQQLDAAxEDAOBgNVBAMMB296b24ucnUxDzAN
BgkqhkiG9w0BCQEWADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMGr
dnRoQe4Dn+8GDx2jviswx0KFo7EGz9T+RtFQblTpgGcSWSx0irRk5ce/kDveDqYh
4ZgNBn4XGnDcogWJAGcq4XymCQq5jBHJ0q7n2GxZ+Tx5vlmyaDNG5zv1lLjgSl0P
22ClM8NM7hCnE3uZz90VVyQbZ9FFKFW2WX+3uJHJnwMtW+w/n8DnQOBlIAjyCRko
sRci9/dgaWFmx+Yscs1UPtG7pgEOO+6hiyFNgZEVGbgZv82jR4g8zr6SxcIT8ZLT
A4oOj+7GFSf8q6yNgp8bDNxbBS3rS2Vv6IrIjXsVfbilWGM3WkAVaOVD6OFJ65wc
AiggDm9oTFz0H2OOAZ0CAwEAAaNPME0wHQYDVR0OBBYEFHP9uHY/St1je+k0cZ/U
0ZO2KWJ8MB8GA1UdIwQYMBaAFHP9uHY/St1je+k0cZ/U0ZO2KWJ8MAsGA1UdEQQE
MAKCADANBgkqhkiG9w0BAQUFAAOCAQEANr3j2VVGCa87fHgJzs3k7IBq0VnXx1C1
tuWsJ96z0MPHcOD2g8OU2ZANBSGn/PhS71fqvRefr4wUKDLt947/Ycriai84lMv8
tWO5KLrHDGG8OuMy5orHVt2okcPFVsVmzJBxIv6fRtrzr152FvvRZegrsgnYloZL
rJdkHHl+9kHS02Nmf3+F6ANAg/jgnr0huJEt1xwB0YtKHTj346ksBP3OOyjolze9
J6xWsrb/fzFceDYNmIW5oyDm9j6ZxmgRa0turlo/pCp4gEojlyTGhg0HzxoPJmky
BQQbQaVBlv7XzP9xayJxOOtKiSXjv6OMFP1XFHZPzvx5/3VbVPrKsQ==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDkDCCAnigAwIBAgIhALuZRcvEN09TpK9d/JMBv7bI+6dAIPmelKPuaKqt6QtE
MA0GCSqGSIb3DQEBBQUAMF4xCTAHBgNVBAYTADEJMAcGA1UECgwAMQkwBwYDVQQL
DAAxFDASBgNVBAMMC2V4YW1wbGUuY29tMQ8wDQYJKoZIhvcNAQkBFgAxFDASBgNV
BAMMC2V4YW1wbGUuY29tMB4XDTIyMDYwNjE3NTUxN1oXDTMyMDYwNjE3NTUxN1ow
SDEJMAcGA1UEBhMAMQkwBwYDVQQKDAAxCTAHBgNVBAsMADEUMBIGA1UEAwwLZXhh
bXBsZS5jb20xDzANBgkqhkiG9w0BCQEWADCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALSIvNQXKhOH6LjCVbWLmX8RZ4+9nyJU8deyAQN+0ia3TtZar4XT
zwa52OUXBtZ5y32QG3bD32AH7kXi/nAZHAT1NsQjvB7kjj33/rW0UnqAuo+LXqYc
msP8XPHEpfNSaLy+o3ML+m4axM9zBQq4Nk4Fzu2qT290mIHaKEZ7E4iqm6bn6q0e
cMxI3sT8x2c2Z6XsDtYkFANgnrZc9wi3rwjtprPb2PUUB8DSZSyPvLKybF2lKS+P
LneekHQnopME/EfWIkJnH9I5490nSTL/fBmIki31VD7EIymPfxhcJqy2aF7kjZZk
dZdo79DJocgVbILIffhPfvIEy6z0g9yNa08CAwEAAaNPME0wHQYDVR0OBBYEFKdG
UhQYqfxUesro4q3pylgXKlzjMB8GA1UdIwQYMBaAFKdGUhQYqfxUesro4q3pylgX
KlzjMAsGA1UdEQQEMAKCADANBgkqhkiG9w0BAQUFAAOCAQEAjubPnFFx1nJoCJn5
GZK9/gvYPJn19djBNXypj7zvOg/XSs3+12tL+Ruc9HtW7HMW4NO/FfYz/f21NEe0
TmppBXOz/RrP4+GUB7MEpiykSo/RlC/nOEabhgOl1NQE0i/Yf2AHrq5Z1A52TpKl
04+RhGsF7nm0EQWIEPmqu/LmgsaIFV0MFW2IqeMeJsnDGCL2CLg0WPxMjCoC3ta2
46uT1Qe85GCi2NbJeZj+rVAV+xy52y5EfLY4N6PTnIKmgb66YXZ3LX0E3ocUYNY0
O9RjSvhWSEkLMQTd0o/5ghQPPMady5Lucb+9iORcdzYywQ9TE1jMN1ooqZnabXk9
Pafatg==
-----END CERTIFICATE-----`
	)

	b := NewConfigBuilder()
	b.readFile = func(name string) ([]byte, error) {
		if name == fakeCertPath {
			return []byte(fakeCerts), nil
		}
		panic("invalid file name")
	}

	err := b.AppendCARoot(fakeCertPath)
	require.NoError(t, err)

	tlsConfig := b.Build()

	require.NotNil(t, tlsConfig)
	require.NotNil(t, tlsConfig.RootCAs)
}
