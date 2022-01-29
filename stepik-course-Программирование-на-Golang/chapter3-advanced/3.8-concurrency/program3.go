package main

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)
	uniq := make([]string, 0)
	for str := range inputStream {
		inUniq := false
		for _, s := range uniq {
			if str == s {
				inUniq = true
			}
		}
		if !inUniq {
			uniq = append(uniq, str)
			outputStream <- str
		}
	}
}
