replacements = {
  "ab" : "c",
  "ac" : "b",
  "bc" : "a",
  "ba" : "c",
  "ca" : "b",
  "cb" : "a"
};


memTable = {}
def StringReduction(s: str) -> int:
    return len(StringReductionRecursive(s))


def StringReductionRecursive(s: str) -> str:
    print(s)
    if len(s) == 1:
        return s

    if s in replacements:
        return replacements[s]
    
    val1 = StringReductionRecursive(s[:int(len(s)/2)])
    val2 = StringReductionRecursive(s[int(len(s)/2):len(s)]) 

    if len(val1) == 2 and len(val2) == 2:
        return val2
    elif len(val1) == 2 and val1[0] != val2[0]:
        return replacements[val1[0] + StringReductionRecursive(val2[0]+val1[0])]
    elif len(val2) == 2 and val1[0] != val2[0]:
        return replacements[StringReductionRecursive(val1[0]+val2[0]) + val2[0]]

    return val1 + val2

print(StringReductionRecursive("bccababbbccbabba"))
