-- search date is exactly the same as existing reservation
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-02-01' < end_date and '2021-02-04' > start_date;
	
	
	
-- start date is before existing reservation, end date is same
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-01-31' < end_date and '2021-02-04' > start_date
	

-- end date is after existing reservation, start date is same
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-02-01' < end_date and '2021-02-07' > start_date
	
	

-- search dates are outside of all existing reservations, but cover the reservation
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-01-01' < end_date and '2021-03-07' > start_date
	
	
-- search dates are outside of all existing reservations
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-07-01' < end_date and '2021-07-07' > start_date
	
	
-- search dates are inside of all existing reservations
SELECT
	count(id)
FROM
	room_restrictions
WHERE
	'2021-02-02' < end_date and '2021-02-03' > start_date
	