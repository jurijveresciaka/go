package math

func IsPrime(n uint64) bool {
  // 0 and 1 not prime.
  if n < 2 {
    return false
  }

  // all even numbers except 2 are no prime.
  if n != 2 && n % 2 == 0 {
    return false
  }

  var i uint64

  for i = 2; i * i <= n ; i++ {
    if n % i == 0 {
      return false
    }
  }

  return true
}
