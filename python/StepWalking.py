

memTable = {}
# faboccini sequence
def StepWalking(n) -> int:
    if n == 0 or n == 1:
        return 1
    elif n in memTable:
        return memTable[n]
    else:
        val = StepWalking(n-1) + StepWalking(n-2)
        memTable[n] = val
        return val

print(StepWalking(5))

