package main
//Table: Weather 
//
// 
//+---------------+---------+
//| Column Name   | Type    |
//+---------------+---------+
//| id            | int     |
//| recordDate    | date    |
//| temperature   | int     |
//+---------------+---------+
//id is the primary key for this table.
//This table contains information about the temperature on a certain day.
// 
//
// 
//
// Write an SQL query to find all dates' Id with higher temperatures compared 
//to its previous dates (yesterday). 
//
// Return the result table in any order. 
//
// The query result format is in the following example. 
//
// 
// Example 1: 
//
// 
//Input: 
//Weather table:
//+----+------------+-------------+
//| id | recordDate | temperature |
//+----+------------+-------------+
//| 1  | 2015-01-01 | 10          |
//| 2  | 2015-01-02 | 25          |
//| 3  | 2015-01-03 | 20          |
//| 4  | 2015-01-04 | 30          |
//+----+------------+-------------+
//Output: 
//+----+
//| id |
//+----+
//| 2  |
//| 4  |
//+----+
//Explanation: 
//In 2015-01-02, the temperature was higher than the previous day (10 -> 25).
//In 2015-01-04, the temperature was higher than the previous day (20 -> 30).
// 
//
// Related Topics Database 👍 1313 👎 432


//There is no code of Go type for this problem

// select a.id from  Weather as a
// join Weather as b
// on a.Temperature > b.Temperature
// and dateDiff(a.recordDate, b.recordDate) = 1

select a.id
from weather as a, weather as b
where a.Temperature > b.Temperature
and datediff(a.recordDat, b.recordDat) = 1