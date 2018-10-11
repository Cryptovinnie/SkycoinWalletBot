package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	s "strings"

	_ "github.com/lib/pq"
	"gopkg.in/telegram-bot-api.v4"
)

var p = fmt.Println

const (
	host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "masterpassword"
		dbname   = "skycoinbotdatabase"
)

func input(x string) string {
	path := "/home/username/go/bin/skycoin-cli"
	input := x //enter argument here to run
	cmd := exec.Command(path, input)

	var out bytes.Buffer
	multi := io.MultiWriter(os.Stdout, &out)
	cmd.Stdout = multi

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf(out.String())
	return out.String()

}

func connectDB() string {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "masterpassword"
		dbname   = "skycoinbotdatabase"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	success := "Successfully Connected to DB"
	fmt.Println(success)

	return success
}

func telegram() {

	bot, err := tgbotapi.NewBotAPI("APIKEYS")
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
		UserName := update.Message.From.UserName
		Message := update.Message.Text //Text received by bot
		chatid := update.Message.From.ID
		Wallet := s.HasPrefix(Message, "/wallet:")              //If Message starts with wallet
		createaddress := s.HasPrefix(Message, "/createaddress") //If Message starts with createaddress
		getaddress := s.HasPrefix(Message, "/getaddress")       //If Message starts with getaddress
		sendsky := s.HasPrefix(Message, "/sendsky")             //If Message starts with sendsky
		p("Message from telegram: ", Message)                   //Message
		p("Chatid from telegram: ", chatid)                     //ChatID
		p("Wallet from telegram: ", Wallet)                     //Wallet
		p("createaddress from telegram: ", createaddress)       //CreateAddress
		p("getaddress from telegram: ", getaddress)             //Getaddress
		p("sendsky from telegram: ", sendsky)                   //SendSky

		switch Wallet {
		case true:
			split := s.SplitAfter(Message, ":")
			address := split[1]
			if address != "" {
				response := fmt.Sprintf("https://explorer.skycoin.net/app/address/%s", address)
				p("response ", response)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				bot.Send(msg)
			}
		}
		switch createaddress {
		case true:

			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
			db, err := sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			err = db.Ping()
			if err != nil {
				panic(err)
			}
			sqlStatement := `SELECT id, public_wallet FROM users WHERE telegram_username=$1;`
			var public_wallet string
			var telegram_username string

			row := db.QueryRow(sqlStatement, UserName)
			switch err := row.Scan(&telegram_username, &public_wallet); err {

			case sql.ErrNoRows:
				fmt.Println("No rows were returned!") //User was not found in DB so create address
				Input := input("generateAddresses")   //String to enter after skycoin-cli
				AddrCreated := Input                  //Save created Address to AddrCreated
				//Then Save created wallet to SQL DB
				sqlStatement := `
										INSERT INTO users (chatid, telegram_username, public_wallet, public_address, private_key)
										VALUES ($1, $2, $3, $4, $5)
										RETURNING id`
				id := 0
				err = db.QueryRow(sqlStatement, chatid, UserName, AddrCreated, "Publickey12213", "Privatekey12123").Scan(&id)
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
		}
	}

}

func main() {
	connectDB()
	telegram()

	x := "listWallets"
	Input := input(x)
	fmt.Println("Input in Main()", Input)

}
