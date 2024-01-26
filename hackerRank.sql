-- first test solution
SELECT
    roll_number,
    name
FROM
    student_information
WHERE advisor IN (
    SELECT
        employee_id
    FROM 
        faculty_information
    WHERE (gender = 'M' AND salary > 15000)
       OR (gender = 'F' AND salary > 20000)
);

-- second test solution
SELECT
    today.stock_code
FROM
    price_today today
JOIN
    price_tomorrow tomorrow ON today.stock_code = tomorrow.stock_code
WHERE 
    today.price < tomorrow.price
    AND (
        today.stock_code = 'TITAN'
        OR today.stock_code = 'MRF'
        OR today.stock_code = 'GOOGL'
    );