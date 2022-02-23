package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCase(t *testing.T, URL string, origins []string, expected bool) {
	result := URLMatchesOrigins(URL, origins)
	assert.Equal(t, expected, result)
}

func Test_Origin_WildcardAlwaysMatches(t *testing.T) {
	testCase(t, "http://example.com/static/foo.js", []string{"https://foo.com/", "*"}, true)
}

func Test_Origin_Matches(t *testing.T) {
	testCase(t, "http://example.com/static/foo.js", []string{"https://foo.com/", "http://example.com/"}, true)
}

func Test_Origin_MatchesWithWildcard(t *testing.T) {
	testCase(t, "http://foo.bar.com/static/foo.js", []string{"https://foo.com/", "http://*.bar.com/"}, true)
}

func Test_Origin_DoesNotMatch(t *testing.T) {
	testCase(t, "http://example.com/static/foo.js", []string{"https://foo.com/", "http://test.com/"}, false)
}

func Test_Origin_DoesNotMatchWithWildcard(t *testing.T) {
	testCase(t, "http://foo.bar.com/static/foo.js", []string{"https://foo.com/", "http://*.baz.com/"}, false)
}
