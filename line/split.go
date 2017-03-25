package line

func split(inputText string) (string, string) {
	length := len(inputText)

	switch length {
	case 0:
		return "", ""
	case 1:
		return string(inputText[0]), ""
	case 2:
		return string(inputText[0:1]), ""
	default:
		return string(inputText[0:1]), string(inputText[2:length])
	}
}
