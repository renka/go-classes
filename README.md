To run first part: 
1. go to directory `classes-api/cmd/admin` and run command `go run *.go`

There is possibility to view, create and delete courses. 
Course - Entire class, means entity with name, start date and end date. It is not stored in database, because system splits it into small daily classes, one class per day. 
Class - entity, stored in database, it has exact date, name, id odf course and capacity. 
For example, request to `POST http://localhost:8081/class` looks like: 
`    {
         "name": "Functional workout",
         "start_date": "2020-07-09",
         "end_date": "2020-07-11",
         "capacity": 20
     }`
This means there will be 3 entities stored in database: 
 `{
      "id": 7,
      "name": "Functional workout",
      "capacity": 20,
      "class_id": "79Fun",
      "date": "2020-07-09"
  }
  {
      "id": 8,
      "name": "Functional workout",
      "capacity": 20,
      "class_id": "79Fun",
      "date": "2020-07-10"
  }
  {
      "id": 9,
      "name": "Functional workout",
      "capacity": 20,
      "class_id": "79Fun",
      "date": "2020-07-11"
  }`
Only one course with unique name can be stored for the same date. 
In terms of current example it is not possible to register another 'Functional workout' class on '2020-07-09' 
But it is possible to add 'Yoga' classes for the same dates. 
There is possibility to remove one specific date by id or entire course by class_id (known issue - if 2 courses have the same id - doth will be removed, it is not important for now)

