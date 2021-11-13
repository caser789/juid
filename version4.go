package uuid

type UUID []byte

func NewRandom() UUID {
	uuid := make([]byte, 16)
	randomBits(uuid)
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}
