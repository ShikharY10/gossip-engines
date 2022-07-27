package main

import (
	"errors"
	"fmt"
	"gbEngine/gbp"
	"gbEngine/mongoAction"
	"gbEngine/redisAction"
	"gbEngine/rmq"
	"gbEngine/utils"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

var clientNode map[string]string = make(map[string]string)

// TODO:
// 1: Read the job from rabbitmq channel
// 2: Get the name of node on which the client is connected.
// 3: Route the address details to the gateway node on which
// 	  the client is connected.

type MAIN struct {
	RMQ        *rmq.RMQ
	RedisDB    *redisAction.Redis
	MongoDB    *mongoAction.Mongo
	EngineName string
}

func (m *MAIN) HandleJob() {
	var trans gbp.Transport

	for job := range m.RMQ.Msgs {
		fmt.Println("new job")
		err := proto.Unmarshal(job.Body, &trans)
		if err != nil {
			log.Println("[ProtoUNMError1] : ", err.Error())
			continue
		}
		// fmt.Println("trans: ", trans)
		if trans.Tp == 1 {
			// Initial Data
			data, err := m.MongoDB.GetAllMsg(trans.Id)
			if err != nil {
				log.Println("Pending Transfers: ", err.Error())
				continue
			}
			tNodeName, err := m.getNodeName(trans.Id)
			if err != nil {
				log.Println("[tNodeNameError2] : ", err.Error())
				continue

			}
			for _, value := range *data {
				var sF gbp.SaveFormat
				err = proto.Unmarshal(value, &sF)
				if err != nil {
					log.Println("[PROTOUNMERROR] : ", err.Error())
					continue
				}

				m.ProducePayload(int(sF.Tp), sF.Data, trans.Id, tNodeName)
			}
		} else if trans.Tp == 2 {
			// Messaging
			smk := m.MongoDB.GetMainKey(trans.Id)

			pT, err := utils.AesDecryption(utils.Decode(smk), trans.Msg)
			if err != nil {
				log.Println("[AESDECRERRROR1] : ", err.Error())
				continue
			}
			var msgp gbp.ChatPayload
			err = proto.Unmarshal(pT, &msgp)
			if err != nil {
				log.Println("[ProtoUNMError] : ", err.Error())
				continue
			}
			tmk := m.MongoDB.GetMainKey(msgp.Tid)
			Mloc := utils.GenerateRandomId()

			var mF gbp.MsgFormat
			mF.Msg = msgp.Msg
			mF.Mloc = Mloc
			mF.Sid = msgp.Sid

			pMF, err := proto.Marshal(&mF)

			if err != nil {
				log.Println("[PROTOMARSHALERRORs]", err.Error())
				continue
			}

			cipherText, err := utils.AesEncryption(utils.Decode(tmk), []byte(pMF))
			if err != nil {
				log.Println("[AESENCRERROR] : ", err.Error())
				continue
			}
			var sF gbp.SaveFormat
			sF.Tp = 2
			sF.Data = cipherText
			PsF, err := proto.Marshal(&sF)
			if err != nil {
				log.Println("[PROTOMARSHALERROR] : ", err.Error())
			}
			err = m.MongoDB.InsertMsg(msgp.Tid, Mloc, PsF)
			if err != nil {
				log.Println("[MONGOINSERTERROR] : ", err.Error())
				continue
			}

			tNodeName, err := m.getNodeName(msgp.Tid)
			if err != nil {
				log.Println("[tNodeNameError1] : ", err.Error())
				continue
			}
			m.ProducePayload(2, cipherText, msgp.Tid, tNodeName)

			fmt.Println("produced for: ", msgp.Tid)
		} else if trans.Tp == 3 {
			// Message Acknowledgment
			smk := m.MongoDB.GetMainKey(trans.Id)
			packC, err := utils.AesDecryption(utils.Decode(smk), trans.Msg)
			if err != nil {
				log.Println("[AESDECRERRROR3] : ", err.Error())
				continue
			}
			var ackC gbp.ChatAck
			err = proto.Unmarshal(packC, &ackC)
			if err != nil {
				log.Println("[PROTOUNMERROR2] : ", err.Error())
			}
			m.MongoDB.DeleteMsg(ackC.MId, ackC.MLoc)
		} else if trans.Tp == 4 {
			// Hand Shack Process One
			fmt.Println("Four")
			smk := m.MongoDB.GetMainKey(trans.Id)
			plaintext, err := utils.AesDecryption(utils.Decode(smk), trans.Msg)
			if err != nil {
				log.Println("[T4-AESDECRERRROR3] : ", err.Error())
				continue
			}

			var HS1 gbp.HandShackP1
			err = proto.Unmarshal(plaintext, &HS1)

			if err != nil {
				log.Println("[T4-PROTOUNMERROR3] : ", err.Error())
			}
			tMid := m.MongoDB.GetMsgIdByNum(HS1.TargetMobile)
			Mloc := utils.GenerateRandomId()
			HS1.Mloc = Mloc
			plain, err := proto.Marshal(&HS1)
			if err != nil {
				log.Println("[T4-PROTOMARSHALERROR] : ", err.Error())
			}
			targetmainKey := m.MongoDB.GetMainKey(tMid)
			targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), plain)
			if err != nil {
				log.Println("[T4-AESENCERROR] : ", err.Error())
			}

			var sF gbp.SaveFormat
			sF.Tp = 4
			sF.Data = targetCipherText
			PsF, err := proto.Marshal(&sF)
			if err != nil {
				log.Println("[T4-PROTOMARSHALERROR3] : ", err.Error())
			}

			err = m.MongoDB.InsertMsg(tMid, Mloc, PsF)
			if err != nil {
				log.Println("[T4-MONGOINSERTERROR] : ", err.Error())
			}

			tNodeName, err := m.getNodeName(tMid)
			if err == nil {
				m.ProducePayload(4, targetCipherText, tMid, tNodeName)
			}
			fmt.Println("Four Completed")
		} else if trans.Tp == 5 {
			// Hand Shack Process Two
			fmt.Println("Five")
			smk := m.MongoDB.GetMainKey(trans.Id)
			plaintext, err := utils.AesDecryption(utils.Decode(smk), trans.Msg)
			if err != nil {
				log.Println("[T5-AESDECRERRROR3] : ", err.Error())
				continue
			}

			var HS2 gbp.HandShackP2
			err = proto.Unmarshal(plaintext, &HS2)
			if err != nil {
				log.Println("[T5-PROTOUNMERROR3] : ", err.Error())
				continue
			}

			if HS2.Permit == 1 {
				Mloc := utils.GenerateRandomId()
				HS2.Mloc = Mloc
				plain, err := proto.Marshal(&HS2)
				if err != nil {
					log.Println("[PROTOUNMERROR] : ", err.Error())
				}
				targetmainKey := m.MongoDB.GetMainKey(HS2.TargetMID)
				targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), plain)
				if err != nil {
					log.Println("[T5-AESENCERROR] : ", err.Error())
					continue
				}

				var sF gbp.SaveFormat
				sF.Tp = 5
				sF.Data = targetCipherText
				PsF, err := proto.Marshal(&sF)
				if err != nil {
					log.Println("[T5-PROTOMARSHALERROR3] : ", err.Error())
					continue
				}

				err = m.MongoDB.InsertMsg(HS2.TargetMID, Mloc, PsF)
				if err != nil {
					log.Println("[T5-MONGOINSERTERROR] : ", err.Error())
					continue
				}
				tNodeName, err := m.getNodeName(HS2.TargetMID)
				if err == nil {
					m.ProducePayload(5, targetCipherText, HS2.TargetMID, tNodeName)
				}

				fmt.Println("s5.1 complete")
				// s5.1 complete

				senderData, err := m.MongoDB.GetUserDataByMID(HS2.SenderMID)
				if err != nil {
					log.Println("[T5- MONGOGETERROR1] : ", err.Error())
					continue
				}
				targetMloc := utils.GenerateRandomId()
				var targetDataPayload gbp.ConnDataTransfer
				targetDataPayload.Hsid = HS2.Hsid
				targetDataPayload.MID = senderData.MsgId
				targetDataPayload.Mloc = targetMloc
				targetDataPayload.Name = senderData.Name
				targetDataPayload.Number = senderData.PhoneNo
				targetDataPayload.ProfilePic = senderData.ProfilePic
				fmt.Println("targetDataPayload: ", &targetDataPayload)
				tMarshalTdata, err := proto.Marshal(&targetDataPayload)
				if err != nil {
					log.Println("[T5-PROTOMARSHALERROR] : ", err.Error())
					continue
				}
				targetPayloadCipher, err := utils.AesEncryption(utils.Decode(targetmainKey), tMarshalTdata)
				if err != nil {
					log.Println("[T5-AESENCERROR] : ", err.Error())
					continue
				}

				var TsF gbp.SaveFormat
				TsF.Tp = 6
				TsF.Data = targetPayloadCipher
				TPsF, err := proto.Marshal(&TsF)
				if err != nil {
					log.Println("[T5-PROTOMARSHALERROR3] : ", err.Error())
					continue
				}
				err = m.MongoDB.InsertMsg(HS2.TargetMID, targetMloc, TPsF)
				if err != nil {
					log.Println("[T5-MONGOINSERTERROR] : ", err.Error())
					continue
				}
				TNodeName, err := m.getNodeName(HS2.TargetMID)
				if err == nil {
					m.ProducePayload(6, targetPayloadCipher, HS2.TargetMID, TNodeName)
				}

				fmt.Println("s5.2 complete")
				// s5.2 complete

				sendermainKey := m.MongoDB.GetMainKey(HS2.SenderMID)
				targetData, err := m.MongoDB.GetUserDataByMID(HS2.TargetMID)
				if err != nil {
					log.Println("[T5- MONGOGETERROR2] : ", err.Error())
					continue
				}

				senderMloc := utils.GenerateRandomId()
				var senderDataPayload gbp.ConnDataTransfer
				senderDataPayload.Hsid = HS2.Hsid
				senderDataPayload.MID = targetData.MsgId
				senderDataPayload.Mloc = senderMloc
				senderDataPayload.Name = targetData.Name
				senderDataPayload.Number = targetData.PhoneNo
				senderDataPayload.ProfilePic = targetData.ProfilePic
				fmt.Println("senderDataPayload: ", &senderDataPayload)
				sMarshalTdata, err := proto.Marshal(&senderDataPayload)
				if err != nil {
					log.Println("[T5-PROTOMARSHALERROR] : ", err.Error())
					continue
				}
				senderPayloadCipher, err := utils.AesEncryption(utils.Decode(sendermainKey), sMarshalTdata)
				if err != nil {
					log.Println("[T5-AESENCERROR] : ", err.Error())
					continue
				}

				var SsF gbp.SaveFormat
				SsF.Tp = 6
				SsF.Data = senderPayloadCipher
				SPsF, err := proto.Marshal(&SsF)
				if err != nil {
					log.Println("[T5-PROTOMARSHALERROR3] : ", err.Error())
					continue
				}
				fmt.Println("senderMloc: ", senderMloc)
				err = m.MongoDB.InsertMsg(HS2.SenderMID, senderMloc, SPsF)
				if err != nil {
					log.Println("[T5-MONGOINSERTERROR] : ", err.Error())
					continue
				}
				SNodeName, err := m.getNodeName(HS2.SenderMID)
				if err == nil {
					m.ProducePayload(6, senderPayloadCipher, HS2.SenderMID, SNodeName)
				}
				fmt.Println("s5.3 complete")
				m.MongoDB.InsertIntoConnection(senderData.MsgId, targetData.MsgId, "done")
				m.MongoDB.InsertIntoConnection(targetData.MsgId, senderData.MsgId, "done")

				fmt.Println("five completed")
			}
		} else if trans.Tp == 7 {
			var hsRemoveNotify gbp.HandshakeDeleteNotify
			err := proto.Unmarshal(trans.Msg, &hsRemoveNotify)
			if err != nil {
				log.Panicln("[T7-PROTOUNMERROR] : ", err.Error())
				continue
			}
			Mloc := utils.GenerateRandomId()
			targetmainKey := m.MongoDB.GetMainKey(hsRemoveNotify.TargetMID)

			var hsRemoveNotifyNew gbp.HandshakeDeleteNotify
			hsRemoveNotifyNew.Mloc = Mloc
			hsRemoveNotifyNew.Number = hsRemoveNotify.Number
			hsRemoveNotifyNew.SenderMID = hsRemoveNotify.SenderMID
			hsRemoveNotifyNew.TargetMID = hsRemoveNotify.TargetMID

			hsNotifyBytes, err := proto.Marshal(&hsRemoveNotifyNew)
			if err != nil {
				log.Panicln("[T7-PROTOMARSHALERROR] : ", err.Error())
				continue
			}

			targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), hsNotifyBytes)
			if err != nil {
				log.Println("[T7-AESENCERROR] : ", err.Error())
				continue
			}

			var sF gbp.SaveFormat
			sF.Tp = 7
			sF.Data = targetCipherText
			PsF, err := proto.Marshal(&sF)
			if err != nil {
				log.Println("[T7-PROTOMARSHALERROR3] : ", err.Error())
			}

			err = m.MongoDB.InsertMsg(hsRemoveNotify.TargetMID, Mloc, PsF)
			if err != nil {
				log.Println("[T7-MONGOINSERTERROR] : ", err.Error())
			}

			tNodeName, err := m.getNodeName(hsRemoveNotify.TargetMID)
			if err == nil {
				m.ProducePayload(7, targetCipherText, hsRemoveNotify.TargetMID, tNodeName)
			}
		} else if trans.Tp == 8 {
			fmt.Println("profile pic change")
			var changeProfile gbp.ChangeProfilePayloads
			err := proto.Unmarshal(trans.Msg, &changeProfile)
			if err != nil {
				log.Panicln("[T8-PROTOUNMERROR] : ", err.Error())
				continue
			}
			for _, mid := range changeProfile.All {
				fmt.Println("notify multiple...")
				Mloc := utils.GenerateRandomId()
				targetmainKey := m.MongoDB.GetMainKey(mid)

				var change gbp.ChangeProfilePayload
				change.Mloc = Mloc
				change.PicData = changeProfile.PicData
				change.SenderMID = changeProfile.SenderMID
				change.TargetMID = mid

				changeBytes, err := proto.Marshal(&change)
				if err != nil {
					fmt.Println("[T8-PROTOMARSHALERROR] : ", err.Error())
				}

				targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), changeBytes)
				if err != nil {
					log.Println("[T8-AESENCERROR] : ", err.Error())
					continue
				}

				var sF gbp.SaveFormat
				sF.Tp = 8
				sF.Data = targetCipherText
				PsF, err := proto.Marshal(&sF)
				if err != nil {
					log.Println("[T8-PROTOMARSHALERROR3] : ", err.Error())
				}

				err = m.MongoDB.InsertMsg(mid, Mloc, PsF)
				if err != nil {
					log.Println("[T8-MONGOINSERTERROR] : ", err.Error())
				}

				tNodeName, err := m.getNodeName(mid)
				if err == nil {
					m.ProducePayload(8, targetCipherText, mid, tNodeName)
				}
				fmt.Println("notified all about oic update")
			}
		} else if trans.Tp == 9 {
			var notifyNum gbp.NotifyChangeNumbers
			err := proto.Unmarshal(trans.Msg, &notifyNum)
			if err != nil {
				log.Panicln("[T9-PROTOUNMERROR] : ", err.Error())
				continue
			}
			for _, mid := range notifyNum.All {
				Mloc := utils.GenerateRandomId()
				targetmainKey := m.MongoDB.GetMainKey(mid)

				var notify gbp.NotifyChangeNumber
				notify.Mloc = Mloc
				notify.Number = notifyNum.Number
				notify.SenderMID = notifyNum.SenderMID
				notify.TargetMID = mid

				notifyBytes, err := proto.Marshal(&notify)
				if err != nil {
					fmt.Println("[T9-PROTOMARSHALERROR] : ", err.Error())
				}

				targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), notifyBytes)
				if err != nil {
					log.Println("[T9-AESENCERROR] : ", err.Error())
					continue
				}

				var sF gbp.SaveFormat
				sF.Tp = 9
				sF.Data = targetCipherText
				PsF, err := proto.Marshal(&sF)
				if err != nil {
					log.Println("[T9-PROTOMARSHALERROR3] : ", err.Error())
				}

				err = m.MongoDB.InsertMsg(mid, Mloc, PsF)
				if err != nil {
					log.Println("[T9-MONGOINSERTERROR] : ", err.Error())
				}

				tNodeName, err := m.getNodeName(mid)
				if err == nil {
					m.ProducePayload(9, targetCipherText, mid, tNodeName)
				}
			}
		} else if trans.Tp == 10 {
			var loginPayload gbp.LoginEnginePayload
			err := proto.Unmarshal(trans.Msg, &loginPayload)
			if err != nil {
				log.Panicln("[T10-PROTOUNMERROR] : ", err.Error())
				continue
			}
			for _, mid := range loginPayload.AllConn {

				Mloc := utils.GenerateRandomId()
				targetmainKey := m.MongoDB.GetMainKey(mid)

				var connReq gbp.LKeyShareRequest
				connReq.TargetMid = mid
				connReq.SenderMid = loginPayload.SenderMid
				connReq.PublicKey = loginPayload.PublicKey
				connReq.Mloc = Mloc

				connReqbytes, err := proto.Marshal(&connReq)
				if err != nil {
					log.Println("[T10-PROTOMARSHALERROR] : ", err.Error())
				}

				targetCipherText, err := utils.AesEncryption(utils.Decode(targetmainKey), connReqbytes)
				if err != nil {
					log.Println("[T9-AESENCERROR] : ", err.Error())
					continue
				}

				var sF gbp.SaveFormat
				sF.Tp = 10
				sF.Data = targetCipherText
				PsF, err := proto.Marshal(&sF)
				if err != nil {
					log.Println("[T9-PROTOMARSHALERROR3] : ", err.Error())
				}

				err = m.MongoDB.InsertMsg(mid, Mloc, PsF)
				if err != nil {
					log.Println("[T9-MONGOINSERTERROR] : ", err.Error())
				}

				tNodeName, err := m.getNodeName(mid)
				if err == nil {
					m.ProducePayload(10, targetCipherText, mid, tNodeName)
				}
			}
		} else if trans.Tp == 11 {
			smk := m.MongoDB.GetMainKey(trans.Id)
			plaintext, err := utils.AesDecryption(utils.Decode(smk), trans.Msg)
			if err != nil {
				log.Println("[T11-AESDECRERRROR3] : ", err.Error())
				continue
			}

			var connKey gbp.ConnectionKey
			err = proto.Unmarshal(plaintext, &connKey)
			if err != nil {
				log.Println("[T5-PROTOUNMERROR3] : ", err.Error())
				continue
			}

			tmk := m.MongoDB.GetMainKey(connKey.TargetMid)
			Mloc := utils.GenerateRandomId()

			var newConnKey gbp.ConnectionKey
			newConnKey.Key = connKey.Key
			newConnKey.Mloc = Mloc
			newConnKey.Number = connKey.Number
			newConnKey.SenderMid = connKey.SenderMid
			newConnKey.TargetMid = connKey.TargetMid

			newConnKeyBytes, err := proto.Marshal(&newConnKey)
			if err != nil {
				log.Println("[T11-PROTOMARSHALERROR]", err.Error())
			}

			cipherText, err := utils.AesEncryption(utils.Decode(tmk), newConnKeyBytes)
			if err != nil {
				log.Println("T11-AESENCERROR", err.Error())
			}

			var sF gbp.SaveFormat
			sF.Tp = 11
			sF.Data = cipherText
			PsF, err := proto.Marshal(&sF)
			if err != nil {
				log.Println("[T9-PROTOMARSHALERROR3] : ", err.Error())
			}

			err = m.MongoDB.InsertMsg(connKey.TargetMid, Mloc, PsF)
			if err != nil {
				log.Println("[T9-MONGOINSERTERROR] : ", err.Error())
			}

			tNodeName, err := m.getNodeName(connKey.TargetMid)
			if err == nil {
				m.ProducePayload(11, cipherText, connKey.TargetMid, tNodeName)
			}

		}
	}
}

func (m *MAIN) getNodeName(Mid string) (string, error) {
	var tNodeName string = ""
	tNodeName = clientNode[Mid]
	if tNodeName == "" {
		tNodeName = m.RedisDB.Client.Get(Mid).Val()
		if tNodeName == "" {
			return "", errors.New("no node names found")
		}
		clientNode[Mid] = tNodeName
		return tNodeName, nil
	}
	clientNode[Mid] = tNodeName
	return tNodeName, nil
}

func (m *MAIN) ProducePayload(tp int, cData []byte, TMid string, tNodeName string) {
	var trans gbp.Transport
	trans.Id = TMid
	trans.Msg = cData
	trans.Tp = int32(tp)
	ptrans, err := proto.Marshal(&trans)
	if err != nil {
		log.Println("[PROTOUNMERROR] : ", err.Error())
	}
	var sD gbp.SendNotify
	sD.Data = ptrans
	sD.TMid = TMid
	sDBytes, err := proto.Marshal(&sD)
	if err != nil {
		log.Println("[ProtoMarshalError2] : ", err.Error())
	}
	err = m.RMQ.Produce(tNodeName, sDBytes)
	if err != nil {
		log.Println("[RMQProduceError] : ", err.Error())
	}
	fmt.Println("Produced for: ", tNodeName)
}

func main() {

	var mongoIP string = "127.0.0.1"
	var rabbitIP string = "127.0.0.1"
	var redisIP string = "127.0.0.1"

	var rabbitUsername string = "guest"
	var rabbitPassword string = "guest"

	var mongoUsername string = "rootuser"
	var mongoPassword string = "rootpass"

	// ---------------------------------------
	val, found := os.LookupEnv("MONGO_LOC_IP")
	if found {
		mongoIP = val
	}

	val, found = os.LookupEnv("RABBITMQ_LOC_IP")
	if found {
		rabbitIP = val
	}

	val, found = os.LookupEnv("REDIS_LOC_IP")
	if found {
		redisIP = val
	}
	// -----------------------------------------

	// -----------------------------------------
	val, found = os.LookupEnv("RABBITMQ_USERNAME")
	if found {
		rabbitUsername = val
	}

	val, found = os.LookupEnv("RABBITMQ_PASSWORD")
	if found {
		rabbitPassword = val
	}
	// -------------------------------------------

	// -------------------------------------------
	val, found = os.LookupEnv("MONGO_USERNAME")
	if found {
		mongoUsername = val
	}

	val, found = os.LookupEnv("MONGO_PASSWORD")
	if found {
		mongoPassword = val
	}
	// -------------------------------------------

	var name string = "EN" + utils.Encode(utils.GenerateAesKey(10))

	var m MAIN
	m.EngineName = name

	var mongodb mongoAction.Mongo
	mongodb.Init(mongoIP, mongoUsername, mongoPassword)

	var redisdb redisAction.Redis
	redisdb.Init(redisIP)

	var RMQ rmq.RMQ
	RMQ.Name = name
	RMQ.Address = rabbitIP
	RMQ.Port = "5672"
	RMQ.Username = rabbitUsername
	RMQ.Password = rabbitPassword
	RMQ.Init()

	m.RMQ = &RMQ
	m.RedisDB = &redisdb
	m.MongoDB = &mongodb

	redisdb.SetEngineName(name)
	m.HandleJob()
}
