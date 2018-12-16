ws2js - Websockets to Javascript
========================================

Send javascript commands from stdin to a live webpage through
websockets.


Use Case
------------------------------

Original use case is to provide an easy way to create live
graphics/animations while using programming languages that don't have
convenient graphics packages.

An example would be drawing things in the HTML canvas from a python
script while it's running.

Could be used in bash by piping output from one program to ws2js

~~~bash
myprogram | ws2js
~~~


How it Works
------------------------------

1. ws2js serves a webpage to `http://localhost:8080`
2. ws2js serves a websocket to `ws://localhost:8080/ws`
3. reads user input (javascript) from stdin until EOF.
4. sends to the websocket.
5. web browser recieves text from the websocket.
6. In a sandboxed iframe, browser runs the javascript.

~~~
user input --> stdin --> websockets --> webpage
~~~
