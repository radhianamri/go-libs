
## New
```go
    scanner, err := scanner.New("./data.txt", 65536)
	if err != nil {
		log.Fatal(err)
	}
	defer scanner.Close()

	for scanner.Scan() {
        fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
```

## NewLineProcesser
```go
   trimSpaceFunc := func(s *string) error {
		*s = strings.TrimSpace(*s)
		if *s == "" {
			return errors.New("empty line")
		}
		return nil
	}

	getFirstWordFunc := func(s *string) error {
		tokens := strings.Split(*s, " ")
		if len(tokens) < 2 {
			return errors.New("missing words")
		}
		*s = tokens[1]
		return nil
	}

	s, err := scanner.NewLineProcesser("./data.txt", 65536, trimSpaceFunc, getFirstWordFunc)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
```