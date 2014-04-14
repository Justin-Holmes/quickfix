package quickfix

import (
	"github.com/quickfixgo/quickfix/message"
)

//The Application interface should be implemented by FIX Applications.
//This is the primary interface for processing messages from a FIX Session.
type Application interface {
	//Notification of a session begin created.
	OnCreate(sessionID SessionID)

	//Notification of a session successfully logging on.
	OnLogon(sessionID SessionID)

	//Notification of a session logging off or disconnecting.
	OnLogout(sessionID SessionID)

	//Notification of admin message being sent to target.
	ToAdmin(msgBuilder message.MessageBuilder, sessionID SessionID)

	//Notification of app message being sent to target.
	ToApp(msgBuilder message.MessageBuilder, sessionID SessionID) error

	//Notification of admin message being received from target.
	FromAdmin(msg message.Message, sessionID SessionID) message.MessageReject

	//Notification of app message being received from target.
	FromApp(msg message.Message, sessionID SessionID) message.MessageReject
}
