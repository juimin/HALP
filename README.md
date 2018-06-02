# HALP

Capstone Project: Providing a mobile solution for community assistance on everyday problems

## Component Status

| Component              | Status                            |
| ----------------- |:---------------------------------:|
| Server            | [![Build Status](https://travis-ci.org/JuiMin/HALP.svg?branch=master)](https://travis-ci.org/JuiMin/HALP) [![Coverage Status](https://coveralls.io/repos/github/JuiMin/HALP/badge.svg?branch=master)](https://coveralls.io/github/JuiMin/HALP?branch=master)      |
| Android Client | Under Construction |
| iOS Client | Under Construction |
| Web Client | -- |
| Web Site   | [Halp App Site](https://halpapp.github.io/) |
| API Server | https://staging.halp.derekwang.net |

## Deployment Status

Staging: **Online**

## Project Overview
The purpose of this project is to develop a service that allows users to provide assistance to each other through a mobile first application that allows for heavy use of photos and photo editing. Users can post questions that include a photo which they either take immediately or upload from their phone. Upon taking or loading the photo, the user can edit the photo by drawing on it to enhance the image, provide clarity of something in the photo or concerning the problem, or draw focus to specific elements in the image. Users who wish to answer the question will be able to provide their own or further edits, which will aid them in providing assistance to the original poster.

For the server development, the majority of our code is based on the 344 Server Side development course in the Informatics Department at the University of Washington. The general server design scheme originates from the material contained in this course. As such, the language that we chose to implment the server in was Go, although we had also considered Python and Node.js as alternatives. Additionally, we considered the option of minimizing server development but this will be discussed more in the rationale section.

The client was developed using React Native, so that the application could be developed for both iOS devices and Android devices simultaneously. Most of our content was tested for android, and care was taken to avoid package components that indicated only android usage so that compatibility with both platforms could be maintained.

## Repository Contents
This reposity contains all the code used in both the back end server configuration and the client side mobile application. These files are separated into appropriately labeled folders, which are themselves the root directories of protected branches that individually track the client development and the server development. At the root level, the repository also contains issue templating, contribution guidelines, the MIT license, as well as the git ignore and configuration files for the Travis CI (Continuous Integration) Service.

Within the client directory, the standard React Native files generated with Create-React-Native-App are stored, which includes a ```/src``` directory containing code categorized into React Components, Redux Reducers, Styling sheets, required IMages, as well as predefined application constants and functions containing fetch HTTP Requests. React native has some special requirements to get up and running with the packages we required, and details on navigating these requirements are within the repository's wiki. Among these special details is information concerning git ignored files that are generated when using create react native app as well as the node modules folder containing our dependency code.

The server contains a collection of code files decribing a dockerized server container which utilizes mongodb and includes scripts to run the code on any server. The server is separated into package components, those being the handlers for http requests, models for data description in storage, tls and scripting, as well as indexes for trie search functionality.

## Technology Decisions

### Server Development
We elected to go for server development as opposed to using a managed backend service like firebase for a couple of reasons. The first among these was that after we discovered the modeling of the data to be relatively simplistic and the fact that we had some experience writing server side code, we could exert more control over the overall state of the application by making the back end ourselves. One advantage of this is that we eliminated a service as a dependency, requiring mainly the packages used to write the software as well as a hosting service with which to make the application available. Another reason is due to the design of our application intending for a large number of users to be serviced, leading us to reason that being able to control the entirety of the application would be beneficial when we eventually required scale ups, a neccessary consideration for most projects but absolutely vital for ours. Although this added more work to our plates in terms of actual tasks to complete, we believed that the resulting application would benefit as a whole in the long run. A personal benefit for some of our developers was that they got to learn more about server side development while others were able to solidify their current understanding.

### Database
The database management system that we opted to use for the current iteration of the HALP service was MongoDB. We chose a NoSQL database management system even though we had what can be considered a well defined server model due to MongoDB's ease of use, the ability to easily change the model easily during development, and the scalability promises that a document based storage system provides. In our research, we noted that the MongoDB system lacked transactions (a feature coming soon in a future version) and that companies like Quora and Reddit utilize SQL services when tracking data. As these are similar and competing applications, we also considered that MongoDB might not be the best option currently and there were advantages to be gained when using SQL. Eventually we decided to settle with MongoDB for the time being because most of the application was already using it and the structure of our code made it relatively easy to change the database system at a later time, so long a we did not deploy yet and we didn't have a large amount of data to migrate.

### Dockerization
Our application was intended for a large amount of users at any given time, with the community providing most of the content and available to other users. As such, high availability and scalability were both things to consider when developing our services. Docker containerization was used to package the server side code so that new instances could be spun up quickly, allowing the system to monitor itself and start up services or shut down services to accomodate load as needed. As most of the required services such as the main database and the session store database(redis) were also available through docker, the server could be easily scripted in order to achieve this scalability, as well as minimizing the need for developer control when starting a server.

### Cloud Service
Once we had decided on our implementation of the back end service, we had to consider how files would be stored and where data would live. The two main services that we looked at for addressing this were Digital Ocean and AWS, Although AWS had a lot of interesting perks and features that came along with their service, we found that we werent' sure what we really needed and found that digital ocean could have a lower upkeep cost when considering the tasks we required. Using Digital Ocean's droplets to contain the server code and spaces to act as the file storage system, we found that Digital Ocean was easy enough to use, could be easily scripted for scaling, and when we considered our Dockerization, migrating at a later date wouldn't be too big a deal so long as we did not have to do any large scale data migration.

### Client Development
Our intention for this application was to provide a mobile first platform with which the users would be able to create and answer posts with ease. For this, we considered mobile capable web applications which are more easily accessible but we decided to go for native code because this would also give us more control over the user experience. Developing a standalone application has the advantage of direct installation, and when we found React-Native had the capability of developing for both Android devices as well as iOS devices at the same time, we thought that we could produce a a more consistent user experience for all users regardless of what phone they were using while maintaining an application that ran on native code. In addition, most of our team was already familiar with both React and Material UI design, which made finding packages that satisfied our eventual User Interface design easier to accomplish. The only exception to this was the drawing library that we used (Canvas), which was the best open source option we could find for quickly implementing drawing over photo capabilities without having to write a complicated piece of software ourselves in the limited time that we had.

### Testing
In terms of testing, we tried to keep the server code tested as much as possible. For this reason, we generally aimed for at least 85% code coverage on unit tests, with several components tested for integration. We believed that the server was important to test as it allowed us to see functionality before deployment and allowed us to keep a good eye on what components required work. We recognized that code coverage is not the best metric for success when writing tests, but it was important that we made sure most of the code was tested in the first place. Our priority was still in writing good tests that tested function output as well as the overall action of a system.

For the client side, we actually didn't spend that much time writing unit tests since we were in a time crunch and were rushed to finish. If this application is ever deployed, testing for the client will have been done so as to give it a similar level of quality assurance as the server, mitigating any user dissatisfaction and technical debt.

## Development Team
| Name              | Role(s)                           | Contact        |
| ----------------- |:---------------------------------:|----------------|
| Alex Gilbert      | Product Management/Developer      | agilbe@uw.edu  |
| Davin Lee         | Project Management/Developer      | davinl@uw.edu  |
| Derek Wang        | Lead Developer                    | d95wang@uw.edu |
| Kanon Shibata     | Lead Designer                     | kshibata@uw.edu|

## Project Timeline
**[AppWeek Timeline (Read Only)](https://app.teamweek.com/#pg/6hGsJu7uJgUkuUMtogdRId_TRHJhxAar)**
