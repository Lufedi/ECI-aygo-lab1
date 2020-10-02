# ECI-aygo-lab1


This is a small laboratory of the course "AYGO" at Escuela Colombiana de Ingeniería Julio Garavito and part of the Informatics's master degree.

#### Objectives

Create an e2e web application that saves strings and returns the 10 most recently created. Also I implemted an application load balancer to balance the load when saving the texts. 

The project was build using `docker-compose`, `go` and `react`, deployed on an `EC2` instance.

Demo: http://ec2-34-209-136-168.us-west-2.compute.amazonaws.com/


#### How to run locally



```
$ git clone https://github.com/Lufedi/ECI-aygo-lab1
$ docker-compose build
$ docker-compose up -d
```


**Warning** If you are running this on your local machine ignore this, otherwise if you are running this on a remote server, change the backend url to your server public domain or IP [here](https://github.com/Lufedi/ECI-aygo-lab1/blob/master/aygolaboneclient/src/App.js#L7). (EC2 domain i.e.)


#### Project structure

- *app-lb-round-robin*: Application load balancer using the RR algorithm. `go`
- *aygolabone*: Backend service that saves and retrieves the strings.`go`
- *aygolaboneclient*: Webapp to send and read the strings. `reactjs`
- *docker-compose*: Docker tool to orchestrate multiple containers in one project 

#### Arquitecture and design

The architecture is as described in the image bellow.



##### Overview

![](https://drive.google.com/uc?export=view&id=12cP4ywZjSRtorUBvFE-LCvTdShpqt0n8)



Three backend services (`aygolabone`) are running on containers at port 4000, those services domains where registered in the `app-lb-round-robin` project manually*, which is a *reversed proxy* and a *gateway* at the same time, the `app-lb-round-robin` picks one of the three services using the round robin algorithm and forwards the request to the selected service. 

This was achieved using an already implemented `Proxy` tool called [SingleHostReverseProxy](https://golang.org/pkg/net/http/httputil/#NewSingleHostReverseProxy)


(*) this step can be avoided using the built-in service discovery tool from docker and docker-compose, but one goal of this laboratory is study the RR load balancing algorithm.


##### The API and the webapp
The Webapp is a simple `react` app that calls the exposed `API` to ger and save the texts. The API signatures are:

GET `/api/text`

POST `api/test`
`{ value: "value" }`

For more details on the API go to [request.http](https://github.com/Lufedi/ECI-aygo-lab1/blob/master/app-lb-round-robin/request.http)


##### The database
Mongo was used to save the texts on a single collection with two fields, `value` and `creation_date`

##### Deployment
The application was deployed in a EC2 Instance (manually)

#### Demo 


![](https://media.giphy.com/media/tu11mmz4LEnxemcLpA/giphy.gif)

This is a picture of the 3 backends services receiving the requests with an even distribution

![](https://drive.google.com/uc?export=view&id=1KxZuWdQn0OyVYSh6efX2yI6MTECMKaqC)

Containers you should see after running `docker-compose`

![](https://drive.google.com/uc?export=view&id=1XXy9GyfkUIYGqjE1gbzRTk9aX2rkBKk3)
###

### TODO
- [x] Base project running in containers
- [ ] Implement health check to discard offline services
- [ ] Document code
- [ ] Validate input in the webapp form
- [ ] Use Nginx