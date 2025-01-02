package advent

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
  r := make([]T2, len(s))
  for i, v := range s {
    r[i] = f(v)
  }
  return r
}

func Filter[T any](s []T, f func(T) bool) []T {
  var r []T
  for _, v := range s {
    if f(v) {
      r = append(r, v)
    }
  }
  return r
}

func Foldl[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
  r := init
  for _, v := range s {
    r = f(v, r)
  }
  return r
}
