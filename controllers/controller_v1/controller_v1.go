package controller_v1

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gbEngine/admin"
	"gbEngine/handler"
	"gbEngine/models"
	"gbEngine/schema"
	"log"

	"google.golang.org/protobuf/proto"
)

type Controller struct {
	Handler *handler.Handler
	Logger  *admin.Logger
}

func (ctrl *Controller) Test_API_ENGINE(payload *schema.Payload, payloadBytes *[]byte) {
	fmt.Println("Calling Test_API_ENGINE")
	type testBody struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		DeliveryId string `json:"deliveryId"`
		SenderId   string `json:"senderId"`
		TargetId   string `json:"targetId"`
	}

	var request testBody
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
	fmt.Println("2")

	request.DeliveryId = deliveryId.Hex()

	jsonEncoded, err := json.Marshal(&request)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	var newPayload schema.Payload
	newPayload.Data = jsonEncoded
	newPayload.Type = payload.Type

	newPayloadBytes, err := proto.Marshal(&newPayload)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	log.Println("Payload length: ", len(newPayloadBytes))
	log.Println("TargetId: ", request.TargetId)

	var deliveryPacket schema.DeliveryPacket
	deliveryPacket.Payload = newPayloadBytes
	deliveryPacket.TargetId = request.TargetId

	deliveryPacketbytes, err := proto.Marshal(&deliveryPacket)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
	fmt.Println("3")

	err = ctrl.Handler.DataBase.SaveDeliveryPacket(*deliveryId, request.ID, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
	fmt.Println("4")

	nodeName, err := ctrl.Handler.Cache.GetUserConnectNode(request.TargetId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	fmt.Println("nodename: ", nodeName)

	err = ctrl.Handler.Queue.Produce(nodeName, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	fmt.Println("5")
}

func (ctrl *Controller) AckPayloadDelivery(payload *schema.Payload, payloadBytes *[]byte) {
	type ackPayload struct {
		DeliveryId string `json:"deliveryId"`
		PacketId   string `json:"packetId"`
	}
	var ack ackPayload
	err := json.Unmarshal(payload.Data, &ack)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	err = ctrl.Handler.DataBase.RemoveDeliveryPacket(ack.DeliveryId, ack.PacketId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
}

func (ctrl *Controller) DeliverPendingPacket(payload *schema.Payload, payloadBytes *[]byte) {
	type PendingPacket struct {
		UserId string `json:"userId"`
	}

	var pendingPacket PendingPacket
	err := json.Unmarshal(payload.Data, &pendingPacket)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	deliveryId, err := ctrl.Handler.DataBase.GetUserDeliveryId(pendingPacket.UserId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	packets, err := ctrl.Handler.DataBase.GetAllDeliveryPackets(*deliveryId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	nodeName, err := ctrl.Handler.Cache.GetUserConnectNode(pendingPacket.UserId)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	for _, packet := range packets {
		packetBytes, err := hex.DecodeString(packet)
		if err != nil {
			ctrl.Logger.LogError(err)
			return
		}
		err = ctrl.Handler.Queue.Produce(nodeName, packetBytes)
		if err != nil {
			ctrl.Logger.LogError(err)
			return
		}
	}
}

func (ctrl *Controller) MakePartnerRequest(payload *schema.Payload, payloadBytes *[]byte) {
	fmt.Println("make partner called")
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

	err = ctrl.Handler.DataBase.RemoveMakePartnerRequest("partnerrequests", request.ResponserId, request.ID)
	if err != nil {
		ctrl.Logger.LogError(err)
		// return
	}

	err = ctrl.Handler.DataBase.RemoveMakePartnerRequest("partnerrequested", request.TargetId, request.ID)
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

func (ctrl *Controller) RemovePartnerNotifier(payload *schema.Payload, payloadBytes *[]byte) {
	var removePartner models.RemovePartner
	err := json.Unmarshal(payload.Data, &removePartner)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	// remove extraditeeId from extraditor's partner section
	err = ctrl.Handler.DataBase.RemovePartner(removePartner.ExtraditorId, removePartner.ExtraditeeId)
	if err != nil {
		return
	}

	// remove extractorId from extraditee's partner section
	err = ctrl.Handler.DataBase.RemovePartner(removePartner.ExtraditeeId, removePartner.ExtraditorId)
	if err != nil {
		return
	}

	// get the deliveryId of extraditee
	deliveryId, err := ctrl.Handler.DataBase.GetUserDeliveryId(removePartner.ExtraditeeId)
	if err != nil {
		return
	}

	// create the delivery payload
	var deliveryPacket schema.DeliveryPacket
	deliveryPacket.Payload = *payloadBytes
	deliveryPacket.TargetId = removePartner.ExtraditeeId

	deliveryPacketbytes, err := proto.Marshal(&deliveryPacket)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	// save the delivery payload in extraditee's payload section using deliveryId
	err = ctrl.Handler.DataBase.SaveDeliveryPacket(*deliveryId, removePartner.ID, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}

	// check if extraditee is connected to any of the gateways
	nodeName, err := ctrl.Handler.Cache.GetUserConnectNode(removePartner.ExtraditeeId)
	if err != nil {
		return
	}

	// if connected then send the delivery payload immediately
	err = ctrl.Handler.Queue.Produce(nodeName, deliveryPacketbytes)
	if err != nil {
		ctrl.Logger.LogError(err)
		return
	}
}

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

// func (c *Controller) AckPayloadDelivery(payload *schema.Payload, payloadBytes []byte) {
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
// }

// func (c *Controller) MakePartnerRequest(payload *schema.Payload, payloadBytes []byte) {

// }

// func (c *Controller) MakePartnerResponse(payload *schema.Payload, payloadBytes []byte) {}
