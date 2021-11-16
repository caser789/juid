package uuid

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type badRand struct{}

func (r badRand) Read(buf []byte) (int, error) {
	for i, _ := range buf {
		buf[i] = byte(i)
	}
	return len(buf), nil
}

func TestBadRand(t *testing.T) {
	SetRand(badRand{})
	uuid1 := New()
	uuid2 := New()
	if uuid1 != uuid2 {
		t.Errorf("expected duplicates, got %q and %q\n", uuid1, uuid2)
	}

	SetRand(nil)
	uuid1 = New()
	uuid2 = New()
	if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q and %q\n", uuid1, uuid2)
	}
}

func TestVersionString(t *testing.T) {
	assert.Equal(t, Version(0b00000000).String(), "VERSION_0")
	assert.Equal(t, Version(0b00000001).String(), "VERSION_1")
	assert.Equal(t, Version(0b00000010).String(), "VERSION_2")
	assert.Equal(t, Version(0b00000011).String(), "VERSION_3")
	assert.Equal(t, Version(0b00000100).String(), "VERSION_4")
	assert.Equal(t, Version(0b00000101).String(), "VERSION_5")
	assert.Equal(t, Version(0b00000110).String(), "VERSION_6")
	assert.Equal(t, Version(0b00000111).String(), "VERSION_7")
	assert.Equal(t, Version(0b00001000).String(), "VERSION_8")
	assert.Equal(t, Version(0b00001001).String(), "VERSION_9")
	assert.Equal(t, Version(0b00001010).String(), "VERSION_10")
	assert.Equal(t, Version(0b00001011).String(), "VERSION_11")
	assert.Equal(t, Version(0b00001100).String(), "VERSION_12")
	assert.Equal(t, Version(0b00001101).String(), "VERSION_13")
	assert.Equal(t, Version(0b00001110).String(), "VERSION_14")
	assert.Equal(t, Version(0b00001111).String(), "VERSION_15")
	assert.Equal(t, Version(0b00010000).String(), "BAD_VERSION_16")
	assert.Equal(t, Version(0b00010001).String(), "BAD_VERSION_17")
}

func TestVariant(t *testing.T) {
	assert.Equal(t, Variant(Invalid).String(), "Invalid")
	assert.Equal(t, Variant(RFC4122).String(), "RFC4122")
	assert.Equal(t, Variant(Reserved).String(), "Reserved")
	assert.Equal(t, Variant(Microsoft).String(), "Microsoft")
	assert.Equal(t, Variant(Future).String(), "Future")
	assert.Equal(t, Variant(5).String(), "BadVariant5")
	assert.Equal(t, Variant(6).String(), "BadVariant6")
}

func TestRandomUUID(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 32; x++ {
		uuid := New()
		s := uuid.String()
		if m[s] {
			t.Errorf("New returned duplicated UUID %s\n", s)
		}
		m[s] = true
		if uuid.Variant() != RFC4122 {
			t.Errorf("Random UUID is variant %d\n", uuid.Variant())
		}
		if v := uuid.Version(); v != 4 {
			t.Errorf("Random UUID of version %s\n", v)
		}
	}
}

func TestUUIDParse(t *testing.T) {
	for x := 1; x < 32; x++ {
		uuid1 := New()
		uuid2 := Must(Parse(uuid1.String()))

		assert.Equal(t, uuid1, uuid2)
	}
}

func TestCoding(t *testing.T) {
	text := "7d444840-9dc0-11d1-b245-5ffdce74fad2"
	urn := "urn:uuid:7d444840-9dc0-11d1-b245-5ffdce74fad2"
	data := UUID{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}
	if v := data.String(); v != text {
		t.Errorf("%x: encoded to %s, expected %s\n", data, v, text)
	}
	if v := data.URN(); v != urn {
		t.Errorf("%x: urn is %s, expected %s\n", data, v, urn)
	}

	uuid := Must(Parse(text))

	assert.Equal(t, uuid, data)
}

func TestNew(t *testing.T) {
	m := make(map[UUID]bool)
	for x := 1; x < 32; x++ {
		s := New()
		if m[s] {
			t.Errorf("New returned duplicated UUID %s\n", s)
		}
		m[s] = true
		uuid := Must(Parse(s.String()))
		if uuid == Nil {
			t.Errorf("New returned %q which does not decode\n", s)
			continue
		}
		if v := uuid.Version(); v != 4 {
			t.Errorf("Random UUID of version %s\n", v)
		}
		if uuid.Variant() != RFC4122 {
			t.Errorf("Random UUID is variant %d\n", uuid.Variant())
		}
	}
}

type test struct {
	in      string
	version Version
	variant Variant
	isuuid  bool
}

var tests = []test{
	{"f47ac10b-58cc-0372-8567-0e02b2c3d479", 0, RFC4122, true},
	{"f47ac10b-58cc-1372-8567-0e02b2c3d479", 1, RFC4122, true},
	{"f47ac10b-58cc-2372-8567-0e02b2c3d479", 2, RFC4122, true},
	{"f47ac10b-58cc-3372-8567-0e02b2c3d479", 3, RFC4122, true},
	{"f47ac10b-58cc-4372-8567-0e02b2c3d479", 4, RFC4122, true},
	{"f47ac10b-58cc-5372-8567-0e02b2c3d479", 5, RFC4122, true},
	{"f47ac10b-58cc-6372-8567-0e02b2c3d479", 6, RFC4122, true},
	{"f47ac10b-58cc-7372-8567-0e02b2c3d479", 7, RFC4122, true},
	{"f47ac10b-58cc-8372-8567-0e02b2c3d479", 8, RFC4122, true},
	{"f47ac10b-58cc-9372-8567-0e02b2c3d479", 9, RFC4122, true},
	{"f47ac10b-58cc-a372-8567-0e02b2c3d479", 10, RFC4122, true},
	{"f47ac10b-58cc-b372-8567-0e02b2c3d479", 11, RFC4122, true},
	{"f47ac10b-58cc-c372-8567-0e02b2c3d479", 12, RFC4122, true},
	{"f47ac10b-58cc-d372-8567-0e02b2c3d479", 13, RFC4122, true},
	{"f47ac10b-58cc-e372-8567-0e02b2c3d479", 14, RFC4122, true},
	{"f47ac10b-58cc-f372-8567-0e02b2c3d479", 15, RFC4122, true},

	{"urn:uuid:f47ac10b-58cc-4372-0567-0e02b2c3d479", 4, Reserved, true},
	{"URN:UUID:f47ac10b-58cc-4372-0567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-0567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-1567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-2567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-3567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-4567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-5567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-6567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-7567-0e02b2c3d479", 4, Reserved, true},
	{"f47ac10b-58cc-4372-8567-0e02b2c3d479", 4, RFC4122, true},
	{"f47ac10b-58cc-4372-9567-0e02b2c3d479", 4, RFC4122, true},
	{"f47ac10b-58cc-4372-a567-0e02b2c3d479", 4, RFC4122, true},
	{"f47ac10b-58cc-4372-b567-0e02b2c3d479", 4, RFC4122, true},
	{"f47ac10b-58cc-4372-c567-0e02b2c3d479", 4, Microsoft, true},
	{"f47ac10b-58cc-4372-d567-0e02b2c3d479", 4, Microsoft, true},
	{"f47ac10b-58cc-4372-e567-0e02b2c3d479", 4, Future, true},
	{"f47ac10b-58cc-4372-f567-0e02b2c3d479", 4, Future, true},

	{"f47ac10b158cc-5372-a567-0e02b2c3d479", 0, Invalid, false},
	{"f47ac10b-58cc25372-a567-0e02b2c3d479", 0, Invalid, false},
	{"f47ac10b-58cc-53723a567-0e02b2c3d479", 0, Invalid, false},
	{"f47ac10b-58cc-5372-a56740e02b2c3d479", 0, Invalid, false},
	{"f47ac10b-58cc-5372-a567-0e02-2c3d479", 0, Invalid, false},
	{"g47ac10b-58cc-4372-a567-0e02b2c3d479", 0, Invalid, false},
}

func testTest(t *testing.T, in string, tt test) {
	uuid, err := Parse(in)
	if ok := (err == nil); ok != tt.isuuid {
		t.Errorf("Parse(%s) got %v expected %v\b", in, ok, tt.isuuid)
	}
	if err != nil {
		return
	}

	if v := uuid.Variant(); v != tt.variant {
		t.Errorf("Variant(%s) got %d expected %d\b", in, v, tt.variant)
	}
	if v := uuid.Version(); v != tt.version {
		t.Errorf("Version(%s) got %d expected %d\b", in, v, tt.version)
	}
}

func TestUUID(t *testing.T) {
	for _, tt := range tests {
		testTest(t, tt.in, tt)
		testTest(t, strings.ToUpper(tt.in), tt)
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := Must(Parse("f47ac10b-58cc-0372-8567-0e02b2c3d479"))
		if uuid == Nil {
			b.Fatal("invalid uuid")
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkUUID_String(b *testing.B) {
	uuid := Must(Parse("f47ac10b-58cc-0372-8567-0e02b2c3d479"))
	if uuid == Nil {
		b.Fatal("invalid uuid")
	}
	for i := 0; i < b.N; i++ {
		if uuid.String() == "" {
			b.Fatal("invalid uuid")
		}
	}
}

func BenchmarkUUID_URN(b *testing.B) {
	uuid := Must(Parse("f47ac10b-58cc-0372-8567-0e02b2c3d479"))
	if uuid == Nil {
		b.Fatal("invalid uuid")
	}
	for i := 0; i < b.N; i++ {
		if uuid.URN() == "" {
			b.Fatal("invalid uuid")
		}
	}
}

func TestFromBytes(t *testing.T) {
	b := []byte{
		0x7d, 0x44, 0x48, 0x40,
		0x9d, 0xc0,
		0x11, 0xd1,
		0xb2, 0x45,
		0x5f, 0xfd, 0xce, 0x74, 0xfa, 0xd2,
	}
	uuid, err := FromBytes(b)
	if err != nil {
		t.Fatalf("%s", err)
	}
	for i := 0; i < len(uuid); i++ {
		if b[i] != uuid[i] {
			t.Fatalf("FromBytes() got %v expected %v\b", uuid[:], b)
		}
	}
}
