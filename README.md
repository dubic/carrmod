# Carrmod

## Overview
Carrmod is a mobile first app that lets car owners stay on top of maintenances and document renewals. Users can get personalized reminders of maintenances when due and documents about to expire. 

## Infrastructure
![](/assets/carrmod.drawio.png)

Carrmod will be a mobile app with a backend infrastructure built on Google cloup platform services. Two microservices deployed on cloud run will handle all backend functionalities. 
<p> The backend API service will handle requests from the mobile client. Authentication, cars, documents, will be handled here. while the notification service will handle email notifications, reminders notifications, maintenance checks, etc.
<p> Both services will use a document (NoSQL) DB since little or no transactional features will be present.
Event driven communication will be handled by pubsub as this allows messages between services to be replayed when one service is down and recovers.
<p> Cloud scheduler will be configured to call notification service daily for reminders.
<p> Notification service will send notifications through email and firebase push messaging to users

## Models
![](/assets/carrmod-models.drawio.png)