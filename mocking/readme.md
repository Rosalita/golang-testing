
Problem:
Your code interacts with something else, like a database. But when tests run, you don't want the tests to change the state of a real database by reading and writing to it.

Solution:
Mocking. By writing come code that acts like a database (a mock database) we can make the tests run using a mock database rather than the real database.