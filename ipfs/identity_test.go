package ipfs

import (
	"bytes"
	"encoding/hex"
	"github.com/tyler-smith/go-bip39"
	"gx/ipfs/QmfWDLQjGjVe4fr5CoztYW2DYYjRysMJrFe1RCsXLPTf46/go-libp2p-crypto"
	"testing"
)

var keyHex string = "080012ac12308209280201000282020100b962ce316a22e4aa1393c92b84925bf27cdd174edb06338ccc97a37698a83795cf5ccce76fb43734748d10d49a1c8ebe03ab7ac18ae18175e6475a7d6674d872209117d5d3c4f11f1994558a07fdbba490583c3af968b38c7abea26f5ee75472eb0829c919ad715153329bca0d1a9686b267f638c10d2f9328014d74aab216aeff00b772da0c567d225f5c2ffd1adc75d0848b9ced11c4c4a30299367a36a0bb34af7d3fbdf260d65b9a2bace264b964d063df396fb87705ae8e2a20aa7f382a0f268574bf2ab44205335e1f680b06f35c3c9b9653acf57e7019757cc0ecf5ac78afc0382aaecdedd7f16109282a00cf783f552931ffebafe5e65f3e1c3262a71a90c694f5c3df69faf750c9542f6b23718ec30dd476990b5d360cba343f50b98a915d5ac54d0f480775d2dd0a4d7f2df17f694b9f2900f0481c2eecf169429899af067122ae216c543f46067d145b7fb3a3763278363333a21457668288d393fe7474735fb62975ac9ff116c7963cc80bc59b57f6777369c817d0972bb8ee12861dd9b0698cdbec72e98619227c2fab4b68ded217a5d2b72f44fd7a19959907c5941bc0b33e5089de1f4627f5ac5c3b24f35970e47d549daa42fef8c9b116a254d87215ccb162a5f37f75d4b08d72153cc3eae01b6f042f9673ae15c2408c5978431b8a81ea4cebf90a2bcaf52d40f99649718b7314b5752ab2ef88aaeb31e102030100010282020061b722d80909bb5daa540b8554ab90ab90053973e71eadad5969d44f7ec7ac04daf81f491e27efb1467dd4913b6a321315fc1d08471bf6111f001d425beedae7cd22ad5d97f206e64da552311f652eb2df22a31cb2b9ed69ee88ed35a6e06832ea8fb2f6bd021ee8aee1ce1c022c318b6e1e3e3a4c07973c09aa2619889fb6352754495dd071106584e89ebeecba6bd9edbeef98179d7609efa23effe80b434d69c7d64cf2e1ff08fddf9585db0f1c69772e2d4b290822346c7da2d6f50779a331b92c3f0d44851f077ebbedf1f95e00886da9d4e2e617c03d762de74409a2b4f6282d396dc615e5781d52f870a3332a2847b9abad1f282bdf123ec18e3b40a3ce3c5683564fcc4fe71b9dcb7bd61d5e6264910ee532ca970a7a620fb42f1e8bd9de7767cc3632144afdc0b132e133b4abf1b2b62a7e2fdc6dbc86e67dc000964712e328089ca5ee31efae17897d49551035721243f956b4f487d9e6fb34e4c61eeb47745f9eb4e089e182c215ff7864a47f799f862808e1a7867c9e4b88de46be27285777930b79a7e50cfa7c33372329a52f1b3f466c6742deb48d1dbef27682ab69a0e0054f06656b16459c7cb7027af2c84b6e37ff6124f207896799f686027c2ece29ba38d4aef254c36292a29ee07436f86d8074a65ad575ef839f75fb195f717e565760935f1f5192164f828750498c13d70cb07520b3a927e89754010282010100dd82b0ed9c4c7f7fb0525f49f242bee7db4ba9164bfac7bb4555f14067ebef9f12a9e67eb3da39fb9e30930a3837c7ae57e8efba26f2ebfdc92cbbfbebe9903bfe9f451775aa3bbed2195b1d354805373136f408bdbe7e58a8f43e2cc594f1d00cfcddf8875091515e3dde700b7476ce3dadf560699bcd80b0022961b4a6ae6e4831aaba8fbd8b1f18076287d38a09f32ddf7a710d4baf049b80dfce7e1d1c19a028106e4bd4cb6cbb6fe9dcf75104c229f2234fb0b96203320cb7541869a23ac3badabd38587e6962d33245e931c2f8a5454e494e3eccf3e8fc034208b955bbdbb8704f59e2823736358fea8d58c621028889780a40b69c41878ace91d429a90282010100d64034b8ebfff3a8eeaa1fcea3026c0d1ba2a56233eda6510119490d41d7367448df6450d60ee373167db8e50df25546aff37246e8597563288c2b2ea52b48e747a2f2988d15934eb5be48c495c8a3ee7108c5749f3a7184afdee50fd80499d40ab680c893b853419a5f9b93b2c0cace8c7da2a783ca57ef9cd393ba34f909cfe742cffe4528d0e2cab529dafe36e52b01d8e3d7a0376df1121f01b7cda808ce72f46e5bb322806717a58d057d368eda32388faa76427bb99f515735d52e91748b0d564a864b1e3a22d26ac35ef80b3b65e87155abec052508ae714836280bd9da00ce5862cced9b0869c4e47f4bb2ffcdd91f59b349cd4528e5996808421979028201006ef3f6a19e47a719937d3f23b00884cd6847f6b3276dbf3beb9807d6c5f725389dc5a2a1d67eb9ca724b4c6bb30b25a70b7baf10a44fcbab7c683ee50a1f1cd3205030f0764cc5ce6838b12de93161036b50665e3684c38eaf5a60065715ab26164b56f814f76342c99cdddd6baf738f0cd024f79d0ea09a140f9120d2b726ea3099483658a78d970e972c6072b253178ee3c0d5fdddea15b9ba8ce046c557fcf1feff73efe1efdf4fe7d7d189c3f849356f870674b3f70738bbabdca0b9d0bbf4ac94c27105ca94091f56b034056beed601e81bc9e8a50a8baa280b9c89110b9f1c3c8d52ada5a1f8324e1b239fd2e657f9823a319086a7b01f2f85170099b902820101009b8d9639c4aeda5b869c61e787169518973de951795e4be87821c1962c7d727dcf02d2349695ccb19c5cbf471e6f1956ef1c9395c3d05f9814b2600fc7bfbb789fdc40c25d7c92ad0ceb10ae1a09c86bb009ac42d5f07ac6d961bd7159674dc053ded975c73a86c814aae1abc4969128917c0e020d7f7584f499f0a1bb202ced937c40c2ddb79aa16e8745d247cbd76af8365e86093ef54597e08740c9c1fc52d7c77bf07c08caad0ed015a92c95d5f5b98cb35211f68719e800088f7f2e0651553be3b0cbfbf2004793fa0bcb01d035d9f17f87b200f46345005223199ecdabe7c9ba071604af66661f54e289418ebe86d7e3191f40b0164c7116861fa8ce11028201005be415091cf086faf22b75d3748f680f860b8540c732574c2ea80d0bf00429891fe6606b6801ad819ad83a1db39be043df0c0c4e60d3bab734ab1649601e3337eb507119a47dd6dfbcff338e68f3b7d3b302de0d1d9ccf94b02ada88d1bb47a60adbbdb7bf51d34f953036dbfda662f6c52b734729a7fa6df27b4a69a723546173a1545800b91e606c2db8d911a94665ebe2b0582c515657e857a1503a68ab167296f7b61fbdcba6a60888e74021ef10aa4af70e5356e432bafd29cbc1c7b8b7695d11bf9a6633f057bd8dc82020339bf197b8a79f5b4bca4cebfd65f7c5e4ff5c20fe5932c5768b3af78596aff5450ec241e33e65fcb07651763fc1fcf20754"

func TestIdentityFromKey(t *testing.T) {
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		t.Error(err)
	}
	identity, err := IdentityFromKey(keyBytes)
	if err != nil {
		t.Error(err)
	}
	if identity.PeerID != "QmdHkAQeKJobghWES9exVUaqXCeMw8katQitnXDKWuKi1F" {
		t.Error("Incorrect identity returned")
	}
	decodedKey, err := crypto.ConfigDecodeKey(identity.PrivKey)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(decodedKey, keyBytes) {
		t.Error("Incorrect private key returned")
	}
}

func TestIdentityKeyFromSeed(t *testing.T) {
	seed := bip39.NewSeed("mule track design catch stairs remain produce evidence cannon opera hamster burst", "Secret Passphrase")
	key, err := IdentityKeyFromSeed(seed, 4096)
	if err != nil {
		t.Error(err)
	}
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(key, keyBytes) {
		t.Error("Failed to extract correct private key from seed")
	}
}

func TestDeterministicReader_Read(t *testing.T) {
	seed := bip39.NewSeed("mule track design catch stairs remain produce evidence cannon opera hamster burst", "Secret Passphrase")
	reader := &DeterministicReader{Seed: seed, Counter: 0}
	b := make([]byte, 32)
	reader.Read(b)
	if hex.EncodeToString(b) != "5742f7c29729cc98dc62bce3104b5b0a1c6f390625cffef34bf2fd471f79ae3b" {
		t.Error("Deterministic reader returned incorrect bytes")
	}
	if reader.Counter != 1 {
		t.Error("Deterministic Reader failed to increment counter")
	}
	b = make([]byte, 512)
	reader.Read(b)
	if hex.EncodeToString(b) != "ca55fa705e93030eb84041689294c9faae0d265ec1aec75c07bfaa3cb56a5a0a24f4c0cb50805fb46498e5e2adeb0e03f0e20dc3936dca0dc8fed0e0f1b48a1def01e07ba9fbcd54971bd95f11779a67633ce8bab5a809cc7780eea88f4dc42828bea89eb76e43d80b986a394f692dea9f6722f3834c6d680f9c50167339b98f510d246292afb93caa1bd6addd693c59c41a9629d09d4aa5935847c255eeb24c7a0610af5ec4c94abbbab92c27553245c3b41779033f139d396e69e748fb3d462c190fbbfbb85cda69a051c788fa0a3bd940a1c5abd15aa07102b17d2ec4ee0f227d32efa2253d88bef7f58520a23bc6b9144b051f3280f0b21fef3f968d22571d3e889f8514918be8b99e3111f524e5e99bc11584153ce56e76e4ff9de7d9b40a2851f2585a98201a80628919922973b5adff6468be4986ec745ad5dc53708b28862ce77fb52b2868641e5d18e767db8ba66e01e6a05975c882be9bf9c527e52a38186443c03655bdb4943e7b8cb0761a57e4bacd1bf247047863350e9f1cdc011303c3f7f61d9201a12a39b2c63c18693f867c1999a5c071fc8783a8dcf71480841e42559f5ae1f4ea13dfde95940dca0f99da8076b04211e938c97d77225f4de9c96c331019bc3ae40eb7f7d84053924aa3e00f4e3bf452ba7c2dacb56b2d9e7ee3b8582004d4c40a3b72c6741f200e11c61cbe00e29f941b6e90d3335535" {
		t.Error("Deterministic reader returned incorrect bytes")
	}
	if reader.Counter != 2 {
		t.Error("Deterministic Reader failed to increment counter")
	}
}
