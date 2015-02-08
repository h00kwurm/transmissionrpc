Transmission RPC
==========

[transmission](https://www.transmissionbt.com/) is an excellent bittorrent client that provides an 
rpc service that can be used for pretty much everything. This provides a simple way to interact
with it using Go. RPC commands are wrapped to be type-strict. Ya know the usual expectations.

I'm doing implementation in a write-as-I-need-it way and I will absolutely maintain this. I am using
this in my own flexget implementation (which is coming soon).

## Example
There is folder containing what I believe to be simple examples I will build upon.

    transmission := transmissionrpc.New("http://192.168.0.106", "9091")
    torrents, err := transmission.GetTorrents()
    if err != nil {
        fmt.Println(err)
    }

## RPC documentation
For reference, you can find the transmission [rpc spec here](https://trac.transmissionbt.com/browser/trunk/extras/rpc-spec.txt).
It doesn't exactly comply with json-rpc (uses "arguments" instead of "params").

## Requests
Either comment to have me write it or (preferably) clone-write-addtoexample-pullrequest.

## License
MIT, attached.
Aditya Natraj. self at anatraj dot com 2015

