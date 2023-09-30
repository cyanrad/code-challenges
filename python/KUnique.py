def KUnique(str):
    k = int(str[0])
    count = {}
    head, tail = 1, 1
    longest = (0,0)

    
    # O(n): time, O(k): memory
    while tail < len(str):  
        if len(count) <= k:
            if str[tail] not in count:
                count[str[tail]] = 1
            else:
                count[str[tail]] += 1

            if len(count) > k and longest[1] - longest[0] < tail-head:
                longest = (head, tail)

            tail += 1

        else:
            count[str[head]] -= 1
            if count[str[head]] == 0:
                count.pop(str[head])

            head += 1

    if longest[1] - longest[0] < tail-head:
        longest = (head, tail)

    return str[longest[0]:longest[1]]

print(KUnique("2aabbcbbbadef"))
