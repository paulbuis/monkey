let fact = fn(n) { if (n == 0) { 1 } else { n * fact(n - 1) } };
fact(6);