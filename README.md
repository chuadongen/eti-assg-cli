# ETI Assignment 1 (Master)

Name: Chua Dong En<br />
Class: P03<br />
ID: S10202623A<br />

## Contents

1. [Repositories](#Repositories)
2. [Requirements and Analysis](#Requirements and Analysis)
3. [Solution Architecture](#Solution-Architecture)
4. [Startup Guide](#Startup-Guide)

This assignment is a microservice based ride-sharing platform which allows passengers to book trips and riders to start/end trips.

## Repositories

| No        | Service Name           | Purpose  | Link  |
| :------------- |:-------------| :-----| :-----|
| 1 | eti-assg-cli (current) | Command line interface that allows both passengers and drivers to use the platform. | [Link](https://github.com/chuadongen/eti-assg-cli) |
| 2 | eti-assg-acc-svc | Account microservice handling data of users. Uses REST. | [Link](https://github.com/chuadongen/eti-assg-acc-svc) |
| 3 | eti-assg-trip-svc | Trips microservice handling data of trips. Uses REST. | [Link](https://github.com/chuadongen/eti-assg-trip-svc) |
| 4 | eti-assg-mysql-db | MySQL for persistant data storage. | [Link](https://github.com/chuadongen/eti-assg-mysql-db) |

## Requirements and Analysis

### Requirements

> During creation of passenger account, first name, last name, mobile number, and email address are required. Subsequently, users can update any information in their account, but they are not able to delete their accounts for audit purposes.

1. Creation of passenger account (First Name, Last Name, Mobile Number, Email Address)
2. Update passenger account information (First Name, Last Name, Mobile Number, Email Address)

> For driver account creation, first name, last name, mobile number, email address, identification number and car license number are required. Drivers can update all information except their identification number. Similarly, a driver account cannot be deleted.
3. Creation of driver account (First Name, Last Name, Mobile Number, Email Address, Identification Number, Car License Number)
4. Update driver account information (First Name, Last Name, Mobile Number, Email Address, Car License Number)

> A passenger can request for a trip with the postal codes of the pick-up and drop-off location. The platform will assign an available driver, who is not driving a passenger, to the trip. This driver will then be able to initiate a start trip or end trip. The passenger can retrieve all trips he/she has taken before in reverse chronological order
5. Passenger request trip (Pick up postal code, Drop off postal code) > Assigns available driver
6. Driver can start trip and end trip
7. Passenger can retrieve all trips in reverse chronological order

> 3.2.1.	Minimum 2 microservices using Go
8. 2 Microservices

> 3.2.2.	Persistent storage of information using database, e.g. MySQL
9. MySQL Database
 
### Analysis

### Driver needs to have availablility indication

Since the system will be automatically assigning available drivers to trips, there is a need to track the availability of drivers

### Drivers are able to reject trips

Drivers should be able to reject trips in the event of unforeseen circumstances

### Drivers are also able to retrieve all trips in reverse chronological order

Drivers should be able to view all their trip information as well

#### Assumptions/Constraints made
1. Drivers don't log off during a trip
2. Passengers cannot cancel trips

## Solution Architecture
### Workflow Diagram

![image](https://user-images.githubusercontent.com/73124349/208286968-54ffd958-6faf-4283-b70d-70c51e520aff.png)
DriverStatus accounts for the availability of drivers

### ER Diagram

![image](https://user-images.githubusercontent.com/73124349/208286489-b72628f1-dee6-4010-86c8-67733352374e.png)

### Architecture Diagram

![image](https://user-images.githubusercontent.com/73124349/208286536-56237076-ed28-44da-b533-a99a0532ab46.png)

In order to satisfy the 2 microservice requirements, it has been logically separated into acc-svc and trip-svc.

* **cli** - This is a command line interface for drivers and passengers to use the ride sharing service. It communicates with both of the microservices via rest.

* **acc-svc** - This is this accounts microservice that allows passengers and drivers to create, login and update their accounts

* **trip-svc** - This is this trips microservice that allows passengers and drivers to create and update trips.

* **mysql-database** - To satisfy the requiremetns, a mysql database was setup for persistent storage. The schema is shown in the ER diagram above.

## Startup Guide

## Setup Database
Clone [eti-assg-mysql-db](https://github.com/chuadongen/eti-assg-mysql-db)
Open up MySQL Workbench and run the init script

## Setup acc-svc
Clone [eti-assg-acc-svc](https://github.com/chuadongen/eti-assg-acc-svc)
Ensure that your console is in the directory of the acc-svc before running the following command

```
go run main.go
```

## Setup trip-svc
Clone [eti-assg-trip-svc](https://github.com/chuadongen/eti-assg-trip-svc)
Ensure that your console is in the directory of the trip-svc before running the following command

```
go run main.go
```

## Setup CLI

Clone [this repository](https://github.com/chuadongen/eti-assg-cli)
Ensure that your console is in the directory of the assg-cli before running the following command
```
go run menu.go
```
