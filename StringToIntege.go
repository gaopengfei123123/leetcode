package leetcode

// 同样的代码跑了两次排名就不一样...
func myAtoi(str string) int {
	arr := []rune(str)
	length := len(arr)

	sign := false
	begin := false
	result := 0
	zero := int(rune('0'))
	nine := int(rune('9'))
	del := int(rune('-'))
	plus := int(rune('+'))
	kong := int(rune(' '))
	max := int(^uint32(0) >> 1)
	min := ^max

	for i := 0; i < length; i++ {
		current := int(arr[i])
		if current >= zero && current <= nine {
			begin = true
			result = result*10 + current - zero
			if i > 0 && int(arr[i-1]) == del {
				sign = true
			}
			if result > max {
				break
			}
		} else {
			if !begin && current == kong {
				continue
			}
			if !begin && (current == del || current == plus) {
				begin = true
				continue
			}
			break
		}
	}

	if sign {
		result *= -1
		if result < min {
			result = min
		}
	} else {
		if result > max {
			result = max
		}
	}
	return result
}
