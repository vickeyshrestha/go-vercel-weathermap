# go-vercel-weathermap
A service that talks to Open Weather API to gather the information based on the provided coordinates. 
This demonstrates the Serverless Functions made with Go utilizing Vercel.

The latitude & longitude are passed as the URL parameters.

<h3>prerequisite </h3>
The following environment variables are required for the app to operate:


| Env Variable |                                                                                           Short Description                                                                                           | 
|-------------------------------------------------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:| 
| api_key |                                                       You need to obtain this yourself. Mine was something like ea***************************8                                                       | 
| base_url | Just set it as https://api.openweathermap.org/data/2.5/weather <br/>usually, this can stay constant as we are using ver 2.5. I added this as env var so that we are prepared for an upcoming version. | 


## Why Vercel?
Since AWS Lambda is cool but isn't always free, Vercel is a cool place to host some serverless functions. This is a simple Golang application utilizing Gin framework configured to run on a Vercel server.

# Running this app 

## Using App locally
This app utilizes Golang and HTML, but is not designed in a regular Go fashion. Hence, there is not main function here. Instead, this is wrapped around the Vercel, a platform primarily designed for front-end development and hosting.
We need a special way to fire up this application if we need to test it locally:
1. Install Vercel client: `npm install -g vercel`
2. From root of the project, run ` vercel dev`
3. If first time, this should download and setup Go ENV withing Vercel box.
4. The server should serve on port `localhost:3001` by default

You can then obtain the global coordinates using Google map. Once obtained, the lat and lon are passed as the URI parameters.

For example,

http://localhost:3001/weather?lat=32.7942144&lon=-97.1702272

For open weather api, more can be found at:

https://openweathermap.org/current