# MotorWorld

Easy Peasey Web app to control particle device

## INSTALL

Get particle package

In Go Src Path /github.com/BobBurns/ clone particle package

`git clone https://github.com/BobBurns/particle.git`

make sure you are on publish branch

`git checkout publish`

Clone this repo and build moto.go

Flash Particle device that subscribes to your events. See examples/

Put Access token in file .at

`echo <access_token> > .at`

## Usage

Start Server

`xargs -a .at ./moto`

Access Server at `localhost:8082/data`

![Alt text](/data/motoworld.png?raw=true )


### Disclaimer

Take care not to be vulnerable to DNS rebinding attacks when running on your
local LAN
https://medium.com/@brannondorsey/attacking-private-networks-from-the-internet-with-dns-rebinding-ea7098a2d325
