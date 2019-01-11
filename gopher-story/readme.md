
Problem 1  
-----------  
function tellStory() has no testsExamine input & output  
input is a gopher  
output is a string  
write a test that feeds some gophers into tellStory()  
and assert that a story about the gopher is returned  
  
Problem 2  
-----------  
function getGopher() gets a gopher from a real database  
we want our test to run independently of that database  
mock the database using the dynamodbiface.DynamoDBAPI package  
so that the test runs independently of the real database.  
  
Problem 3  
-----------  
The randomStory() function gets a random story based  
on the time. So the result will be different each time  
