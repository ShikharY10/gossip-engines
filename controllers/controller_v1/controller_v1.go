package controller_v1

import (
	"encoding/json"
	"gbEngine/admin"
	"gbEngine/handler"
	"gbEngine/models"
	"gbEngine/schema"

	"google.golang.org/protobuf/proto"
)

type Controller struct {
	Handler *handler.Handler
	Logger  *admin.Logger
}

func (ctrl *Controller) MakePartnerRequest(payload *schema.Payload, payloadBytes *[]byte) {
	var request models.PartnerRequest
	err := json.Unmarshal(payload.Data, &request)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
	deliveryId, err := ctrl.Handler.DataBase.GetUserDeliveryId(request.TargetId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	var deliveryPacket schema.DeliveryPacket
	deliveryPacket.Payload = *payloadBytes
	deliveryPacket.TargetId = request.TargetId

	deliveryPacketbytes, err := proto.Marshal(&deliveryPacket)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	err = ctrl.Handler.DataBase.SaveDeliveryPacket(*deliveryId, request.ID, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	nodeName, err := ctrl.Handler.Cache.GetUserConnectNode(request.TargetId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	err = ctrl.Handler.Queue.Produce(nodeName, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
}

func (ctrl *Controller) MakePartnerResponse(payload *schema.Payload, payloadBytes *[]byte) {
	var request models.PartnerResponse
	err := json.Unmarshal(payload.Data, &request)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
	deliveryId, err := ctrl.Handler.DataBase.GetUserDeliveryId(request.TargetId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	var deliveryPacket schema.DeliveryPacket
	deliveryPacket.Payload = *payloadBytes
	deliveryPacket.TargetId = request.TargetId

	deliveryPacketbytes, err := proto.Marshal(&deliveryPacket)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	err = ctrl.Handler.DataBase.SaveDeliveryPacket(*deliveryId, request.ID, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	nodeName, err := ctrl.Handler.Cache.GetUserConnectNode(request.TargetId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	err = ctrl.Handler.Queue.Produce(nodeName, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
}

func (ctrl *Controller) RemovePartnerNotifier(payload *schema.Payload, payloadBytes *[]byte) {}

// func (c *Controller) SendPendingPayloads(payload *schema.Payload) error {
// 	transfers, err := c.Handler.DataBase.GetUserAllDeliveryPacket(string(payload.Data))
// 	if err != nil {
// 		return err
// 	} else {
// 		for _, deliveryPackets := range transfers {
// 			var deliveryPacket schema.DeliveryPacket
// 			err = proto.Unmarshal([]byte(deliveryPackets), &deliveryPacket)
// 			if err != nil {
// 				return err
// 			} else {
// 				targetNodeName, err := c.Handler.Cache.GetUserNodeName(deliveryPacket.TargetId)
// 				if err == nil {
// 					err = c.Handler.Queue.Produce(targetNodeName, deliveryPacket.Payload)
// 					if err != nil {
// 						return err
// 					}
// 				}
// 			}
// 		}
// 		return nil
// 	}
// }

func (c *Controller) BypassChat(payload *schema.Payload, payloadBytes []byte) {
	// var newChat schema.NewMessage
	// err := proto.Unmarshal(payload.Data, &newChat)
	// if err != nil {
	// 	c.logs.RegisterLog(err.Error())
	// } else {
	// 	var savePayload schema.SavePayload
	// 	savePayload.Payload = payloadBytes
	// 	savePayload.TargetId = newChat.TargetId

	// 	savePayloadBinary, err := proto.Marshal(&savePayload)
	// 	if err != nil {
	// 		c.logs.RegisterLog(err.Error())
	// 	} else {
	// 		err = c.handler.MongoDB.InsertPayload(newChat.TransferId, newChat.PayloadKey, string(savePayloadBinary))
	// 		if err != nil {
	// 			c.logs.RegisterLog(err.Error())
	// 		}
	// 	}

	// 	targetNodeName, err := c.handler.RedisDB.GetUserNodeName(newChat.TargetId)
	// 	if err == nil {
	// 		err = c.handler.RabbitMQ.Produce(targetNodeName, payloadBytes)
	// 		if err != nil {
	// 			c.logs.RegisterLog(err.Error())
	// 		}
	// 	}
	// }
}

func (c *Controller) AckPayloadDelivery(payload *schema.Payload, payloadBytes []byte) {
	// var payloadAck schema.PayloadAcknowledgement
	// err := proto.Unmarshal(payload.Data, &payloadAck)
	// if err != nil {
	// 	c.logs.RegisterLog(err.Error())
	// } else {
	// 	err = c.handler.MongoDB.DeleteOnePayload(payloadAck.TransferId, payloadAck.PayloadKey)
	// 	if err != nil {
	// 		c.logs.RegisterLog(err.Error())
	// 	}
	// }
}

// func (c *Controller) MakePartnerRequest(payload *schema.Payload, payloadBytes []byte) {

// }

// func (c *Controller) MakePartnerResponse(payload *schema.Payload, payloadBytes []byte) {}
