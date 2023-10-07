"""
Have the function TreeConstructor(strArr) take the array of strings stored in strArr,
which will contain pairs of integers in the following format: (i1,i2),
where i1 represents a child node in a tree and the second integer i2 signifies 
that it is the parent of i1.

Your program should, in this case, return the string true  because a valid binary tree can be formed.
If a proper binary tree cannot be formed with the integer pairs, then return the string false .
"""

def parse(s:str) -> list[str]:
    return s[1:-1].split(",")

def TreeConstructor(strs: list[str]) -> bool:
    nodes = {}
    hasParents = set()
    for rawStr in strs:
        s = parse(rawStr)

        if s[0] in hasParents:
            return False

        if s[1] in nodes:
            if len(nodes[s[1]]) == 2:
                return False
            else:
                nodes[s[1]].append(s[0])
        else:
            nodes[s[1]] = [s[0]]

        hasParents.add(s[0])
    return True

print(TreeConstructor(["(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"]))
print(TreeConstructor(["(1,2)", "(3,2)", "(2,12)", "(5,2)"]))
print(TreeConstructor(["(1,2)", "(3,2)", "(2,12)"]))
