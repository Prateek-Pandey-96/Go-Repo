This is a tasks api with simple logic but not so simple implementation

Functional Features :-

1> Create/Read/Update/Delete users
2> Create/Read/Update for tasks

What makes it a bit different from simple crud apps is:-

1> The user controller is fully unit tested with go-sqlmock
2> Redis is integrated to fetch tasks for a user efficiently 
3> Stale data was not an issue for my use case and the cache is invalidated every 15 minutes


Api is fully dockerized and can be used as a standalone microservice