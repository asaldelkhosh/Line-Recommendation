# Broadcaster
Building a WebRTC video and audio Broadcaster in Golang using ION-SFU, and media devices

### What is this project?

### What is WebRTC ?
WebRTC or **Web Real-Time Communication** is a free open-source project providing 
web browsers and mobile applications with **real-time communication** (RTC) with APIs.
It allows audio and video communication to work inside web pages by allowing direct
**peer-to-peer** communication, eliminating the need to install plugins or download native
apps.

The technology behind WebRTC are implemented as an open web standard and available
as regular Javascript APIs in all major browsers.

As said before, major components of WebRTC includes several JS APIs:
- **getUserMedia** acquires the audio and video media
- **RTCPeerConnection** enables audio and video communication between peers. It performs signal processing, codec handling, peer-to-peer communication and ...
- **RTCDataChannel** allows bidirectional communication of arbitrary data between peers.

<p align="center">
    <img src="https://www.researchgate.net/profile/Martin-Meszaros-3/publication/328334940/figure/fig18/AS:682651635707904@1539768241002/WebRTC-triangle-with-SDES-and-DTLS-key-exchange-As-discussed-in-section-35-DTLS-SRTP.png" width="400" />
</p>

### What is SFU ?
SFU stands for **Selective Forwarding Unit**. Also known in the specifications as SFM (Selective Forwarding Middlebox).
At times, the term is used to describe a type of video routing device, while at other times
it will be used to indicate the support of routing technology and not a specific device.

An SFU is capable of receiving multiple media streams and then decide which of these media
streams should be sent to which participants.

SFU is a video routing service which allows webrtc sessions to scale more efficiently.

<p align="center">
    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSTiIJSpXcglL4a1_Z8pKNVx7kH7B8nh72xkXPC2xv-5tR6bmNkq2lMjM4PR2URk2HADws&usqp=CAU" width="392" />
</p>

### Why using ION-SFU ?
This package provides a simple, flexible, high performance Go implementation of a WebRTC SFU. 
It can be called directly or through a gPRC or json-rpc interface.

ION-SFU supports real-time processing on media streamed through the sfu using ion-avp.

Ion-avp is an extensible audio/video processing service designed for use with ios-sfu.

### What is media devices ?
The MediaDevices interface provides access to connected media input devices like cameras
and microphones, as well as screen sharing. In essence, it lets you obtain access to any 
hardware source of media data.

### How does this project work ?
WebRTC serves multiple purposes; together with the Media Capture and Streams API, they provide
powerful multimedia capabilities to the web, including support for audio and video conferencing,
file exchange, screen sharing, identity management. Connection between peers can be made
without requiring any special drivers or plugins.

Connections between two peers are represented by the **RTCPeerConnection** interface.
Once a connection has been established and opened using RTCPeerConnection, media streams (**MediaStream**)
and data channels (**RTCDataChannel**) can be added to the connection.

Media streams can consist of any number of tracks of media information; **tracks**, which are
represented by objects based on the **MediaStreamTrack** interface, may contain one of a number
of types of media data, including audio, video, and text.

Most streams consist of at least one audio track and likely also a video track, and
can be used to send and receive both live media or stored media information (such as streamed movie).

First we create a WebRTC connection between our local computer and a remote peer.
then we use bidirectional channels to transfer data between peers.

Once a user runs the application, we use Media devices to get the input data from our user
and send it to our ION-SFU server.

ION-SFU server on the other hand, gets the data and returns the response to all other peers, including our local machine.

### Dialer package
websocket - dialer

### Config package 
ICE server - SDPSemantics

### Engine package
Media engine - vpx - code selector

### How to use this project ?