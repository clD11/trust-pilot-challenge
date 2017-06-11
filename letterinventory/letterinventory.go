package letterinventory

type LetterInventory struct {
	count [26]int
	size  int
}

func CreateLetterInventory(str string) LetterInventory {
	letterInventory := LetterInventory{}
	letterInventory.Add(str)
	return letterInventory
}

func (l *LetterInventory) Subtract(str string) {
	for _, char := range []byte(str) {
		index := char - 97
		if index >= 0 && index <= 25 {
			l.count[index]--
			l.size--
		}
	}
}

func (l *LetterInventory) Add(str string) {
	for _, char := range []byte(str) {
		index := char - 97
		if index >= 0 && index <= 25 {
			l.count[index]++
			l.size++
		}
	}
}

func (l *LetterInventory) Contains(str string) bool {
	letterInventory := CreateLetterInventory(str)
	for i := range letterInventory.count {
		if letterInventory.count[i] > l.count[i] {
			return false
		}
	}
	return true
}

func (l LetterInventory) GetCount() [26]int {
	return l.count
}

func (l LetterInventory) GetSize() int {
	return l.size
}
