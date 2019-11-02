#/bin/bash

#rm -rf /Users/tomkk/.lsd
#rm -rf /Users/tomkk/.lscli
#lsd init kknode --chain-id lotterychain
#lscli keys add jack > jack.account
#lscli keys add alice > alice.account

lsd add-genesis-account $(lscli keys show jack -a) 1000lotterytoken,100000000stake
lsd add-genesis-account $(lscli keys show alice -a) 1000lottertoken,100000000stake

lscli config chain-id lotterychain
lscli config output json
lscli config indent true
lscli config trust-node true

#lsd gentx --name jack

#lsd collect-gentxs
#lsd validate-genesis