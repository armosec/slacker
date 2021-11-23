# slacker
a simplistic slack bot

example:
'''

	bot, _ := slacker.SlackBotInit("A1110001114", "xoxb-1337133713375-1337133713371-13371337133713371337RFI", true)
	bot.SendCriticalMessage("", "scan failed")
'''

channelID:
A1110001114

token:
"xoxb-1337133713375-1337133713371-13371337133713371337RFI"

can work with user & bot tokens:
bots must be able to:


channels:join
channels:read
groups:read
chat:write 
im:read
mpim:read


