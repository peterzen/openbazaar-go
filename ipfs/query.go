package ipfs

import (
	"github.com/ipfs/go-ipfs/commands"
	"gx/ipfs/QmbkGVaN9W6RYJK4Ws5FvMKXKDqdRQ5snhtaa92qP6L8eU/go-libp2p-routing/notifications"
	peer "gx/ipfs/QmfMmLGoKzCHDN7cGgk64PJr4iipzidDRME8HABSJqvmhC/go-libp2p-peer"
)

func Query(ctx commands.Context, peerID string) ([]peer.ID, error) {
	var peers []peer.ID
	args := []string{"dht", "query", peerID}
	req, cmd, err := NewRequest(ctx, args)
	if err != nil {
		return peers, err
	}
	res := commands.NewResponse(req)
	cmd.Run(req, res)
	resp := res.Output()
	if res.Error() != nil {
		log.Error(res.Error())
		return peers, res.Error()
	}
	peerChan := resp.(<-chan interface{})
	peerMap := make(map[string]peer.ID)
	for p := range peerChan {
		peerMap[p.(*notifications.QueryEvent).ID.Pretty()] = p.(*notifications.QueryEvent).ID
		if len(p.(*notifications.QueryEvent).Responses) > 0 {
			for _, r := range p.(*notifications.QueryEvent).Responses {
				peerMap[r.ID.Pretty()] = r.ID
			}
		}
	}
	for _, v := range peerMap {
		peers = append(peers, v)
	}
	return peers, nil
}
