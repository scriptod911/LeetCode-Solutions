"""
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

 

Example 1:

Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input: s = "aab", p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input: s = "mississippi", p = "mis*is*p*."
Output: false
 

Constraints:

1 <= s.length <= 20
1 <= p.length <= 30
s contains only lowercase English letters.
p contains only lowercase English letters, '.', and '*'.
It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
"""


func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}
	if p[1] != '*' {
		if len(s) == 0 || (s[0] != p[0] && p[0] != '.') {
			return false
		}
		return isMatch(s[1:], p[1:])
	}
	if isMatch(s, p[2:]) {
		return true
	}
	if len(s) == 0 || (s[0] != p[0] && p[0] != '.') {
		return false
	}
	return isMatch(s[1:], p)
}

// how to remove the third occurrence of a duplicate line in regex? 
// -> ^(?!.*(.).*\1).*$ 

// What is $[\s\S]* going to do?
// -> it will match any string of characters, including newlines. 