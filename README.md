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
1. Finishing setting up grpc
2. Build clients
3. Add ssl
4. Add rest
5. Build basic frontend
6. Build out backend with full model implementation, add sql backend of some sort
7. Add features to UI and go from there