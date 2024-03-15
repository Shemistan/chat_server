package converter

import (
	"github.com/Shemistan/chat_server/internal/model"
	pb "github.com/Shemistan/chat_server/pkg/chat_api_v1"
)

// RPCCreateChatToModelChat конвертер из rpc в model
func RPCCreateChatToModelChat(req *pb.CreateRequest) model.Chat {
	if req == nil {
		return model.Chat{}
	}

	return model.Chat{
		Name:  req.GetChatName(),
		Users: req.GetUserLogins(),
	}
}

// RPCSendMessageToModelMessage конвертер из rpc в model
func RPCSendMessageToModelMessage(req *pb.SendMessageRequest) model.Message {
	if req == nil {
		return model.Message{}
	}

	return model.Message{
		ChatName:  req.GetChatName(),
		UserLogin: req.GetUserLogin(),
		Message:   req.GetMessage(),
	}
}
