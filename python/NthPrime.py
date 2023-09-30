import math

def isPrime(n:int) -> bool:
    for i in range(2, int(math.sqrt(n))+1):
        if n % i == 0:
            return False
    return True


def nthPrime(n:int) -> int:
    primeCounter = 0
    i = 2
    while True:
        if isPrime(i):
            primeCounter += 1
            if primeCounter == n:
                return i
        i += 1

print(nthPrime(16))
