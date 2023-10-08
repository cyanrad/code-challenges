def topKFrequent(nums, k):
    counter = {}
    for n in nums:
        counter[n] = counter.get(n, 0)+1

    frequency = []
    for _ in range(len(nums)):
        frequency.append([])

    for key, value in counter.items():
        frequency[value-1].append(key)


    result = []
    finished = False
    for i in range(len(nums)-1, -1, -1):
        for j in range(len(frequency[i])-1, -1, -1):
            result.append(frequency[i][j])
            if len(result) == k:
                finished = True
                break
        if finished:
            break
    return result




print(topKFrequent([1,1,1,2,2,3], 2))
print(topKFrequent([1], 1))