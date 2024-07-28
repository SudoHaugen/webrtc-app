package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

const signalingServerURL = "ws://localhost:8080/ws"

var ws *websocket.Conn

func main() {
	var err error
	ws, _, err = websocket.DefaultDialer.Dial(signalingServerURL, nil)

	if err != nil {
		log.Fatalf("Failed to connect to signaling server: %v", err)
	}

	defer ws.Close()

	//Create new Peer Connection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})

	if err != nil {
		log.Fatalf("Failed to create PeerConnection: %v", err)
	}

	// Handle of ICE candidates
	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}

		sendToSignalingServer(c.ToJSON())
	})

	//Create offer and start connection
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Fatalf("Failed to create offer: %v", err)
	}

	//set the local descritption to the offer
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		log.Fatalf("Failed to set local description: %v", err)
	}

	sendToSignalingServer(offer)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		handleMessage(msg, peerConnection)
	}
}

func sendToSignalingServer(data interface{}) {
	err := ws.WriteJSON(data)

	if err != nil {
		log.Fatalf("Failed to send to signaling server: %v", err)
	}
}

func handleMessage(msg []byte, peerConnection *webrtc.PeerConnection) {
	var signalData map[string]interface{}
	err := json.Unmarshal(msg, &signalData)
	if err != nil {
		panic(err)
	}

	log.Printf("Handling message: %v", signalData)

	if signalData["type"] == "offer" {
		offer := signalData["sdp"].(string)
		err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  offer,
		})

		if err != nil {
			panic(err)
		}

		answer, err := peerConnection.CreateAnswer(nil)

		if err != nil {
			panic(err)
		}

		err = peerConnection.SetLocalDescription(answer)
		if err != nil {
			panic(err)
		}

		sendToSignalingServer(answer.SDP)
	} else if signalData["type"] == "answer" {
		answer := signalData["sdp"].(string)
		err := peerConnection.SetRemoteDescription(webrtc.SessionDescription{
			Type: webrtc.SDPTypeAnswer,
			SDP:  answer,
		})

		if err != nil {
			panic(err)
		}
	} else if signalData["type"] == "candidate" {
		candidate := signalData["canidate"].(map[string]interface{})
		sdpMLineIndex := uint16(candidate["sdpMLineIndex"].(float64))
		sdpMid := candidate["sdpMid"].(string)

		err := peerConnection.AddICECandidate(webrtc.ICECandidateInit{
			Candidate:     signalData["sdpMid"].(string),
			SDPMLineIndex: &sdpMLineIndex,
			SDPMid:        &sdpMid,
		})

		if err != nil {
			panic(err)
		}
	}
}
