# Skycoin Wallet Bot

The idea is to create and inline telegram bot that can be used within telegram group conversations, with hooks into telegram cli to send tips to other users by simply typing `/sendsky 100 @V1nnie` initially this bot will be limited to small tips 0.1SKY or 1000 drops. 

At the begining this will be similar to how a centralized wallet works without needing users to enter private keys, but will progress to being controlled entirely by end users. 

with functionalty that is possible from the SKYCOIN-CLI such as 

Generate wallet
Check balances 
Send to Multiple address
etc

## Install and Build ## 
```
mkdir -p $GOPATH/src/github.com/Cryptovinnie
cd $GOPATH/src/github.com/Cryptovinnie
git clone https://github.com/Cryptovinnie/SkycoinWalletBot.git
```

## Install PostGreSQL ##

```
$ sudo apt-get update
$ sudo apt-get install postgresql postgresql-contrib
```

Switch over to the postgres account on your server by typing:

`$ sudo -i -u postgres`

You can now access a Postgres prompt immediately by typing:

`$ psql`

You will be logged in and able to interact with the database management system right away.

Exit out of the PostgreSQL prompt by typing:

`$ \q`

## Enabling Postgres command line tools ##
If you are using the default terminal, you are going to want to modify the file at ~/.bash_profile. If you are using something like Oh My Zsh you are going to want to modify the file ~/.zshrc.

To edit this file you likely need to open it via the terminal, so open your terminal and type `nano ~/.bash_profile` 

Once your zbash_profile or .zshrc file is open, add the following line to the end of the file:

`export PATH=$PATH:/Applications/Postgres.app/Contents/Versions/latest/bin`

After that you will need to quit and restart your terminal This is to make sure it reloads with the changes you just made.

Once you have restarted your terminal, try running psql.

`psql -U postgres`
You should get the following output.

`psql (9.6.0)
Type "help" for help.`

## Creating a Postgres database ##
The first thing we need to do is connect to Postgres with our postgres role. To do this we want to type the following into our terminal.

`psql -U postgres`

Creating Database 
We will create the DB which will include all fields that are needed. fields will be 

`| telegram_username | public_wallet | public_key | private_key |`

`CREATE DATABASE skycoinbalancesDB;`

Next we want to connect to our database. We do that by typing the following.

`\c skycoinbalancesDB`

```
CREATE TABLE users (
  id SERIAL,
  telegram_username TEXT UNIQUE NOT NULL,
  public_wallet TEXT UNIQUE NOT NULL,
  public_key TEXT UNIQUE NOT NULL,
  private_key TEXT UNIQUE NOT NULL 
);
```

Next enter a test entry into DB 
```
psql -U postgres -d skycoinbalancesDB
INSERT INTO users (id, telegram_username, public_wallet, public_key, private_key)  
VALUES (0, '@testing', 'pubwallet123', 'pubkey123', 'privkey123');
```  

You should see the output `INSERT 0 1` after inserting this row.

If you would like to see the data you just inserted into your table, as well as the auto-incrementing id, you can do so by running the following SQL.  

```SELECT * FROM users;

 id | telegram_username | public_wallet | public_key |    private_key   
----+-------------------+---------------+------------+----------------  
  1 |    @testing       | pubwallet123  |  pubkey123 |    privkey123  
  
 ```  
 
 ## Configuration ## 
 
 In Telegrambot.go under `func main()` please enter telegram bot token.

 ```golang
 func main() {
        //Telegram messenger
        bot, err := tgbotapi.NewBotAPI("TELEGRAM-BOT-TOKEN-HERE")
        if err != nil {
                log.Panic(err)
        }
 ```
 
 ## Different Commands ## 
 If Bot receives any messages starting with the bellow they will be linked to a command. 
 
 ```golang
 Message := update.Message.Text //Message received by bot
 p("Message from telegram: ", Message) //Print message received
```
### /wallet ###

`Wallet := s.HasPrefix(Message, "/wallet:") //If Message starts with wallet:`  

If message to bot contains "/wallet:"  

Get Telegram username and unique chatID!!! 

Split String after Semi-colon (:)  
eg. "/wallet:1234567890skyaddress" 
split too "1234567890skyaddress" 

Connect to Skycoin-CLI and get wallet balance. return this in a message via TelegramBot.

```
address:1234567890skyaddress
balance: 1000 SKY 
coinhours: 100K
```




### /createaddress ###
`createaddress := s.HasPrefix(Message, "/createaddress") //If Message starts with createaddress`  

If message to bot contains "/createaddress"
  
Get Telegram username and unique chatID!!!  
 
Search SQL Database to see if Username already exists. 
  
If Username is in DB send back message with users public address 

If userdoes not exist, Connect to Skycoin-CLI, GenerateKeypair, then Save to SQL DB 


### /getaddress ###
`getaddress := s.HasPrefix(Message, "/getaddress")//If Message starts with getaddress`  

If message to bot contains "/getaddress"
Do same as createaddress. 

### /sendsky ### 
`sendsky := s.HasPrefix(Message, "/sendsky")//If Message starts with sendsky`  

If message to bot contains "/sendsky"  
eg `/sendsky 100 @V1nnie` message sent from @Synth

Get Telegram username and unique chatID!!!  

```
- telegram_username = @Synth
- chatId = 1111111
```

Message in format "/sendsky 100 @V1nnie"  

Check SQL Database for sender Telegram username. -@Synth

```
id | telegram_username | public_wallet | public_key |    private_key   
----+------------------+---------------+------------+----------------  
  1 |    @Synth   ✅   | pubwallet123  | pubkey123 |    privkey123  
  2 |    @V1nnie   ❗❗  | V1nnieaddress |  V1nniepub |    V1nniepriv
 ```  
 
Connect to Skycoin-CLI and check spendable balance, Save this as a variable.

var address = pubwallet123
var SPENDABLE = 1000

Split String up get @Username address and int amount eg `100`, -@V1nnie❗
Check this address in SQL database for wallet address, If no address exists do `/createaddress` 


Once @Username exists in SQL Database, Send int amount to @Username Public address  
`100 SKY from @synth (pubwallet123) to @V1nnie (V1nnieaddress)`

Connect to Skycoin-CLI "sendto" (entercli command arg here) send from address @synth (publickey123) to @V1nnie

Then send confirmation to Telegram user who initiated transaction + if chatID exists in SQL Database send Verification message. 














