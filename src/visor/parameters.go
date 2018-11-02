package visor

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 100000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 25
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
	"2C8moj8UmNcJRSruKB7HZwBjSToYXortSDo",
	"o5Rf5zeTJ7bgzkQ4367SVgu342R5S6pKJ5",
	"fzaCZDB8XXaoxBQ5VUpPiHfRPJPy5uthqu",
	"2R6RrG6fkGoCBM8DvHRnbgB5k2DU2sPfjVX",
	"QWejMvjvRb7wS8KBmP5XyX5umZTkaneKWp",
	"VGgZgQRgf9M4xeXeZym2o3iFE2WVfeVhWq",
	"2MN3bisp57a4QhDG9ESu5Chs6ZuSpi1KDUS",
	"qqcsr43fz9717Dtpsf5j7m6EzC9mMSRiy5",
	"267xzCnnj2GQYLFK2Qb5FxT3nyWropXNXBa",
	"RJmhQps75PhC4fjvG492vrCGxzeY1V3WUT",
	"25YTrZjmk4M5XRM8ZeXUwWnBb7AdQk2BJbf",
	"dRMaWRd6UsFou1UATaJvPzBvqN8nCK1ZTy",
	"mJyxHLzjLviJHas4fz6ZQqrmLzLyMWrv7i",
	"4UAHUyRtcvRexDFQ6om9N3ZnD5NkgZaJ9k",
	"21TeNDPCjh5XyfTriuNzgGxbE9tuUqPGA4e",
	"26TC4Rf4ECSFUdMqLdh2E7tcKFzjV9DjnLp",
	"39RGALcQDtgLB8SLTwCwFEQQPNA6kApndJ",
	"u9bSCdXD7bkrPztpnuzDbosvLPmJudGAKU",
	"2PA58WunafLcj1YVbucfZKXG14Ssf8PMRwP",
	"2ateNuuTaZSx4um6QxfBJmvAqrwNf15C3qY",
	"2PRikJ1roG91agqDVw3Xdu5nKRTaNsQLwnC",
	"sjbb2VsAWw1xVqHeAAoQrRdhZ9Y8Zx5GeP",
	"Ht8JDE1yG3xycr1SZTFRqgdcEnDfpMZ5Hm",
	"2Yf6qk5CLyCheLBSQxXu5AWMTbtBVKcrRnm",
	"2caFMbgxj1QNMVt5p4eH9XFLkJrGNx8f91j",
	"7XcPoumMqJ2K6wqJCT4RG9ArDA3aszhfUw",
	"QvvHC1ijeenkrmvGTT5yL4nfnaN9oWuHre",
	"oGXeCwYNmCn9qHhTtYzaAQUb1ZS6xipAJZ",
	"2ibTjh1Ed5AGMZzpTbhtkAxdqXEUmDbqiHe",
	"uKhFRP4M3gYczya36dPSyy8fHKhJRVQAzG",
	"SFC9BPggFYbNdzALQmZBUDNZ8L67wANvQ7",
	"4kDn6FwG3tS4DjFdrKrksnwgsRuCYRu2Z8",
	"2UyURaDwkGn9Gh2mgPmsb949QEZrXEER6G3",
	"ySNAFStWQTb5ZwB9LYmcU8h6hyNSTdS1iQ",
	"2ML21LsZrcJPkeDYD8L5WPmQeNeuRQWXCAh",
	"2aLmdGKgJk6NiSCoQbsx2HU6du9UjxHMrSa",
	"suWJFJGqPM5YAcnqa8k7XmBEsCX4Y2shyz",
	"FNr2aPkz6owCxUmeRi9ohw2aRwNAGe7R3N",
	"29UiVK8C1HyiocwjSgQQpB4RSdgFhVBPmEx",
	"LK813jjSoPByJ2HzaJnVmAseKWdW755uHX",
	"2k1WuDe1RC9nXZeAUHynTAsaTRsR4t7Eexf",
	"pS24g9zJ2CmMbwe7gXsGd4yurnCZSdgM4V",
	"2i3ToePt3jrLxmU69MUvznT6vYPvYDmsXTE",
	"2YqLRahy9rCYfLmMtWieJFresKww8WtRq58",
	"2WpWXnU4xjYHeyUPuc1pis2Nb1qqZJs7vuN",
	"2mTvv3S6UdFE4m816jgc5qpfXfTYwbAQv1y",
	"2BbQE7GQ44dFB3fkvgRgJkn8Vq1uJdgPPRN",
	"Sk1jqXLbug35XxfZPGht1KsDcq3NW5Uegv",
	"FhPiYhF9RtpbJiq5JJScrj6pRSwNiNuWtx",
	"7yrmXXVzaSveuMsh5o3WoWSrp9ezVTbqc1",
	"2VjK2wcTNypN17h6MX2YYWVttXDemHw6bMf",
	"2XQPD9Hj2koPHwxYHbGBufByu4uxxzoYexT",
	"chwS3Uf5amemAvaK6Y8Ky3H8gmLttBsKCi",
	"2Q9r2MLrD8EEQBotdYtxdgavcZXy6AZ1WWq",
	"4uNXmqGiRvfTswt3a1pHP49DEYWCcgovZB",
	"HpVWHNX5M5MNSwmbgNFjuxw5dyWnpNB8Tt",
	"tBb1gw5V2ekjdExYFnjeafDi99Xeh41Vtb",
	"2Uct6hqVtVrm5BNmgUwmgqTVFah9i9o27HR",
	"2PvAaFpa4FbT1uNgVXhCAfESMJ1H6QDwBSr",
	"1Hv9H8ywNj9wFgqAHwT4sMefLKDF9796kW",
	"2WzW8tXaRg5j5ety4T9SRmBqDi7xz5zpB1p",
	"26JAjJ1EmPbzMePWcLEtveprAwjpZnYMeVR",
	"riezonbg5LGo4Kopyo6CLBsyxmhsGGiFnX",
	"2cMgv3xSUU6tKC5VZai8jUBW5RD4Z9uinZC",
	"Fp5u72TjsWKHxWL8uX7nYSVUbgMVn1BLBJ",
	"XvtXBBNSWJgaCNkXicMfsXq2YeWRPzmxxp",
	"Gkc9EvyvhvHn1aBh7iBhUpvv76X6PY8mj7",
	"2i2yDPrTXSbquf2HPx8Jbu3gcjFS7KjRccL",
	"wyU2hbNHC2hvrzzKRAtpHT8yf3Jdva746o",
	"2Sh6r8URsJS8nMquQ3gKfVzcbiGsvgNibyF",
	"23rmAHhZBPddnsTWWUA9hCqNjVFT5NXcq5q",
	"2AsTbA1iiz8eTeQujyMYJJXXz3epQuedkCx",
	"2REkRteLYULS3SwsXP43pEUF4Q1zHsmyPyo",
	"hS9Ng36UxRfCv8tvpwrZtAPejiYtghUzqC",
	"2mokZP7BXuU44zkB5Vh6ThXdTYE4cznAgeH",
	"2keNLGdDhFtoRJ54AvCcbpT2cXWzaDWVc6i",
	"ReTKfnH752kEykE6xs1CxofykspcCYwABj",
	"peGnVk3ZDwTNs3XmPu1XqcNSFZt99hKFrh",
	"iQu1k8KPQbyFAG83jQaiiorweF6HmUNNPn",
	"xr3PBJGF4u5JhvEkccJZoQXb4pXardE7QB",
	"XZfwqRrVACSoQCr6NLcJbaBFqXwWqYiwSj",
	"4ap82kJYu2oz4FQAzD6H9FnHt9sL4oGn87",
	"2RKgp3M56YRzcNaLA91RUJaeuy7VhaxTJBC",
	"ayToJ5vGc1ckG4BczRDsV1FaQEXiouqBc2",
	"HHWG1HEZR7DXstNir9VjkKbpjz4kACxPrE",
	"2JafZAkqhBWr42yGgZFoM38wRXTYbPSxQPo",
	"H2dbVRtkaLVC9hzNZXQHomG5k5ixWH7Coc",
	"2TiyMRwDzjRSvNuMyqYgiLkJjLGzEXrcX3W",
	"2Kt8uWPdxxXvE18Jhic2LZ1iWqJAacKVMTY",
	"cSthuXCL2iMg84Kj7euJzR4Aoow24tdH7a",
	"nARvMm7pixSoYH3VQRY9RMDX88Jywfk1Qw",
	"ZxfJPGVEFZavHcXrfch8R9PRUtfeSHtFNu",
	"oN7ccZHPcVCRaVQH1Sfqn111TbV599pJL9",
	"8bXe5a6KpkLxjvaPgtKVingC524dA7ABwE",
	"tyewjDtMvb7FrrpWQGG9KTj7KhBeCmX3MN",
	"2X7PH9d1e1iKWuJBYCG9SjwUpGfyQasUNuK",
	"2itibpAsQCRz6FNnNsSwFGvASsyN9LLCVWQ",
	"NfmR9yxQ278DEFtRJA59LMpbHPKwhesi6d",
	"Ez6Dp8c636eS2B5S2mBDbSbBn7Yey4tGgj",
	"2T32wc6bC5pKAuMuXNyVU4ABLPJbf9WUt7G",
}
