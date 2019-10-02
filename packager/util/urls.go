package util

import (
	"net/url"
	"strings"
	//"github.com/pkg/errors"
	//"golang.org/x/net/idna"
)

// Formats the URL in a way that guarantees conformance to the URI production
// of the https://tools.ietf.org/html/rfc3986 grammar (i.e. only absolute URLs).
// This is required by the Link header
// (https://tools.ietf.org/html/rfc5988#section-5). Other parts of SXG refer to
// https://url.spec.whatwg.org/ instead. My assumption/hope is that the IETF
// grammar is a strict subset of the WHATWG grammar, as the latter is intended
// to be more lenient.
func FormatURL(u *url.URL) (string, error) {
	// Evaluate "/..", by resolving the URL as a reference from itself.
	// This prevents malformed URLs from eluding the PathRE protections.
	// TODO(twifkak): Test this.
	ret := u.ResolveReference(u)

	// Escape special characters in the Opaque.
	ret.Opaque = url.PathEscape(ret.Opaque)

	//if !ret.IsAbs() {
	//	return "", errors.New("URL is not absolute")
	//}

	//if ret.Scheme != "" && ret.Opaque == "" && ret.Host == "" && (ret.ForceQuery || ret.RawQuery != "") {
	//	return "", errors.New("URL has query without authority")
	//}

	//host, err := idna.ToASCII(ret.Host)
	//if err != nil {
	//	return "", errors.Wrapf(err, "converting %q to punycode", ret.Host)
	//}
	//ret.Host = host
	//println("host:", host)

	ret.Host = strings.ReplaceAll(ret.Host, "<", "%3C")

	// Escape special characters in the query component such as "<" or "|"
	// (but not "&" or "=").
	ret.RawQuery = url.PathEscape(ret.RawQuery)

	return ret.String(), nil
}
