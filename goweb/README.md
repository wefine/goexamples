# go-web
Source code for the book "Web Development with Go" published by Apress.

## TaskManager RESTful API
Sample REST API app with Go and MongoDB for chapter 9.

### Updates on 22-Jun-2016 in TaskManager App
1. Upgraded the package jwt-go and modified the source in common/auth.go.
2. Authentication middleware now sets the user name into the HTTP Request context using package "github.com/gorilla/context".
3. Added a logging implementation in common/logger.go. 
4. Minor code refactoring in the app.
