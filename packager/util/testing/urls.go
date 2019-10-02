package testing

import (
	"strings"
)

// uriparser built with:
// $ cd vendor/uriparser-0.9.3 && mkdir build && cd build
// $ cmake -DCMAKE_BUILD_TYPE=Release -DURIPARSER_BUILD_TESTS=off -DURIPARSER_BUILD_DOCS=off ..
// $ make

// #cgo CFLAGS: -I${SRCDIR}/../../../vendor/uriparser-0.9.3/include
// #cgo LDFLAGS: -L${SRCDIR}/../../../vendor/uriparser-0.9.3/build -luriparser
//
// #include "uriparser/Uri.h"
// int URLErrorCode(const char* url) {
//   UriUriA parsed;
//   return uriParseSingleUriA(&parsed, url, NULL);
// }
// int URLsMatch(const char* url1, const char* url2) {
//   UriUriA parsed1, parsed2;
//   if (uriParseSingleUriA(&parsed1, url1, NULL)) return 0;
//   if (uriParseSingleUriA(&parsed2, url2, NULL)) return 0;
//   if (uriNormalizeSyntaxA(&parsed1)) return 0;
//   if (uriNormalizeSyntaxA(&parsed2)) return 0;
//   return uriEqualsUriA(&parsed1, &parsed2);
// }
import "C"

func containsNUL(url string) bool {
	// Strings with NUL byte can't be converted to a C-string, because
	// uriparser may mistakenly think they're valid based only on the
	// portion before the terminator. TODO(twifkak): Pass [data,len]
	// instead, e.g. via the *Ex form of ParseSingleUri.
	return strings.IndexByte(url, 0) >= 0
}

func URLIsValid(url string) bool {
	if containsNUL(url) {
		return false
	}
	return C.URLErrorCode(C.CString(url)) == 0
}

func URLsMatch(url1, url2 string) bool {
	if containsNUL(url1) || containsNUL(url2) {
		return false
	}
	return C.URLsMatch(C.CString(url1), C.CString(url2)) == 1
}
