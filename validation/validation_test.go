package main_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type testEntry struct {
	field      interface{}
	tag        string
	expectPass bool
}

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func tester(t *testing.T, testEntries []testEntry) {
	for _, entry := range testEntries {
		errs := validate.Var(entry.field, entry.tag)
		if entry.expectPass {
			assert.Nil(t, errs)
		} else {
			assert.Error(t, errs)
			assert.NotEmpty(t, errs)
		}
	}
}

func TestBasicValidation(t *testing.T) {
	testEntries := []testEntry{
		{`blah`, "required", true},
		{``, "required", false},
		{`sam`, "eq=sam", true},
		{`sam`, "min=3", true},
		{`s`, "min=3", false},
		{`los angeles`, "len=11", true},
		{`los`, "len=11", false},
		{20, "eq=20", true},
		{20, "gte=12,lte=120", true},
		{10, "gte=12,lte=120", false},
		{200, "gte=12,lte=120", false},
		{`AU`, "iso3166_1_alpha2|iso3166_1_alpha3", true},    // We can combine 2 tags using |
	}

	tester(t, testEntries)
}

func TestNetworkValidation(t *testing.T) {
	testEntries := []testEntry{
		{`10.12.150.0/32`, "cidr", true},
		{`10.12.150.0`, "cidr", false},
		{`abc`, "cidr", false},
		{`10.12.150.0`, "ip", true},
		{`10.12.150.0/32`, "ip", false},
		{`abc`, "ip", false},
		{`www.github.com`, "fqdn", true},
		{`github.com`, "fqdn", true},
		{`github.com/login`, "fqdn", false},
		{`10.12.150.0`, "fqdn", false},
		{`www.github.com:8000`, "hostname_port", true},
		{`localhost:8000`, "hostname_port", true},
		{`http://github.com:8000/profile/sam?format=json`, "uri", true},
		{`http://10.12.150.0/profile/sam`, "uri", true},
	}

	tester(t, testEntries)
}

func TestStringValidation(t *testing.T) {
	testEntries := []testEntry{
		{`abcdefghijklmnopqrstuvwxyz`, "alpha", true},
		{`abc123`, "alpha", false},
		{`123`, "alpha", false},
		{`abc`, "alphanum", true},
		{`abc123`, "alphanum", true},
		{`123`, "alphanum", true},
		{`abc123好`, "alphanum", false},
		{`abc123好`, "alphanumunicode", true},
		{`abc123`, "ascii", true},
		{`øµ`, "ascii", false},
		{`true`, "boolean", true},
		{`TRUE`, "boolean", true},
		{`false`, "boolean", true},
		{`FALSE`, "boolean", true},
		{true, "boolean", false},   // Boolean value validated as false
		{false, "boolean", false},  // Boolean value validated as false
		{`123abcdefg`, "contains=abc", true},
		{`123abcdefg`, "contains=xyz", false},
		{`123abcdefg`, "containsany=1fa", true},
		{`123abcdefg`, "excludes=xyz", true},
		{`123abcdefg`, "lowercase", true},
		{`123ABCDEFG`, "lowercase", false},
		{`123abcdefg`, "uppercase", false},
		{`123ABCDEFG`, "uppercase", true},
		{`123`, "numeric", true},
		{`abc`, "numeric", false},
		{`abc`, "startsnotwith=c", true},
		{`abc`, "startsnotwith=a", false},
		{`abc`, "startswith=a", true},
		{`abc`, "startswith=c", false},
		{"\xab\xf3\x3d\x45\xe2\x65\x1a\x8e", "multibyte", true},
		{`abc`, "multibyte", false},
	}

	tester(t, testEntries)
}

func TestFormatValidation(t *testing.T) {
	testEntries := []testEntry{
		{`aGVsbG8gd29ybGQ=`, "base64", true},
		{`abc`, "base64", false},
		{`2021-09-12`, "datetime=2006-01-02", true},
		{`2021-09-12T08:29:23`, "datetime=2006-01-02T15:04:05", true},
		{`+16175551212`, "e164", true},
		{`+447555123456`, "e164", true},
		{`6175551212`, "e164", false},
		{`+1-617-555-1212`, "e164", false},
		{`+44-7-555-123-456`, "e164", false},
		{`sam@example.com`, "email", true},
		{`example.com`, "email", false},
		{`sam@example`, "email", false},
		{`0xEF`, "hexadecimal", true},
		{`0xef`, "hexadecimal", true},
		{`EF`, "hexadecimal", true},
		{`EFG`, "hexadecimal", false},
		{`#c33`, "iscolor", true},
		{`#cc3333`, "iscolor", true},
		{`rgb(204,51,51)`, "iscolor", true},
		{`rgba(204,51,51,0.5)`, "iscolor", true},
		{`AU`, "iso3166_1_alpha2", true},
		{`YZ`, "iso3166_1_alpha2", false},
		{`au`, "iso3166_1_alpha2", false},
		{`AUS`, "iso3166_1_alpha3", true},
		{`XYZ`, "iso3166_1_alpha3", false},
		{`aus`, "iso3166_1_alpha3", false},
		{`AUD`, "iso4217", true},
		{`XYD`, "iso4217", false},
		{`aud`, "iso4217", false},
		{`{"name": "sam"}`, "json", true},
		{`{name: "sam"}`, "json", false},
		{`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiIsImlhdCI6MTUxNjIzOTAyMn0.eXipevc_vYLoox-5ZDiXGxkNA_YczIVBQke2QuEokTg`, "jwt", true},
		{`180`, "longitude", true},
		{180, "longitude", true},
		{`180.00`, "longitude", true},
		{`-180`, "longitude", true},
		{-180, "longitude", true},
		{`-180.00`, "longitude", true},
		{`200.00`, "longitude", false},
		{`-200.00`, "longitude", false},
		{`90`, "latitude", true},
		{90, "latitude", true},
		{`90.00`, "latitude", true},
		{`-90`, "latitude", true},
		{-90, "latitude", true},
		{`-90.00`, "latitude", true},
		{`100.00`, "latitude", false},
		{`-100.00`, "latitude", false},
		{`90731`, "postcode_iso3166_alpha2=US", true},
		{`90731-1234`, "postcode_iso3166_alpha2=US", true},
		{`AB123`, "postcode_iso3166_alpha2=US", false},
		{`K1A 0B1`, "postcode_iso3166_alpha2=CA", true},
		{`123-45-6789`, "ssn", true},
		{`123456789`, "ssn", false},
		{`US/Eastern`, "timezone", true},
		{`us/eastern`, "timezone", true},
		{`UTC`, "timezone", true},
		{`d44c8d26-f304-4834-a8cf-a2ad0f018164`, "uuid4", true},
	}

	tester(t, testEntries)
}
