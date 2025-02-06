package _rsa

import (
	"fmt"
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_hash"
	"testing"
)

func TestGenerate(t *testing.T) {
	{
		rsa := Generate()
		fmt.Println(rsa.Pub)
		fmt.Println(rsa.Pri)
	}
}
func TestEncode(t *testing.T) {
	// no need to test
}
func TestDecode(t *testing.T) {
	{
		var pub string = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtPXQuyD47ae0i/dpbvC7
cNgvksdD/IUy4iTHqFmZXTdINJzdnc1GvGHiQgoxZo+aiCbb8Iv0Hkte4RKDfh8j
3K7aXSl73t3cQg/LavzBdfBiK59df0U1WLAM7hzaSFq8oDz2ZpH4HrANVECc3Zsa
VDZTHCvcTGrHBh4YzPjiwy0X4AXfQ7xVStv5OxZqdyyLDe/H0xeWj4cN4CVr21lI
GTQtxLBO+aI5P+7moUqyir0Pg3cml8CMJyhzIWn9bbmWVSznEcqqcxkbq1N4Y8Z1
kkkS1oqZV6sfaqdOUvXHZeEWGEvAPtQZTmsXIbbcgMQjddxxB1j1JLM3/hsWcBRL
pXg7XtEEHxPog8C2GqSMRegJ0wQGNEg3BpaPKhRRec1G6LRDpmYkx+ixqxXtVz1J
3X09+nnS7svi7jjitSi4r8M+kMjA9+ehtHx8YOoyXCiyRcZHWr7E/xbReFXXWN+7
Y6HmECoMtxOpmK/4YK8d0cy8zw2slMLY5XDJMsSUSQJqUFIC/z8+fIOdpdIrsep0
DvZrPhrg0SZTe4nMBZPjtfxG2Gbwsn6oS+M8vn1FZTUbvaeqvAMVbxgxCzhsnMv6
ItMZ4j6sVzAWvuQM5j4eWBXte2qk59yUh/drwmDy9l/gQ96MUciqZEEJItIJDPYy
rx56NcNN25VfX06/pW92CAcCAwEAAQ==
-----END PUBLIC KEY-----`
		var pri string = `-----BEGIN PRIVATE KEY-----
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
Gf2locNcrQ9l7yC1LNnKPGSL0R//
-----END PRIVATE KEY-----`
		var expect string = "hello world!"
		encode := Encode(expect, pub)
		get := Decode(encode, pri)
		_assert.Equal(t, expect, get)
	}
}
func TestSign(t *testing.T) {
	// no need to test
	t.SkipNow()
}
func TestVerify(t *testing.T) {
	{
		var pub string = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAtPXQuyD47ae0i/dpbvC7
cNgvksdD/IUy4iTHqFmZXTdINJzdnc1GvGHiQgoxZo+aiCbb8Iv0Hkte4RKDfh8j
3K7aXSl73t3cQg/LavzBdfBiK59df0U1WLAM7hzaSFq8oDz2ZpH4HrANVECc3Zsa
VDZTHCvcTGrHBh4YzPjiwy0X4AXfQ7xVStv5OxZqdyyLDe/H0xeWj4cN4CVr21lI
GTQtxLBO+aI5P+7moUqyir0Pg3cml8CMJyhzIWn9bbmWVSznEcqqcxkbq1N4Y8Z1
kkkS1oqZV6sfaqdOUvXHZeEWGEvAPtQZTmsXIbbcgMQjddxxB1j1JLM3/hsWcBRL
pXg7XtEEHxPog8C2GqSMRegJ0wQGNEg3BpaPKhRRec1G6LRDpmYkx+ixqxXtVz1J
3X09+nnS7svi7jjitSi4r8M+kMjA9+ehtHx8YOoyXCiyRcZHWr7E/xbReFXXWN+7
Y6HmECoMtxOpmK/4YK8d0cy8zw2slMLY5XDJMsSUSQJqUFIC/z8+fIOdpdIrsep0
DvZrPhrg0SZTe4nMBZPjtfxG2Gbwsn6oS+M8vn1FZTUbvaeqvAMVbxgxCzhsnMv6
ItMZ4j6sVzAWvuQM5j4eWBXte2qk59yUh/drwmDy9l/gQ96MUciqZEEJItIJDPYy
rx56NcNN25VfX06/pW92CAcCAwEAAQ==
-----END PUBLIC KEY-----`
		var pri string = `-----BEGIN PRIVATE KEY-----
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
Gf2locNcrQ9l7yC1LNnKPGSL0R//
-----END PRIVATE KEY-----`
		var data string = "hello world"
		var hashed []byte = _hash.DecodeString(_hash.Sha256(data))

		sign := Sign(hashed, pri)
		var expect bool = true
		get := Verify(hashed, sign, pub)
		_assert.Equal(t, expect, get)
	}
}
