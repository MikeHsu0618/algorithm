//Table: Users
//
// 
//+----------------+---------+
//| Column Name    | Type    |
//+----------------+---------+
//| user_id        | int     |
//| name           | varchar |
//+----------------+---------+
//user_id is the primary key for this table.
//This table contains the ID and the name of the user. The name consists of 
//only lowercase and uppercase characters.
// 
//
// 
//
// Write an SQL query to fix the names so that only the first character is 
//uppercase and the rest are lowercase. 
//
// Return the result table ordered by user_id. 
//
// The query result format is in the following example. 
//
// 
// Example 1: 
//
// 
//Input: 
//Users table:
//+---------+-------+
//| user_id | name  |
//+---------+-------+
//| 1       | aLice |
//| 2       | bOB   |
//+---------+-------+
//Output: 
//+---------+-------+
//| user_id | name  |
//+---------+-------+
//| 1       | Alice |
//| 2       | Bob   |
//+---------+-------+
// 
//
// Related Topics Database 👍 432 👎 65


//There is no code of Go type for this problem
select
user_id, concat(upper(left(name, 1)),lower(substr(name,2))) as name
from Users
order by user_id asc