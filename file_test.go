package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

func TestMimeTypeFromExtension(t *testing.T) {
	cases := []struct{ in, want string }{
		{"photo.JPG", "image/jpeg"},
		{"image.png", "image/png"},
		{"doc.pdf", "application/pdf"},
		{"README", "application/octet-stream"},
		{"file.unknownext", "application/octet-stream"},
	}
	for _, c := range cases {
		if got := utilities.MimeTypeFromExtension(c.in); got != c.want {
			t.Errorf("MimeTypeFromExtension(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
