

GAPE  
  
Simple recursive filesystem notifier that writes to local or remote syslog servers.  
Compiled version can be found in the bin directory. It was compile on Debian Jesse,but will  
probably work on any other distribution.  
It currently logs any notifications that are known to work  
on any operating system like:  

 - Remove   
 - Create   
 - Write   
 - Rename  


Configuration:  

 - syslogproto: => udp or tcp  
 - sysloghost: Syslog server to send log messages to  
 - syslogport: Syslog server port  
 - localonly: Log only to local syslog (no network) => true or false  
 - stdout: Also output to stout => true or false  
 - paths: An array of paths to watch, these should be directories  

Usage:

 Download or clone this repo and in the bin directory you can find an example  
 config file, adjust this and copy it to /etc or somewhere else in which case  
 you start it with ./gape -config <path o your config>  
 If you take the default /etc directory simply do ./gape or whatever init or  
 systemd script you want to use.  There are several tools for daemonizing, at a later  
 stage I might build this in (if time permits :)  

Remember when an underlying watch has been removed, a restart of Gape is needed.  
Based on the excellent library from https://github.com/rjeczalik/notify  
If you found any bugs or are using it to full hapiness drop me an email.  
