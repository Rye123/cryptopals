package aes

func circularRotateOnce(data []byte, dirLeft bool) []byte {
	result := make([]byte, len(data))
	if len(data) == 0 {
		return result
	}
	
	if dirLeft {
		leftmost := data[0]
		for i := 0; i < len(data); i++ {
			if i+1 == len(data) {
				result[i] = leftmost
			} else {
				result[i] = data[i+1]
			}
		}
	} else {
		rightmost := data[len(data)-1]
		result[0] = rightmost
		for i := 1; i < len(data); i++ {
			result[i] = data[i-1]
		}
	}
	return result
}

// Returns a new bytestring of `data` rotated by `offset` bytes. This is a circular rotation.
// Direction is determined by `dirLeft` -- if true, rotation is to the left, otherwise rotation is to the right.
func CircularRotate(data []byte, offset int, dirLeft bool) []byte {
	result := make([]byte, len(data))
	copy(result, data)
	for i := 0; i < offset; i++ {
		result = circularRotateOnce(result, dirLeft)
	}
	return result
}
