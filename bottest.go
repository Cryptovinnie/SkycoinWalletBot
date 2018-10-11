package main

import s "strings"
import (
  	//"database/sql"
   	"fmt"
   	"log"
   	"gopkg.in/telegram-bot-api.v4"
      _ "github.com/lib/pq"
	viper "github.com/spf13/viper"
	"github.com/Cryptovinnie/SkycoinWalletBot/config"
)
var p = fmt.Println


func main() {
	//get telegrambotapikey from config file. 
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration
	
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	var telegramapikeys = configuration.Telegram.Apikey
	
	log.Printf("TelegramAPI is %s", telegramapikeys)
	
	//Database settings 
	log.Printf("host is %s", configuration.SqlDatabase.Host)
	log.Printf("port is %d", configuration.SqlDatabase.Port)
	log.Printf("user is %s", configuration.SqlDatabase.User)
	log.Printf("password is %s", configuration.SqlDatabase.Password)
	log.Printf("dbname is %s", configuration.SqlDatabase.Dbname)
	
	
	var host     = configuration.SqlDatabase.Host
	var port     = configuration.SqlDatabase.Port
	var user     = configuration.SqlDatabase.User
	var password = configuration.SqlDatabase.Password
	var dbname   = configuration.SqlDatabase.Dbname
	
	
        //Telegram messenger
        bot, err := tgbotapi.NewBotAPI(telegramapikeys)
        if err != nil {
                log.Panic(err)
		}
		
 		bot.Debug = true

        log.Printf("Authorized on account %s", bot.Self.UserName)

        u := tgbotapi.NewUpdate(0)
        u.Timeout = 60

        updates, err := bot.GetUpdatesChan(u)

        for update := range updates {
                if update.Message == nil {
                        continue
                }

                log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
                Message := update.Message.Text //Text received by bot
                p("Message from telegram: ", Message) //Print text received
                //Message := "wallet:1234567890" //testing string

              Wallet := s.HasPrefix(Message, "/wallet:") //If Message starts with wallet
              createaddress := s.HasPrefix(Message, "/createaddress") //If Message starts with createaddress
              getaddress := s.HasPrefix(Message, "/getaddress")//If Message starts with createaddress
              sendsky := s.HasPrefix(Message, "/sendsky")//If Message starts with createaddress


		switch Wallet {
		case true: //If Message starts with wallet: then do this
		p("Message Switch worked: ", s.HasPrefix(Message, "wallet:"))
		UserName := update.Message.From.UserName
		p("Username is ", UserName)
		split := s.SplitAfter(Message, ":")
	
		p("Split message ", split)
		p("Split message1 ", split[0])
		p("Split message2 ", split[1])
		p("Len: ", len(Message))
		p()
	
		address := split[1]
		if address != ""{
			response := fmt.Sprintf("https://explorer.skycoin.net/app/address/%s", address)
			p("response ", response)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
			}

			default: // If none of the above options do this.
				p("Message Switch not worked: ", s.HasPrefix(Message, "wallet"))
		}
	}
}
