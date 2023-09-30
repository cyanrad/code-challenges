
def RadwanArithGeo(arr): 
    # seq must at least be 3 ints long

    # arithmatic
    if arr[1] - arr[0] == arr[2] - arr[1]:
        factor = arr[1] - arr[0]
        for i in range(2, len(arr)):
            if arr[i] - arr[i-1] != factor:
                return -1
        return "Arithmatic"

    # geometric
    # divison by zero check should be in place
    if arr[1] / arr[0] == arr[2] / arr[1]:
        factor = arr[1] / arr[0]
        for i in range(2, len(arr)):
            if arr[i] / arr[i-1] != factor:
                return -1
        return "Geometric"

    return -1

def ArithGeo(arr):
    arith,geo = True, True
    for num in range(1, len(arr) - 1):
        if arr[num + 1] - arr[num] != arr[num] - arr[num-1]:
            arith = False
        if arr[num +1] / arr[num] != arr[num] / arr[num-1]:
            geo = False
    return "Arithmetic" if arith else "Geometric" if geo else -1

print(ArithGeo([2, 0, 18, 54]))
