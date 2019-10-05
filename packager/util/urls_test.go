package util

import (
	//"fmt"
	"net/url"
	//"strings"
	"testing"

	//"github.com/alecthomas/participle/lexer"
	//"github.com/alecthomas/participle/lexer/ebnf"
	util_test "github.com/ampproject/amppackager/packager/util/testing"
	"github.com/flyingmutant/rapid"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Run this test with `-rapid.checks=100000`.

/*
// Translated from https://tools.ietf.org/html/rfc3986#appendix-A.
// URI is the only terminal symbol. Note that participle don't seem to support
// matching empty strings against non-terminal productions, so every production
// that is supposed to be able to match an empty string has to instead be
// wrapped in [...] at the "callsite".
var rfc3986 = lexer.Must(ebnf.New(`
   URI           = scheme ":" [hierPart]  [ "?" [query]  ] [ "#" [fragment] ] .

   hierPart      = "//" [authority] [pathAbempty]
                 | pathAbsolute
                 | pathRootless
                 | [pathEmpty] .

   // Unused, but left here for reference (pun only slightly intended):
   // uriReference  = URI | [relativeRef] .
   // absoluteURI   = scheme ":" [hierPart]  [ "?" [query] ] .
   // relativeRef   = [relativePart] [ "?" [query] ] [ "#" [fragment] ] .
   // relativePart  = "//" [authority] [pathAbempty]
   //               | pathAbsolute
   //               | pathNoscheme
   //               | [pathEmpty] .

   scheme        = alpha { alpha | digit | "+" | "-" | "." } .

   authority     = [ userinfo "@" ] [host] [ ":" [port] ] .
   userinfo      = { unreserved | pctEncoded | subDelims | ":" } .
   host          = ipLiteral | ipv4address | [regName] .
   port          = { digit } .

   ipLiteral     = "[" ( ipv6address | ipvFuture  ) "]" .

   ipvFuture     = "v" hexdig { hexdig } "." ipvFuturePart { ipvFuturePart } .
   ipvFuturePart = unreserved | subDelims | ":" .

   ipv6address   = h16 ":" h16 ":" h16 ":" h16 ":" h16 ":" h16 ":" ls32
                 | "::" h16 ":" h16 ":" h16 ":" h16 ":" h16 ":" ls32
		 | [ h16 ] "::" h16 ":" h16 ":" h16 ":" h16 ":" ls32
		 | [ h16 [ ":" h16 ] ] "::" h16 ":" h16 ":" h16 ":" ls32
		 | [ h16 [ ":" h16 [ ":" h16 ] ] ] "::" h16 ":" h16 ":" ls32
		 | [ h16 [ ":" h16 [ ":" h16 [ ":" h16 ] ] ] ] "::" h16 ":" ls32
		 | [ h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 ] ] ] ] ] "::" ls32
		 | [ h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 ] ] ] ] ] ] "::" h16
		 | [ h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 [ ":" h16 ] ] ] ] ] ] ] "::" .

   h16           = hexdig [ hexdig [ hexdig [ hexdig ] ] ] .
   ls32          = ( h16 ":" h16 ) | ipv4address .
   ipv4address   = decOctet "." decOctet "." decOctet "." decOctet .

   decOctet      = digit                         // 0-9
                 | "\x31" … "\x39" digit         // 10-99
                 | "1" digit digit               // 100-199
                 | "2" "\x30" … "\x34" digit   // 200-249
                 | "25" "\x30" … "\x35" .      // 250-255

   regName       = { unreserved | pctEncoded | subDelims } .

   // Unused, but left here for reference:
   // path          = pathAbempty    // begins with "/" or is empty
   //               | pathAbsolute   // begins with "/" but not "//"
   //               | pathNoscheme   // begins with a non-colon segment
   //               | pathRootless   // begins with a segment
   //               | [pathEmpty] .    // zero characters

   pathAbempty   = { "/" [segment] } .
   pathAbsolute  = "/" [ segmentNz { "/" [segment] } ] .
   pathNoscheme  = segmentNzNc { "/" [segment] } .
   pathRootless  = segmentNz { "/" [segment] } .
   // NOTE: Participle doesn't support zero-length rules; instead we do a fake
   // non-epsilon rule, plus the [...] trick explained above.
   pathEmpty     = "WOOooOWWWooOWWOooooWWWwwWW".

   segment       = { pchar } .
   segmentNz     = pchar { pchar } .
   segmentNzNc   = segmentNzNcPart { segmentNzNcPart } .  // non-zero-length segment without any colon ":"
   segmentNzNcPart   = unreserved | pctEncoded | subDelims | "@" .

   pchar         = unreserved | pctEncoded | subDelims | ":" | "@" .

   query         = { pchar | "/" | "?" } .

   fragment      = { pchar | "/" | "?" } .

   pctEncoded    = "%" hexdig hexdig .

   unreserved    = alpha | digit | "-" | "." | "_" | "~" .
   reserved      = genDelims | subDelims .
   genDelims     = ":" | "/" | "?" | "#" | "[" | "]" | "@" .
   subDelims     = "!" | "$" | "&" | "'" | "(" | ")"
                 | "*" | "+" | "," | ";" | "=" .

   // Imported from https://tools.ietf.org/html/rfc2234#section-6.1.
   alpha          =  "\x41" … "\x5A" | "\x61" … "\x7A" .  // A-Z / a-z
   //alpha          =  "A" … "Z" | "a" … "z" .
   digit          =  "\x30" … "\x39" .  // 0-9
   hexdig         =  digit | "A" | "B" | "C" | "D" | "E" | "F" .
`))
*/

func EncodingIsValid1(t *testing.T, rawURL string) {
	u, err := url.Parse(rawURL)
	require.NoError(t, err)
	s, err := FormatURL(u)
	require.NoError(t, err)
	println("s =", s)
	require.True(t, util_test.URLIsValid(s))
}

func EncodingIsValid(t *rapid.T, rawURL string) {
	u, err := url.Parse(rawURL)
	require.NoError(t, err)
	//fmt.Printf("%#v\n", u)

	s, err := FormatURL(u)

	// This test is only interesting for URLs that are actually formatted;
	// don't waste trials on errors.
	if err != nil {
		t.SkipNow()
	}

	//fmt.Println(s)

	require.True(t, util_test.URLIsValid(s))
}

// Ensure that, given any URL (valid or not), the resultant format is a
// valid URL.
func TestAllEncodingsAreValid(t *testing.T) {
	urls := rapid.Strings().Filter(func(raw string) bool {
		u, err := url.Parse(raw)
		return err == nil && u.IsAbs()
	})
	rapid.Check(t, func(t *rapid.T) {
		EncodingIsValid(t, urls.Draw(t, "url").(string))
	})
}

// Same as above, limited to https:// URLs, to help the fuzzer find interesting
// cases.
func TestAllEncodingsAreValid_HTTPS(t *testing.T) {
	urls := rapid.Strings().Filter(func(raw string) bool {
		//fmt.Printf("%q\n", raw)
		u, err := url.Parse("https://" + raw)
		return err == nil && u.IsAbs()
	})
	rapid.Check(t, func(t *rapid.T) {
		EncodingIsValid(t, "https://" + urls.Draw(t, "url").(string))
	})
}

func TestKnownBug(t *testing.T) {
	EncodingIsValid1(t, "https://~[")
}

/*
// Not a bug any more?
func TestKnownBug_CombiningChar(t *testing.T) {
	rawURL := "https://" + "҉ᾉ\U000f1c84"
	u, err := url.Parse(rawURL)
	require.NoError(t, err)
	//fmt.Printf("%#v\n", u)
	s, err := FormatURL(u)
	require.NoError(t, err)
	fmt.Println(s)

	require.True(t, util_test.URLIsValid(s))
}
*/

func EncodingPreservesMeaning(t *rapid.T, rawURL string) {
	u, err := url.Parse(rawURL)
	require.NoError(t, err)

	s, err := FormatURL(u)

	// This test is only interesting for URLs that are actually formatted;
	// don't waste trials on errors.
	if err != nil {
		t.SkipNow()
	}

	require.True(t, util_test.URLsMatch(rawURL, s), "%s -> %s", rawURL, s)
}

// Ensure that, given a valid URL, the only modifications are those that
// preserve semantics, e.g.
// https://en.wikipedia.org/wiki/URL_normalization#Normalizations_that_preserve_semantics.
func TestValidURLsAreSemanticallyEquivalent_HTTPS(t *testing.T) {
	urls := rapid.Strings().Filter(func(raw string) bool {
		u, err := url.Parse("https://" + raw)
		return err == nil && u.IsAbs()
	})
	rapid.Check(t, func(t *rapid.T) {
		EncodingPreservesMeaning(t, "https://" + urls.Draw(t, "url").(string))
	})
}

func TestKnownBug2(t *testing.T) {
	EncodingPreservesMeaning(t, "https://")
}
