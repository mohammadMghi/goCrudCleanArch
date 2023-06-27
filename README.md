# goCrudCleanArch
This is a simple CRUD created with clean architecture


## What is clean architecture ?

Clean architecture is software design that separation of concerns , identify the Core Business Logic , Define Layers , Establish Dependency Flow , Implement Use Cases,Use Dependency Injection to achive
a software that approptrate for development and maintan.

Note : clean arch is a principle we should base on this principle implement our software design.

First of all we have to know about our logic , but what is meaning of logic here? In fact logic is proccess of working a organization (like bank , shop .etc) we went to automasion this flow with the computers and this proccess named logic of project.

How to understand logic of the project that we want to implement? It's a very good question . for do this we have many approach , in my opinion one of the best is DDD .
For understand of Domain Driven Design i'll write a new article in future.

## Sparation of concens: 

separation of concerns mean separate as much as possible for achieve loose coupling and SOLID principle .
For example database is one of the concerns and we have to prioritize database concerns based on future changes . for example two year later we want to change or add new databse , in this sutation we should implement or change the databse whithout chenging the logic of business . So we have to implement it with abstraction layer . abstraction layer can spreate these concerns.
