package _for

func Repeat(s string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + s
	}
	return repeated
}
