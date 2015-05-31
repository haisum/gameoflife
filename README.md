Game of life simulator
--------------------------

Source Documentation: http://godoc.org/github.com/haisum/gameoflife

This project is Conway's game of life's simulation. It's written in golang and is capable of displaying simulation for large life spaces limited only by display space, memory and CPU power.


Installation
--------------------

Recommended way is to download following binaries which I compiled for all popular platforms:

- Linux: https://drive.google.com/file/d/0B9oMRwzY0tXsdk1TV3VQZVVRTEk/view?usp=sharing
- Mac: https://drive.google.com/file/d/0B9oMRwzY0tXsdFE2UUNsbkVid3c/view?usp=sharing
- Windows: https://drive.google.com/file/d/0B9oMRwzY0tXsSWpUVEZfMlA1NUE/view?usp=sharing

####Manual build

- Install golang and git. 
- Create a new directory `mysimulation`
- Set environment variable `GOPATH="/path/to/mysimulation"` (Linux: `export GOPATH="/path/to/mysimulation"`).
- Open terminal, cd to  `mysimulation` and run `go get github.com/haisum/gameoflife`. 
- `cd` to `mysimulation/src/github.com/haisum/gameoflife`. 
- Compile simulator by executing `go build _simulator/simulator.go`. 
- You should have a binary named "simulator" or "simulator.exe" in your current directory.

####Vagrant

Vagrant file isn't yet available for this project, because of lesser time I had. You can still use this awesome vagrant setup provided by @dcoxall at https://github.com/dcoxall/vagrant-golang

Usage
-------------------

This project supports two display modes: terminal and http. Terminal mode displays simulation in command line and Http mode starts a Http server so you can see simulation on a browser.

Options
------------

- **-a** List of alive cells. Format: x1:y1,x2:y2,x3:y3,....,xn:yn.
- **-y** Number of columns in life space 
- **-x** Number of rows in life space
- **-t** Only applicable when -d is set to terminal. If passed, text only output is shown without any colors. Useful for systems where ansi coloring is not supported and program outputs garbage text.
- **-r** Refresh rate for animation in milliseconds.
- **-d** Display interface for simulation. Two values are supported right now: "terminal" and "http"
- **-p** Port number to listen on for http requests. Only applicable when -d is set to http.

Screenshots
---------------

####Output on a linux terminal

![Usage with basic options](http://i.imgur.com/jcRlOAp.png)

####Using http display option (-d http)

![-d http](http://i.imgur.com/baP9vxG.png)

####Browser output
![Browser](http://i.imgur.com/mVCtLcb.png)

####Text only option (-t)
![Text only](http://i.imgur.com/viSwqlU.png)

####Output on Mac Terminal
![Mac terminal](http://i.imgur.com/PaB1Fh7.jpg)

####Simulation on windows
![Imgur](http://i.imgur.com/czLPcLz.jpg)

Note: Due to lack of proper text output support on windows command line, there are sometimes distorted charachters. I recommend using `-d http` option when on windows.

![Imgur](http://i.imgur.com/uZ7b9dT.jpg)

Optimized and Can be Optimized
----------------------

This simulation uses lesser memory than most of other projects I found. It does so by recording only live cells on a life space. So rest of grid doesn't occupy memory.

Algorithm for computing next generation of life has a lot of room for improvement, in fact in its current state it has worst case O(N)squre complexity which could be reduced by utlizing Quad Trees and HashLife algorithm. May be we'll do it at some other time.