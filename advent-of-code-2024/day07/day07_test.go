package main

import("testing")

func BenchmarkMain(b *testing.B) {
  for i := 0; i < b.N; i++ {
    solve(mult, plus, conc)
  }
  b.ReportAllocs()
}
