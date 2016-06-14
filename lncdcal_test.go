package lncdcal

import (
  "testing"
  "lncdcal"
  "os"
  "fmt"
  "io/ioutil"
)

var realini string = "gcal.ini";
var testini string = "gsettings_test.ini";

func writeTestfile(key string, email string) error{
  txt := []byte(fmt.Sprintf("pemFile=%s\nserviceEmail=%s",key,email))
  err := ioutil.WriteFile(testini,txt,0644)
  return err
}


// we error if there is no file
func TestGsettings_nofile(t *testing.T){
  if _,err:= os.Stat(testini); ! os.IsNotExist(err) {
   t.Skipf("%s exists! this should happen!", testini)
  }

  email,pem,err := lncdcal.Gsettings(testini)
  if err == nil {
    t.Error("there is no file to read, how did we not error?",email,pem )
  }
}

// read settings correctly
func TestGsettings_parsefile(t *testing.T){
  if _,err:= os.Stat(testini); ! os.IsNotExist(err) {
   t.Skipf("%s exists! this should happen!", testini)
  }

  var expectpem   string = "key.pem"
  var expectemail string = "fake@gmail.com"

  if err := writeTestfile(expectpem,expectemail); err != nil {
   t.Skipf("could not write %s: %v", testini, err)
  }
  defer os.Remove(testini)

  email,pem,err := lncdcal.Gsettings(testini)
  if err   != nil ||
     email != expectemail ||
     pem   != expectpem {
    t.Errorf(
      "did not read %s correctly, got vs expect: '%s' vs '%s', '%s' vs '%s'!",
      testini,
      email,expectemail,
      pem,expectpem )
  }

}



// authenticate with a real file
func TestLogin(t *testing.T){
  if _,err:= os.Stat(realini); os.IsNotExist(err) {
   t.Skipf("do not have %s to test gsettings",realini)
  }

  _,err := lncdcal.Login(realini)
  if err != nil {
     t.Error("failed to log in!")
  }
}

