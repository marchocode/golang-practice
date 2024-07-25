package main

func longestCommonPrefix(strs []string) string {

	index := 0
	bytes := make([]byte, 0)

	for {

		var b byte
		var t bool = true

		for i, str := range strs {

			if index > len(str)-1 {
				t = false
				break
			}

			if i == 0 {
				b = str[index]
				continue
			}

			if b != str[index] {
				t = false
				break
			}

		}

		index++

		if t {
			bytes = append(bytes, b)
			continue
		}

		break
	}

	return string(bytes)
}
