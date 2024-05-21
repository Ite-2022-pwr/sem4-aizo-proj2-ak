package utils

// HeapSort sortuje tablicę metodą sortowania przez kopcowanie
func HeapSort[T any](array []T, cmp func(a, b T) bool) {
  buildHeap(array, cmp)

  for i := len(array) - 1; i > 0; i-- {
    array[0], array[i] = array[i], array[0]
    heapify(array, i, 0, cmp)
  }
}

// buildHeap buduje kopiec z tablicy
func buildHeap[T any](array []T, cmp func(a, b T) bool) {
  n := len(array)
  for i := n / 2; i >= 0; i-- {
    heapify(array, n, i, cmp)
  }
}

func heapify[T any](array []T, size, root int, cmp func(a, b T) bool) {
  left, right := 2 * root + 1, 2 * root + 2
  largest := root

  if left < size && cmp(array[left], array[largest]) {
    largest = left
  }

  if right < size && cmp(array[right], array[largest]) {
    largest = right
  }

  if largest != root {
    array[root], array[largest] = array[largest], array[root]
    heapify(array, size, largest, cmp)
  }
}
