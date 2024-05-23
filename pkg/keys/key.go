package keys

import "fmt"

func GenerateKey(id1, id2 int) string {
	return fmt.Sprintf("%d:%d", id1, id2)
}

func ExtractIds(key string) (int, int) {
	var id1, id2 int
	fmt.Sscanf(key, "%d:%d", &id1, &id2)
	return id1, id2
}
