# appointy

The task is to develop a basic version of meeting scheduling API. You are only required to develop
the API for the system. Below are the details.<br>

### Meetings should have the following Attributes.<br>
All fields are mandatory unless marked optional: <br>
● Id<br>
● Title<br>
● Participants<br>
● Start Time<br>
● End Time<br>
● Creation Timestamp<br>

### Participants should have the following Attributes.
All fields are mandatory unless marked optional:<br>
● Name<br>
● Email<br>
● RSVP (i.e. Yes/No/MayBe/Not Answered)<br>

### You are required to Design and Develop an HTTP JSON API capable of the following operations,<br>
● Schedule a meeting<br>
1.  Should be a POST request<br>
1.  Use JSON request body<br>
1.  URL should be ‘/meetings’<br>
1.  Must return the meeting in JSON format<br>

####  Get a meeting using id<br>
1.  Should be a GET request<br>
1  Id should be in the url parameter<br>
1.  URL should be ‘/meeting/<id here>’<br>
1.  Must return the meeting in JSON format<br>
  
#### List all meetings within a time frame<br>
1.  Should be a GET request<br>
1.  URL should be ‘/meetings?start=<start time here>&end=<end time here>’<br>
1.  Must return a an array of meetings in JSON format that are within the time range<br>
  
#### List all meetings of a participant<br>
1.  Should be a GET request<br>
1.  URL should be ‘/meetings?participant=<email id>’<br>
1.  Must return a an array of meetings in JSON format that have the participant received in the email within the time range<br>
  
### Additional Constraints/Requirements:<br>
● The API should be developed using Go.<br>
● MongoDB should be used for storage.<br>
● Only packages/libraries listed here and here can be used.<br>
Scoring:<br>
#### Completion Percentage<br>
1.  Total working endpoints among the ones listed above.<br>
1.  Meetings should not be overlapped i.e. one participant (uniquely identified by email) should not have 2 or more meetings with RSVP Yes with any overlap between their times.<br>
#### Quality of Code<br>
1.  Reusability<br>
1.  Consistency in naming variables, methods, functions, types<br>
1.  Idiomatic i.e. in Go’s style<br>
● Make the server thread safe i.e. it should not have any race conditions especially when two meetings are being booked simultaneously for the same participant with overlapping time.<br>
● Add pagination to the list endpoint<br>
● Add unit tests<br>
## Resources:
● Completing the Golang tour should give one a good grip over the language. Do this well and you will complete the task with ease.<br>
● This article should give you an idea on getting started with Web Application Development in Go.<br>
● This book covers both the workings of web and Go based servers.<br>
● This covers getting started with MongoDB in Go.<br>
