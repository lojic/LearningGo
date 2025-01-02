package main

// run benchmark via:  go test -bench=.
import ("testing")

func BenchmarkMain(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Solve()
  }
  b.ReportAllocs()
}
