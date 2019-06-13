// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://286d9b5690f0c2a322f5bf31775fb06f2992d3de001d9a9ab62513b813d48f5607a959a8e499f37153f2e78c74c751f1756e56588d02d032c1c4a92c002229ba@35.187.233.103:33333",
	"enode://f3a6df4d7a1c1566f54deb0449770a88403d03313911e08af88d312011de7234d4a6231073678bacbb93df036d5f48e5c419cf6a58cda7fff0a04d6786175c37@139.180.137.154:33333",
	"enode://866fbc2c7dd95adc8db8ceb91442ea276a788c92cb3b755b4adac52d2012343e49a35e6194224960e0257cb466be657354a680dfef78b227090ea1a417aa5bee@52.74.133.33:33333",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Dccs test network.
var TestnetBootnodes = []string{
	"enode://ee47802e44de30d3073e4646657bb9367611235546faa11fd40a6ed2983e3a40983e035d66d84a1525726e329da9eac4dd139a009fd8dec27b98f58bd6985528@108.61.148.72:30301",
	"enode://5deadedcda84af8f40b9fdea66d93d56d408ab674e0c4939bdd26750b1820bdd78bbac7cd5de9e7a44c609d7be7c1cb08842d5adc3df7fc3bbcf3fda9c378ffd@199.247.15.240:30301",
	"enode://f8572a8177f20972139ada1d41344ff62481fbde6abbf6cb2ffcde42139eb9c67248dd3bd45c5672238de41b2ff8fbf746bb6ebe46a6d3e6a5a83fcebb0fb797@45.77.22.149:30301",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",

	// "enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	// "enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	// "enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DccsBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Dccs network.
var DccsBootnodes = []string{
	"enode://286d9b5690f0c2a322f5bf31775fb06f2992d3de001d9a9ab62513b813d48f5607a959a8e499f37153f2e78c74c751f1756e56588d02d032c1c4a92c002229ba@35.187.233.103:33333",
	"enode://f3a6df4d7a1c1566f54deb0449770a88403d03313911e08af88d312011de7234d4a6231073678bacbb93df036d5f48e5c419cf6a58cda7fff0a04d6786175c37@139.180.137.154:33333",
	"enode://866fbc2c7dd95adc8db8ceb91442ea276a788c92cb3b755b4adac52d2012343e49a35e6194224960e0257cb466be657354a680dfef78b227090ea1a417aa5bee@52.74.133.33:33333",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstrem bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://f21f3c81e728ed5dd48805280ce98fb789906a9c425446da3d993206c817b012abfe2179e53d35c3e5cdea74a870152f81a2ca43a872db495df447157276c418@13.228.68.50:33333",
	"enode://1ad838bc92f7ce34e2be3d3655fb61435553721166cb83a84e485c49c74be9c79d7aa590fb6f63c93818aa43131afa306d208ecb67419848dc9eb6a4f0858ed0@54.38.241.133:33333",
	"enode://18b96db04dd2d7bc46bf331c88781299d6209738e16629cf9f34388f5f436370cc87d888611b82a22244eb59ed358dbcdf88fc9cb65a5602f5516c90ab72033f@159.89.229.137:33333",
	"enode://0e81ffc72ba4e2ed30789c4c78347d520fb35b8e46c108e012bd906a4f44b3e453e48840b05c2b48ad93f4d7c21dbff5ba7affc9fbff7b1fbc0711d37bfaa911@139.99.96.128:33333",
	"enode://1b6a4140194a30ecdfc9cfd8ab33f96f6ca2fee737a581fe76f05e985e859b7d9a21c5d5c76ea413956f6b3970db9aa22fa8d76e020eb9fdbb3f49a496a4d84e@167.99.103.67:33333",
}
