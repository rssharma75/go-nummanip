package calc

func Add(numbers ... int) (error,int) {
	
	if len(numbers < 2)
		return errors.New("Need more than 2 numbers"),0
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}
	return nil,sum
}
