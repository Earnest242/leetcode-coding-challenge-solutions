# Baseball challenge

## Instructions
You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:

An integer x.<br/>
Record a new score of x.<br/>
'+'.<br/>
Record a new score that is the sum of the previous two scores.<br/>
'D'.<br/>
Record a new score that is the double of the previous score.<br/>
'C'.<br/>
Invalidate the previous score, removing it from the record.<br/>
Return the sum of all the scores on the record after applying all the operations.<br/>

The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.<br/>

 

Example 1:<br/>

Input: ops = ["5","2","C","D","+"]<br/>
Output: 30<br/>
Explanation:<br/>
"5" - Add 5 to the record, record is now [5].<br/>
"2" - Add 2 to the record, record is now [5, 2].<br/>
"C" - Invalidate and remove the previous score, record is now [5].<br/>
"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].<br/>
"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].<br/>
The total sum is 5 + 10 + 15 = 30.<br/>
Example 2:<br/>

Input: ops = ["5","-2","4","C","D","9","+","+"]<br/>
Output: 27<br/>
Explanation:<br/>
"5" - Add 5 to the record, record is now [5].<br/>
"-2" - Add -2 to the record, record is now [5, -2].<br/>
"4" - Add 4 to the record, record is now [5, -2, 4].<br/>
"C" - Invalidate and remove the previous score, record is now [5, -2].<br/>
"D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].<br/>
"9" - Add 9 to the record, record is now [5, -2, -4, 9].<br/>
"+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].<br/>
"+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].<br/>
The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.<br/>
Example 3:<br/>

Input: ops = ["1","C"]<br/>
Output: 0<br/>
Explanation:<br/>
"1" - Add 1 to the record, record is now [1].<br/>
"C" - Invalidate and remove the previous score, record is now [].<br/>
Since the record is empty, the total sum is 0.<br/>
 

Constraints:<br/>

1 <= operations.length <= 1000<br/>
operations[i] is "C", "D", "+", or a string representing an integer in the range [-3 * 104, 3 * 104].<br/>
For operation "+", there will always be at least two previous scores on the record.<br/>
For operations "C" and "D", there will always be at least one previous score on the record.<br/>
