# go-vercel-weathermap
A service that talks to Open Weather API to gather the information based on the provided coordinates. 

The latitude & longitude are passed as the URL parameters.

<h3>prerequisite </h3>
The following environment variables are required for the app to operate:


| Env Variable |                                                                                           Short Description                                                                                           | 
|-------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:| 
| api_key |                                                       You need to obtain this yourself. Mine was something like ea***************************8                                                       | 
| base_url | Just set it as https://api.openweathermap.org/data/2.5/weather <br/>usually, this can stay constant as we are using ver 2.5. I added this as env var so that we are prepared for an upcoming version. | 


## Why Vercel?
Since AWS Lambda is cool but isn't always free, Vercel is a cool place to host some serverless functions. This is a simple Golang application utilizing Gin framework configured to run on a Vercel server.

## Using App 

You can then obtain the global coordinates using Google map. Once obtained, the lat and lon are passed as the URI parameters.

For example,

http://localhost:8080/weather?lat=32.7942144&lon=-97.1702272

For open weather api, more can be found at:

https://openweathermap.org/current