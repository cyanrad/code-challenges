class Solution:
    def isPalindrome(self, s: str) -> bool:
        start, end = 0, len(s)-1
        while start <= end:
            if not s[start].isalnum():
                start += 1
            elif not s[end].isalnum():
                end -= 1
            elif s[end].lower() != s[start].lower():
                return False
            else:
                end -= 1
                start += 1
        return True

