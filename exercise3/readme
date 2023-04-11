Write a function:
  func Solution(s string, T string) string

that given two strings S and T consisting of N and M characters, respectively,
determines whether string T can be obtained from string S by at most one
simple operation from the set specified below. The function should return a
string:

- "ADD c" if string T can be obtained from string S by inserting a single 
character "c" at the end of the string;
- "JOIN c" if string T can be obtained from string S by merging two adjacent
characters "c" into one (exactly one merge is performed);
- "SWAP c d" if string T can be obtained from string S swapping two characters
"c" and "d" (these characters should be distinct and in the same order as in
string S; exactly one swap is performed);
- "NOTHING" if no operation is needed (strings T and S are equal);
- "IMPOSSIBLE" if none of the above works.

Note that by characters "c" and "d" from the operations above, we mean any 
English alphabet lowercase letters.

For example:
- given S="nice" and T="nicer", the function should return "ADD r";
- given S="meet" and T="met", the function should return "JOIN e";
- given S="late" and T="tale", the function should return "SWAP l t";
- given S="o" and T="odd", the function should return "IMPOSSIBLE".

Assume that:
- N and M are integers within the range [1...100];
- string S is made only of lowercase letters (a-z);
- string T is made only of lowecase letters (a-z).

In your solution, focus on correctness. The perfomance of your solution will
not be the focus of the assesment. 
