package AsyncLogging

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Basic logger struct. Use message channel for async logging and writer for synchronous logging
type TLog struct {
	dest             io.Writer
	m                *sync.Mutex
	msgCh            chan string
	errCh            chan error
	shutdownCh       chan struct{}
	shutdownComplete chan struct{}
}
//creates new log object that writes to given iowriter or if not then to standard out
func NewLogObject(w io.Writer) *TLog{
	if w==nil{
		w=os.Stdout
	}
	return &TLog{
		dest: w,
		msgCh:make(chan string),
		errCh:make(chan error),
		m: &sync.Mutex{},
	}
}

func (logobj TLog) Start(){
   for {
	   msg:=<-logobj.errCh
	   go logobj.write(msg,nil)
   }
}

func (logobj TLog) formatMessage(msg string) string{
	if !strings.HasSuffix(msg,"\n"){
		msg += "\n"
	}
	return fmt.Sprintf("[%v]-%v",time.Now().Format("2022-11-25 17:07:07"),msg)
}

func (logobj TLog) write(msg string, wg *sync.WaitGroup){
	logobj.m.Lock()
	defer logobj.m.Unlock()
    _,err:=logobj.dest.Write([]byte(logobj.formatMessage(msg)))
	if err!=nil{
		go func(err error){
			logobj.errCh<-err
		} (err)
	}
}















