
Problem 1  
-----------  
Function tellStory() has no tests  
Examine input & output  
Input is a gopher  
Output is a string  
Write a test that feeds some gophers into tellStory()  
Assert that a story about the gopher is returned  
  
Problem 2  
-----------  
Function getGopher() gets a gopher from a real database  
We want our test to run independently of that database  
Mock the database using the dynamodbiface.DynamoDBAPI package  
This will let the tests run independently of the real database.  
  
Problem 3  
-----------  
The randomStory() function gets a random story based on the time.  
This means the result of randomStory() is not always going to be the same.  
Mock the time so that randomStory() can be tested at a set moment in time.  
  
Problem 4  
-----------  
The gopher-story code was not written TDD style.  
Use TDD to write the code again.  
Is it better?  
Has it's testability improved?  