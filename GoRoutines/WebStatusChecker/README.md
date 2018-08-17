### Web Status Checker

HTTP ping daemon to check the status of websites created using Go routines.
It uses channels to signal between go routines to synchronize the time between
two consecutive pings to a website.

Can be used to check multiple websites at once.
