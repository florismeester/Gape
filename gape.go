package main


/* Easy to use recursive directory filesystem changes notification tool that can send output to 
   remote or local syslog servers, based on github.com/rjeczalik/notify
   Floris Meester floris@grid6.io 
*/


import (
    "log"
    "encoding/json"
    "github.com/rjeczalik/notify"
    "flag"
    "os"
    "fmt"
    "log/syslog"
)

type Configuration struct {
    Sysloghost string
    Syslogproto string
    Syslogport string
    Stdout bool
    Localonly bool
    Paths []string
}


func main(){

        // Open configuration file
        conf := flag.String("config","/etc/gape.conf", "Path to configuration file")
        flag.Parse()
        file, err := os.Open(*conf)
        if err != nil {
                log.Fatal("Can't find configuration file, try 'gape -config <path> ", *conf)
        }
        decoder := json.NewDecoder(file)
        configuration := Configuration{}
        err = decoder.Decode(&configuration)
        if err != nil {
                fmt.Println("error opening configuration:", err)
        }

	// Create the notification channel
	c := make(chan notify.EventInfo, 1)

        // Create a syslog writer for logging local or remote
	if configuration.Localonly{
	        logger, err := syslog.New(syslog.LOG_NOTICE, "Gape")
                if err == nil {
                	log.SetOutput(logger)
		} else {
			log.Fatal(err)
		}
        }else {
        	logger, err := syslog.Dial(configuration.Syslogproto, configuration.Sysloghost +
			":" + configuration.Syslogport, syslog.LOG_NOTICE, "Gape")
                	if err == nil {
                	log.SetOutput(logger)
        	}else{
			log.Fatal(err)
		}	
	}

	// Create the notification watches
	for _,item := range configuration.Paths {
		
		// Check if item exists and is a directory otherwise bailout
		fd, err := os.Stat(item)
		if err != nil {
			log.Fatal(err)
		}
		if !fd.IsDir(){
			log.Fatal("Not a directory: ", item)
		}
		// I might move the notification options to the config file
		if err := notify.Watch(item, c, notify.Remove, notify.Create, notify.Write, notify.Rename ); err != nil {
    			log.Fatal(err)
		}
	}
	
	defer notify.Stop(c)

	// Loop forever and receive  events from the channel.
	for {
		ei := <-c
		log.Print(ei)
		if configuration.Stdout {
			fmt.Println(ei)
		}
	}
}


