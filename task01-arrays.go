package homework

func average(input [15]float32) (result float32) {

	var count float32

	for _, v := range input {
		count += v
	}
	result = count / 15

	return
}
