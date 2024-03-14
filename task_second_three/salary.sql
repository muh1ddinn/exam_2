
Create table If Not Exists Employee (id int, name varchar(255), salary int, departmentId int)
Create table If Not Exists Department (id int, name varchar(255))
Truncate table Employee
INSERT INTO Employee(id, name, salary, departmentId) VALUES (1, 'Joe', '70000', 1), (3, 'Henry', '80000', 2),
 ('4', 'Sam', '60000', '2'),
 ('5', 'Max', '90000', '1');
Truncate table Department
INSERT INTO Department (id, name) VALUES (1, 'IT'),(2, 'Sales')





SELECT 
    d.name AS Department, 
    e.name AS Employee, 
    e.salary AS Salary
FROM Employee e                                                  
JOIN Department d ON e.departmentId = d.id
WHERE NOT EXISTS (
    SELECT 1 
    FROM Employee e2 
    WHERE e2.departmentId = e.departmentId 
      AND e2.salary > e.salary
)
ORDER BY Department;
 department | employee | salary 
------------+----------+--------
 IT         | Joe      |  70000
 Sales      | Henry    |  80000
(2 rows)
