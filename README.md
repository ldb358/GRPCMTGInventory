# MTG Card Inventory

## What is it:

I build this tool with the goal of inventorying my old MTG collection in order to sell/trade/know what I have. It is a simple tool for building and maintaining a MTG card inventory but with some tools to make the process less painful for someone not super familiar with magic cards.

## Technology:

The second part of this project was an excuse to work the GRPC and Vue.js
which are two technologies that I have been interested in for a while.


The current architecture is:

Backend: Golang + grpc (no database)
Frontend: Vue.js (not started, need research to connect the two)

The next steps(I think):
1. ~~Finishing setting up grpc~~
2. ~~Build client~~
3. ~~Add ssl~~
4. ~~Add rest~~
5. Build out basic inventory + card data protobuf
6. Update the api to reflect the full format
7. Build basic frontend
8. Add JWT token
9. Build out backend with full model implementation, add sql backend of some sort
10. Add features to UI and go from there

## Usage:

To build the client, server, and certificates you can just run:

```
make all
make certs
```

From there you can just run:

`./bin/server`

to start the server. The client can be run with:

`./bin/client`

And that is it for now!


### Resouces:
https://mtg.gamepedia.com/Template:List_of_Magic_sets
https://andrewgioia.github.io/Keyrune/icons.html?icon=lea
https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
https://github.com/grpc-ecosystem/grpc-gateway
