package uu

const (
	//MaxBytesPerLine is maximum bytes allowed in per line
	MaxBytesPerLine = 45

	//PCharStart is integer representation of first printable character 'Space'
	UUCharStart = 0x20

	UUCharPseudoZero = 0x60

	PCharNewline = 0x0A

	PCharCR = 0x0D
)

func Conv3Bytes(aBytes [3]byte) [4]byte {
	lRet := [4]byte{}

	lInt32 := uint32(0)
	for cIdx := 0; cIdx < 3; cIdx++ {
		lInt32 <<= 8
		lInt32 |= uint32(aBytes[cIdx])
	}

	for cIdx := 0; cIdx < 4; cIdx++ {
		lRet[3-cIdx] = UUCharStart + byte(lInt32&0x3F)
		lInt32 >>= 6
	}

	return lRet
}

func EncodeLine(aBytes []byte) []byte {
	if len(aBytes) > MaxBytesPerLine {
		return nil
	}

	if len(aBytes) == 0 {
		return []byte{UUCharPseudoZero, PCharCR, PCharNewline}
	}

	lRet := make([]byte, 0)

	lRet = append(lRet, byte(UUCharStart+len(aBytes)))

	for cIdx := 0; cIdx < (len(aBytes) / 3); cIdx++ {
		bIn := [3]byte{}
		copy(bIn[:], aBytes[cIdx*3:((cIdx+1)*3)])
		bOut := Conv3Bytes(bIn)
		lRet = append(lRet, bOut[0], bOut[1], bOut[2], bOut[3])
	}

	{
		bLeft := len(aBytes) % 3
		if bLeft > 0 {
			bIn := [3]byte{}
			copy(bIn[0:bLeft], aBytes[len(aBytes)-bLeft:])
			bOut := Conv3Bytes(bIn)
			lRet = append(lRet, bOut[0], bOut[1], bOut[2], bOut[3])
		}
	}

	lRet = append(lRet, PCharCR, PCharNewline)

	return lRet
}
