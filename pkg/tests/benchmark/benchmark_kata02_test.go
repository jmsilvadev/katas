package benchmark_test

import (
	"testing"

	"github.com/jmsilvadev/cycloid/katas02/iterable"
	"github.com/jmsilvadev/cycloid/katas02/recursive"
	recursivereference "github.com/jmsilvadev/cycloid/katas02/recursive_reference"
	slicerecursive "github.com/jmsilvadev/cycloid/katas02/slice_recursive"
)

func benchmarkIterable(s int, l []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		iterable.Chop(s, l)
	}
}

func benchmarkRecursive(s int, l []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		recursive.Chop(s, l)
	}
}

func benchmarkRecursiveReference(s int, l []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		recursivereference.Chop(s, l)
	}
}

func benchmarkSlicing(s int, l []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		slicerecursive.Chop(s, l)
	}
}

func BenchmarkIterable1(b *testing.B) { benchmarkIterable(1, []int{1, 3, 5, 7}, b) }
func BenchmarkIterable2(b *testing.B) { benchmarkIterable(3, []int{1, 3, 5, 7}, b) }
func BenchmarkIterable3(b *testing.B) { benchmarkIterable(5, []int{1, 3, 5, 7}, b) }
func BenchmarkIterable4(b *testing.B) { benchmarkIterable(7, []int{1, 3, 5, 7}, b) }

func BenchmarkRecursive1(b *testing.B) { benchmarkRecursive(1, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursive2(b *testing.B) { benchmarkRecursive(3, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursive3(b *testing.B) { benchmarkRecursive(5, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursive4(b *testing.B) { benchmarkRecursive(7, []int{1, 3, 5, 7}, b) }

func BenchmarkRecursiveReference1(b *testing.B) { benchmarkRecursiveReference(1, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursiveReference2(b *testing.B) { benchmarkRecursiveReference(3, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursiveReference3(b *testing.B) { benchmarkRecursiveReference(5, []int{1, 3, 5, 7}, b) }
func BenchmarkRecursiveReference4(b *testing.B) { benchmarkRecursiveReference(7, []int{1, 3, 5, 7}, b) }

func BenchmarkSlicing1(b *testing.B) { benchmarkSlicing(1, []int{1, 3, 5, 7}, b) }
func BenchmarkSlicing2(b *testing.B) { benchmarkSlicing(3, []int{1, 3, 5, 7}, b) }
func BenchmarkSlicing3(b *testing.B) { benchmarkSlicing(5, []int{1, 3, 5, 7}, b) }
func BenchmarkSlicing4(b *testing.B) { benchmarkSlicing(7, []int{1, 3, 5, 7}, b) }
