package lncdcal

import (

  "io/ioutil"
  "log"
  "time"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/jwt"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/calendar/v3"
  "gopkg.in/ini.v1"
  "fmt"
 
  //"reflect"
  //"flag"
)

// read settings from a specified ini file
//   serviceEmail=xxx@yyy.iam.gserviceaccount.com
//   pemFile=key.pem
func Gsettings(authini string) (string,string,error){
  cfg, err := ini.Load(authini) //e.g. "gcal.ini"
  if err!=nil {
    //log.Fatalf("cannot open %s!", authini) //TODO add authin to output
    return "","",fmt.Errorf("cannot open %s",authini)
  }
  serviceEmail,eerr := cfg.Section("").GetKey("serviceEmail")
  keyPath,kerr := cfg.Section("").GetKey("pemFile")
  if eerr !=nil || kerr != nil { 
    return "","",fmt.Errorf("error reading %s: %v %v",authini,eerr,kerr)
  }

  return serviceEmail.String(), keyPath.String(), nil
}



// Your credentials should be obtained from the Google
// Developer Console (https://console.developers.google.com).
//Email: "xxx@developer.gserviceaccount.com",
//  OR   xxxx@yyyy.iam.gserviceaccount.com
//PrimaveKey:
// The contents of your RSA private key or your PEM file
// that contains a private key.
// If you have a p12 file instead, you
// can use `openssl` to export the private key into a pem file.
//
//    $ openssl pkcs12 -in key.p12 -passin pass:notasecret -out key.pem -nodes
//
// The field only supports PEM containers with no passphrase.
// The openssl command will convert p12 keys to passphrase-less PEM containers.
// If you would like to impersonate a user, you can
// create a transport with a subject. The following GET
// request will be made on the behalf of user@example.com.
// Optional.
//Subject: "user@example.com",

func Login(authini string) (*calendar.Service, error) {


  serviceEmail,keyPath,err := Gsettings(authini)
  if err!=nil {
    log.Fatal(err)
  }

  keyBytes, err := ioutil.ReadFile(keyPath)
  if err != nil {
   log.Fatal("cannot open pemFile")
  }

  conf := &jwt.Config{
      Email: serviceEmail,
      PrivateKey: keyBytes,
      Scopes: []string{
          "https://www.googleapis.com/auth/calendar",
      },
      TokenURL: google.JWTTokenURL,
  }
  client := conf.Client(oauth2.NoContext)
  cal,err := calendar.New(client)

  return cal,err
}

// put time into a fomate that calendar understands
func TimeToCal( t time.Time) *calendar.EventDateTime {
 return &calendar.EventDateTime{DateTime: t.Format(time.RFC3339)}
}


// add an event to the primary
func addEvent(calServ *calendar.Service,
             summary string ,
             start  time.Time,
             end    time.Time) {

 event := calendar.Event{
   Summary: summary,
   Start: TimeToCal(start),
   End: TimeToCal(end),
 }

 calServ.Events.Insert("primary",&event)
}


// func main() {
//   cal,err := login()
//   if err != nil {
//      log.Fatal(err)
//   }
// 
//   // https://jacobmartins.com/2016/03/08/practical-golang-using-google-drive-and-calendar/
//   events,err := cal.Events.List("primary").TimeMin(time.Now().Format(time.RFC3339)).MaxResults(5).Do()
// 
//   log.Println(events)
// 
// 
// }


