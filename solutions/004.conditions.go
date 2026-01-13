package solutions

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	if age <= 0 || balance < 0 {
		return "INVALID"
	}
	if age < 18 || hasID == false {
		return "REJECTED"
	}
	if isVIP == true && balance >= 10000 {
		return "VIP_ACCESS"
	}
	if balance >= 1000 {
		return "STANDARD_ACCESS"
	}
	return "LIMITED_ACCESS"
}

func EvaluateGrade(score int) string {
	switch {
	case score < 0 || score > 100:
		return "INVALID"
	case score <= 100 && score >= 90:
		return "A"
	case score <= 89 && score >= 80:
		return "B"
	case score <= 79 && score >= 70:
		return "C"
	case score <= 69 && score >= 60:
		return "D"
	case score <= 59:
		return "F"
	}
	return "INVALID"
}
