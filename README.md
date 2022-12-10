# JSvarXSSTester
A golang tool entirely written with ChatGPT that takes URLS on stdin, extracts javascript variable names from the webpages, then uses them as query parameters 5 at a time checking the response bodies for reflections. 

There is a lot that could be improved and implemented but the point was to get ChatGPT to do all the work and frankly, it took probably longer than it would have if I just wrote it myself. The main point is that it works and is capable of finding reflections that lead to XSS. 

I might refine it further, still using ChatGPT, I also might not. just a bit of fun :) 
