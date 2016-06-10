package lncdcal

import (
  "testing"
  "lncdcal"
  "os"
  "fmt"
)

var realini string = "gcal.ini";
var testini string = "gsettings_test.ini";

func writeTestfile(key string, email string) error{
  txt := []byte(fmt.Spritnf("pemFile=%s\nserviceEmail=%s",key,email))
  err := ioutil.WriteFile(testini,txt,0644)
}


// we error if there is no file
func Test_gsettings_nofile(t *testing.T){
  if _,err:= os.Stat(testini); ! os.IsNotExist(err) {
   t.Skipf("%s exists! this should happen!", testini)
  }

  email,pem,err = gsettings(testini)
  if err == nil {
    t.Error("there is no file to read, how did we not error?",email,pem )
  }
}

// read settings correctly
func Test_gsettings_nofile(t *testing.T){
  if _,err:= os.Stat(testini); ! os.IsNotExist(err) {
   t.Skipf("%s exists! this should happen!", testini)
  }

  writeTestfile("key.pem","fake@email.com")
  defer os.Remove(testini)

  email,pem,err = gsettings(testini)
  if err   == nil &&
     email == "fake@email.com" &&
     pem   == "key.pem" {
    t.Errorf("did not read %s correctly, got e: %s, k: %s!",testini,email,pem )
  }

}



// authenticate with a real file
func Test_login(t *testing.T){
  if _,err:= os.Stat(realini); os.IsNotExist(err) {
   t.Skipf("do not have %s to test gsettings",realini)
  }

  cal,err := login(realini)
  if err != nil {
     t.Error("failed to log in!")
  }
}

