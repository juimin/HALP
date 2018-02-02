# Contribution Guidelines

## Basic Rules
- Pull Requests are mandatory
- Code Reviews are mandatory
- Testing of your code is mandatory (so write your own tests)

## Starting Work Flow
1. CLONE the github repo
- If you are working on the server, checkout server and then do work with server as the base.
- Basically same rules for the client
- NO ONE should be merging anything to master

2. CHECKOUT the desired branch you want to start from
3. CREATE A NEW BRANCH to work on so you aren't interfering with someone else's branch
4. CODE your stuff
5. TEST your code
6. CREATE a pull request
7. ASK for a code review
8. REVIEWER will merge the code upon approval

## Pull Request Merge Requirements
- Code Coverage > 80%
- Build must be passing
- Ideally 2 Code Reviews (sometimes 1 will be sufficient)

## Flow
The project should do something like this so we don't get screwed. (Standard Dev path)

Master <- Staging <- Development

Development will be split into each client type (ios, android, web) and the server

Code pushed to a development branch will be merged into staging. If everything passes in staging, it should be automatically merged to master

![flow diagram](https://github.com/JuiMin/HALP/blob/master/development-flow.jpeg)

