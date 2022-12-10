package main
//Table: Customers 
//
// 
//+-------------+---------+
//| Column Name | Type    |
//+-------------+---------+
//| id          | int     |
//| name        | varchar |
//+-------------+---------+
//id is the primary key column for this table.
//Each row of this table indicates the ID and name of a customer.
// 
//
// 
//
// Table: Orders 
//
// 
//+-------------+------+
//| Column Name | Type |
//+-------------+------+
//| id          | int  |
//| customerId  | int  |
//+-------------+------+
//id is the primary key column for this table.
//customerId is a foreign key of the ID from the Customers table.
//Each row of this table indicates the ID of an order and the ID of the 
//customer who ordered it.
// 
//
// 
//
// Write an SQL query to report all customers who never order anything. 
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
//Customers table:
//+----+-------+
//| id | name  |
//+----+-------+
//| 1  | Joe   |
//| 2  | Henry |
//| 3  | Sam   |
//| 4  | Max   |
//+----+-------+
//Orders table:
//+----+------------+
//| id | customerId |
//+----+------------+
//| 1  | 3          |
//| 2  | 1          |
//+----+------------+
//Output: 
//+-----------+
//| Customers |
//+-----------+
//| Henry     |
//| Max       |
//+-----------+
// 
//
// Related Topics Database 👍 1767 👎 100


//There is no code of Go type for this problem

# SELECT A.Name from Customers A
# WHERE NOT EXISTS (SELECT 1 FROM Orders B WHERE A.Id = B.CustomerId)

# SELECT A.Name from Customers A
# LEFT JOIN Orders B on  a.Id = B.CustomerId
# WHERE b.CustomerId is NULL

# SELECT A.Name from Customers A
# WHERE A.Id NOT IN (SELECT B.CustomerId from Orders B)

select Name as Customers from Customers
where id not in (select customerId from orders)