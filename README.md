# DiscordRAT
A RAT using discord as a command and control center written in go

#Building
You need to get this libary https://github.com/HaiHakkuIku/DiscordRAT
then run "go build -ldflags -H=windowsgui main.go" as that will hide the window 

#Instructions
This is for educational purposes only im not responsible for what you do with this!

Create a bot @ https://discordapp.com/developers/applications/me you will need an account
In your config add your bot token to the first line and that can be found under token
![alt tag](http://i.imgur.com/TUoPqZB.png)

Then on the second line add the id of the account that will control the bot this can be found by enabling developer option
![alt tag](http://i.imgur.com/gCkjaoQ.png)

then by right clicking your name and clicking "Copy ID"
![alt tag](http://i.imgur.com/JN90d9v.png)

Then in the third line add what you want to call your bot and you will have to address the bot by this name to run commands.

#Usage
Add your bot to your server using this url https://discordapp.com/oauth2/authorize?client_id=YOUR_BOT_CLIENT_ID&scope=bot&permissions=0

replace "YOUR_BOT_CLIENT_ID" with your client ID which can be found above the Token
![alt tag](http://i.imgur.com/TUoPqZB.png)

Now all you need to do to get your bot to run a command is to address it by its name you set earlier in this example "pepe" and run a command here you can see me running the help command and then the hello command

![alt tag](http://i.imgur.com/P58DeNB.png)
