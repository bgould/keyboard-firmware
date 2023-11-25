package main

type TOTPKey struct {
	Name string
	Key  SecureString
}

type SecureString string

func (s SecureString) String() string {
	return "<redacted>"
}

var (
	totpKeys = []TOTPKey{
		TOTPKey{Name: "GitHub", Key: "KWNKKXLJYPHGSFCB"},
	}
)
