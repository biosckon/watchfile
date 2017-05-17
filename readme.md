### Watch a file and run a command when it is done
    
```Shell
$ watchfile index.html diff index.html index_old.html
```
    This command example will watch file index.html for changes. Once detected run a command ```diff index.html index_old.html```. A command to run can be anything.
    