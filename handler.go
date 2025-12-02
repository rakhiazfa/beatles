package beatles

type Handler = func(c Context) error

func mergeHandlers(args ...[]Handler) []Handler {
	totalHandlers := 0

	for _, handlers := range args {
		totalHandlers += len(handlers)
	}

	merged := make([]Handler, 0, totalHandlers)
	for _, handlers := range args {
		merged = append(merged, handlers...)
	}

	return merged
}
