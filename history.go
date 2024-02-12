package main

type History struct {
	entries []string
	index   int
}

func (h *History) ResetIndex() {
	h.index = len(h.entries)
}

func (h *History) IsEmpty() bool {
	return len(h.entries) == 0
}

func (h *History) isAtTop() bool {
	return len(h.entries) == h.index+1
}

func (h *History) Push(str string) {
	h.entries = append(h.entries, str)
	h.index += 1
}

func (h *History) Next() (string, bool) {
	if h.isAtTop() {
		return "", false
	}

	if h.index < len(h.entries)-1 {
		h.index += 1

	}

	return h.entries[h.index], true
}

func (h *History) Pop() (string, bool) {

	if h.IsEmpty() {
		return "", false
	}

	if h.index > 0 {
		h.index -= 1
	}

	return h.entries[h.index], true

}
