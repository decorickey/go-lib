package stream

// Filter は条件を満たす要素を抽出するジェネリクス対応のフィルタ関数。
func Filter[T any](values []T, fn func(v T) bool) []T {
	results := make([]T, 0, len(values))
	for _, v := range values {
		if fn(v) {
			results = append(results, v)
		}
	}
	return results
}
