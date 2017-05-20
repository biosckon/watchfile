### Watch a file and run a command when it is done
    
```Shell
$ watchfile index.html diff index.html index_old.html
```
    
This command example will watch file index.html for changes. Once detected run a command `diff index.html index_old.html`. A command to run can be anything.

The watcher continues to run untill you Ctrl-C or kill it.

To make it ran at the backgroud all the time example use:
```Shell
$ nohup watchfile index.html diff index.html index._old.html &
```
this will make it immunune to bash terminal closure (`nohup`) and shall continue to run the it at the background (`&`). 

### installation
Download the required binary here: https://github.com/biosckon/watchfile/releases/

Alternatively for go programmers... you know what to do.
