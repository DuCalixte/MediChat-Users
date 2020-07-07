package notifier

import (
  // "fmt"
	// "log"
  "encoding/json"

	// "github.com/anaskhan96/go-vapidkeys"
  webpush "github.com/SherClockHolmes/webpush-go"

  "github.com/DuCalixte/MediChat-Users/settings"
)

const (
  subscription    = ``
  // Subscriber = settings.NotificationSetting.Subscriber
  // PublicKey = settings.NotificationSetting.PublicKey
  // PrivateKey = settings.NotificationSetting.PrivateKey
)

// var PrivateKey string
// var PrivateKey string

// func GenerateKeys() {
//   var err error
//   PrivateKey, PublicKey, err = vapidkeys.Generate()
// 	if err != nil {
// 		// log.Fatal(err)
//     log.Printf("Unable to generate keys %v", err)
//     return
// 	}
// }

func InitNotifier() {
  // Decode subscription
	notifier := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), notifier)

  var subscriber = settings.NotificationSetting.Subscriber
  var publicKey = settings.NotificationSetting.PublicKey
  var privateKey = settings.NotificationSetting.PrivateKey

	// Send Notification
	response, err := webpush.SendNotification([]byte("Test"), notifier, &webpush.Options{
		Subscriber:      subscriber,
		VAPIDPublicKey:  publicKey,
		VAPIDPrivateKey: privateKey,
		TTL:             30,
	})
	if err != nil {
		// TODO: Handle error
	}
	defer response.Body.Close()

}
