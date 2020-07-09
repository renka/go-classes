Postman collection for testing: 
`https://www.getpostman.com/collections/0342583d8c136c090d42`

To run first part: 
1. go to directory `classes-api/cmd/admin` and run command `go run *.go`
2. Application starts on port `8081`

There is possibility to view, create and delete courses.
I have chosen the next way of data structuring, despite it looks a bit complicated, but storing data this way gives more flexibility and allows to cancel one class instead of whole course and gives quicker way to book class for one specific date.  
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

To run second part: 
1. go to directory `classes-api/cmd/user` and run command `go run *.go`
2. Application starts on port `8082`

This allows user to create booking using Username. Class with specific id can be booked only once for the same name. 
User also can to review all bookings for his name and delete specific booking. 
As far as there is no user authentication - no user validation was implemented. 
To get list of available classes need to call endoint from admin application.
If class was deleted - bookings for this class still in the system unless user deletes them

There were also added some unit tests - to check data validation logic. 
There is not much logic to test for now, so tests were added mostly to show working with Go testing framework