package letterinventory

type LetterInventory struct {
	count [26]int
	size  int
}

func CreateLetterInventory(str string) LetterInventory {
	letterInventory := LetterInventory{}
	letterInventory.init(str)
	return letterInventory
}

func (li *LetterInventory) init(str string) {
	for _, char := range []byte(str) {
		index := char - 97
		if index >= 0 && index <= 25 {
			li.count[index]++
			li.size++
		}
	}
}

func (li *LetterInventory) Add(letterInventory LetterInventory) {
	for key, value := range letterInventory.Count() {
		li.count[key] += value
		li.size += value
	}
}

func (li *LetterInventory) Subtract(letterInventory LetterInventory) {
	for key, value := range letterInventory.Count() {
		li.count[key] -= value
		li.size -= value
	}
}

func (li LetterInventory) Count() [26]int {
	return li.count
}

func (li LetterInventory) Size() int {
	return li.size
}

func (li LetterInventory) ContainsNegative() bool {
	for _, value := range li.count {
		if value < 0 {
			return true
		}
	}
	return false
}

func (li LetterInventory) Contains(letterInventory LetterInventory) bool {
	for i := range letterInventory.count {
		if letterInventory.count[i] > li.count[i] {
			return false
		}
	}
	return true
}
