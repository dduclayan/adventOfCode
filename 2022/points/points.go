package points

var (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func NewLettersToPointsMap() map[string]int {
	ltpMap := make(map[string]int)
	for i, v := range letters {
		i++
		ltpMap[string(v)] = i
	}
	return ltpMap
}
