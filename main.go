package  main

import (
	"log"	
	viper "github.com/spf13/viper"
	"github.com/Cryptovinnie/SkycoinWalletBot/config"

)

func main() {
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



	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)
	log.Printf("TelegramAPI is %s", configuration.Telegram.Apikey)
	log.Printf("SQL1 host is %s", configuration.Sqldatabase.Host)
	log.Printf("SQL2 port is %d", configuration.Sqldatabase.Port)
	log.Printf("SQL3 user is %s", configuration.Sqldatabase.User)
	log.Printf("SQL4 password is %s", configuration.Sqldatabase.Password)
	log.Printf("SQL5 dbname is %s", configuration.Sqldatabase.Dbname)


}
