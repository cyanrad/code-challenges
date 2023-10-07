"""
have the function RemoveBrackets(str) take the str string parameter being passed,
which will contain only the characters "(" and ")", and determine the minimum number
of brackets that need to be removed to create a string of correctly matched brackets.
For example: if str is "(()))" then your program should return the number 1.
The answer could potentially be 0,
and there will always be at least one set of matching brackets in the string.
"""

def RemoveBrackets(s):
    count = 0
    for c in s:
        if c == "(":
            count += 1
        else:
            count -= 1
    return abs(count)

print(RemoveBrackets("(())()))"))
