package main

import s "strings"
import (
  	"database/sql"
   	"fmt"
   	"log"
   	"gopkg.in/telegram-bot-api.v4"
      _ "github.com/lib/pq"
)
var p = fmt.Println

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "masterpassword"
  dbname   = "skycoinbalancesDB"
)

func main() {
        //Telegram messenger
        bot, err := tgbotapi.NewBotAPI("TELEGRAMBOTAPIKEYS")
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

	switch createaddress {
        case true: //If message starts with createaddress:
                       p("createaddress Switch worked: ", s.HasPrefix(Message, "createaddress"))
                       UserName := update.Message.From.UserName
                       p("Username is ", UserName)

                       //Check database for userid existing.
                        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
                        "password=%s dbname=%s sslmode=disable",
                        host, port, user, password, dbname)
                        db, err := sql.Open("postgres", psqlInfo)
				if err != nil { 
					panic(err) 
					}
			defer db.Close()

                        sqlStatement := `SELECT id, public_wallet FROM users WHERE telegram_username=$1;`
                                     var public_wallet string
                                     var telegram_username string

                                     row := db.QueryRow(sqlStatement, UserName)
                                     switch err := row.Scan(&telegram_username, &public_wallet); err {

                                     case sql.ErrNoRows:
                                     fmt.Println("No rows were returned!") //User was not found in DB so create address

                                     // Run address creation script here
                                     //Connect to Skycoin-CLI and run address_gen2.go
                                     //Save output to SQL database.
                                     AddrCreated := "abcdefghijklmnop"// testing AddrCreated
																		
				//sqlinfo bellow
				psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
						"password=%s dbname=%s sslmode=disable",
						host, port, user, password, dbname)
					db, err := sql.Open("postgres", psqlInfo)
							if err != nil {
							panic(err)
							}
							defer db.Close()


								// Save created wallet to SQL DB
								sqlStatement := `
								INSERT INTO users (telegram_username, public_wallet, public_key, private_key)
								VALUES ($1, $2, $3, $4)
								RETURNING id`
									id := 0
								err = db.QueryRow(sqlStatement, UserName, AddrCreated, "Publickey1", "Privatekey1").Scan(&id)
								if err != nil {
								panic(err)
								}
								fmt.Println("New record ID is:", id)
								//End of SQL info. Still in For Updates loop under wallet: switch.


																		//send message back
																		msg := tgbotapi.NewMessage(update.Message.Chat.ID, AddrCreated)
																		bot.Send(msg)

																		}
																		//Address alread created and exists.
																		msg := tgbotapi.NewMessage(update.Message.Chat.ID, public_wallet)
																		bot.Send(msg)
																	default:
																		p("createaddress Switch not worked: ", s.HasPrefix(Message, "createaddress"))
														}
								
				switch getaddress {
						case true:
														p("getaddress Switch worked: ", s.HasPrefix(Message, "createaddress"))
														UserName := update.Message.From.UserName
														p("Username is ", UserName)
														psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
																		"password=%s dbname=%s sslmode=disable",
																		host, port, user, password, dbname)
																db, err := sql.Open("postgres", psqlInfo)
																if err != nil {
																		panic(err)
																}
																defer db.Close()
								
								
	sqlStatement := `SELECT id, public_wallet FROM users WHERE telegram_username=$1;`
	var public_wallet string
	var id int
																var telegram_username string
																		row := db.QueryRow(sqlStatement, UserName)
																		switch err := row.Scan(&telegram_username, &public_wallet); err {
																		case sql.ErrNoRows:
																		fmt.Println("No rows were returned!")
																		case nil:
																		fmt.Println(id, public_wallet)

																		                              //send message back
																									  msg := tgbotapi.NewMessage(update.Message.Chat.ID, public_wallet)
																									  bot.Send(msg)
															  
																									  default:
																									  panic(err)
																									  }
															  
															  
															  
																									default:
																										p("getaddress Switch not worked: ", s.HasPrefix(Message, "getaddress"))
																										}
																				
																										switch sendsky {
																										case true:
																												p("sendsky Switch worked: ", s.HasPrefix(Message, "sendsky"))
																										default:
																										p("sendsky Switch not worked: ", s.HasPrefix(Message, "sendsky"))
																										}
																										if Message == "/help" {
																										start :=
																												`
																												/wallet:123Skyaddresshere
																												to check balance
																				
																												/createaddress
																												to link telegram to skyaddress
																				
																												/getaddress
																												get your public wallet address
																				
																												/sendsky
																												send sky to user
																												`
																				
																												msg := tgbotapi.NewMessage(update.Message.Chat.ID, start)
																												bot.Send(msg)
																										}
																						}
																				}
																				
																		
																				
       
        
