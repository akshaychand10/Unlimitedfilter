// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
<b>Hᴇʏ %v ɪᴍ %v ᴀɴ Aᴡᴇsᴏᴍᴇ Filter bot with global filter support</b>

<b>I can save a custom reply for a word in any chat. Check my /help menu for more details.</b>
	`,
	"ABOUT": `
Creator: <a href='http://t.me/iamLegend789bot'>༺Leͥgeͣnͫd༻</a>
Library: <a href='https://docs.pyrogram.org'>Pyrogram 1.4.16</a>
Updates: <a href='https://t.me/AKprojects4'>AK projects</a>
Server: <a href='https://www.digitalocean.com'>Do</a>
Build Status: v1.0.3 [stable]
	`,

	"MF": `
Filters:
Filter is the feature were users can set automated replies for a particular keyword and the bot will respond whenever a keyword is found the message

NOTE:
1. bot should have admin privillage in order to reply filters in a chat.
2. only admins can add filters in a chat.
3. filters does support all the telegram markdowns, medias 
5. there are some easter eggs, try to find it out.

Commands and Usage:
/filter   - add a filter

/filters - list all the filters of a chat

/stop  - delete a specific filter (separate keywords with spaces for deleting multiple filters at a time)
`,

	"GF": `
Connections:
Used to connect bot to PM which let will you to execute both normal filter related commands and some other sensitive commands right from the PM that will
reflect in the group which helps you to keep the filter additions and other stuffs private and helps to prevent flooding.

NOTE:
1. Only admins can add a connection.
2. In a chat you can simply use the /connect for starting a connection and in PM you must specify chat id right after the command.

Commands and Usage:

/connect <chat id> - connect a particular chat to your PM

/disconnect <chat id> - disconnect from a chat
`,
	"CONNECT": `
Connections:
Used to connect bot to PM which let will you to execute both normal filter related commands and some other sensitive commands right from the PM that will
reflect in the group which helps you to keep the filter additions and other stuffs private and helps to prevent flooding.

NOTE:
1. Only admins can add a connection.
2. In a chat you can simply use the /connect for starting a connection and in PM you must specify chat id right after the command.

Commands and Usage:

/connect <chat id> - connect a particular chat to your PM

/disconnect <chat id> - disconnect from a chat
`,

	"BROADCAST": `
<b>The broadcast feature allows bot admins to broadcast a message to all of the bot's users.</b>

<I>Broadcasts are of two types :</i>
 ~ Full Broadcast - Broadcast to all of the bot users 

   <code>/broadcast</code>
 ~ Concast - Broadcast to only users who are connected to a chat 

   <code>/concast</code>
`,

	"HELP": `
<b>To know how to use my modules use the buttons below to get all my commands with usage examples !</b>
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{
			{Text: "🧭 Help 🧭", CallbackData: "edit(HELP)"},
		},
	},
	"ABOUT": {
		{
                        {Text: "Home", CallbackData: "edit(HELP)"},
		},
	},
	"STATS": {
		{
			{Text: "Back", CallbackData: "edit(ABOUT)"},
			{Text: "Refresh", CallbackData: "stats"},
		},
	},
	"HELP": {
		       {{Text: "Filter", CallbackData: "edit(MF)"},
			{Text: "About", CallbackData: "edit(ABOUT)"},
		}, {
			{Text: "Connection", CallbackData: "edit(GF)"}, 
                        {Text: "Broadcast", CallbackData: "edit(BROADCAST)"},
                }, {
                        {Text: "Update Channel", Url: "t.me/AKprojects4"},
                        {Text: "🫂 Support 🫂", Url: "t.me/AK_projects_Bot"},
                },
		       {{Text: "Back", CallbackData: "edit(START)"}},
	},
}
