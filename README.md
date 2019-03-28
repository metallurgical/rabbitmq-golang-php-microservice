# Introduction
Make a communication between Golang with PHP by using RabbitMQ as message broker that hosted in CloudAMQP. These days, microservice pattern or concept(you name it) have being spread and used widely because of its advantages over a monolith application. With microservices, we can decoupling the new/exisiting features from main codebase and isolated into its own language(can develop using any programming language besides from our main codebase) or environment without cluttering core codebase. 

## Terminology
These microservices eventually need to communicate to each other to receive and sending the data from our main application which exposed by HTTP from the outside world. Outside world here can be assumed as a Client(Mobile apps, web application, embedded) make a HTTP call to the backend server. Backend application will handle the rest of operations by distributed the process into multiple form.

Data would be passed to RabbitMQ(message broker) by a producer and the rest of microservices(consumers) will execute the task that exist inside a queue of RabbitMQ. As example, client want to send sms and at the same time send an email to specified member. 

                                                  sms process  
                                                     /        \  
                                                    /           \
[user] --> inititated process to send sms and email                --> pass to Message Broker 
                                                    \            /              |
                                                      \         /               |
                                                  email process           microservices
                                                                                |
                                                                                /\
                                                                               /  \
                                                                              /    \
                                                                        Email(PHP) Sms(Golang)
                                                                        
                                                                        
For this section, i assume we got two different 
