# SkycoinWalletBot

Install PostGreSQL

```$ sudo apt-get update`
$ sudo apt-get install postgresql postgresql-contrib```

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

```CREATE TABLE users (
  id SERIAL,
  telegram_username TEXT UNIQUE NOT NULL,
  public_wallet TEXT UNIQUE NOT NULL,
  public_key TEXT UNIQUE NOT NULL,
  private_key TEXT UNIQUE NOT NULL 
);```

Next enter a test entry into DB 
```psql -U postgres -d skycoinbalancesDB
INSERT INTO users (id, telegram_username, public_wallet, public_key, private_key)  
VALUES (0, '@testing', 'pubwallet123', 'pubkey123', 'privkey123');```  

You should see the output `INSERT 0 1` after inserting this row.

If you would like to see the data you just inserted into your table, as well as the auto-incrementing id, you can do so by running the following SQL.  

`SELECT * FROM users;`

``` id | telegram_username | public_wallet | public_key |    private_key   
----+-------------------+---------------+------------+----------------  
  1 |    @testing       | pubwallet123  |  pubkey123 |    privkey123  ```  








