def Consecutive(arr):
    largest, smallest = 0, 0
    if arr[0] > arr[1]:
        largest, smallest = arr[0], arr[1]
    else:
        largest, smallest = arr[1], arr[0]

    # O(n): time, O(1): space
    for x in arr:
        if x > largest:
            largest = x
        elif x < smallest:
            smallest = x
    return largest - smallest + 1 - len(arr)

print(Consecutive([5,10,15]))
print(Consecutive([-2,10,4]))
