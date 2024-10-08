#!/bin/bash

no_order=(1000 1500 2000)
# no_matcher=(1 2 5 10)
# no_send_to=(1 1 2 5)

cd ..

cleanup() {
  echo "Cleaning up..."
  kill $ganache_pid 2>/dev/null
  sync
  sudo sh -c 'echo 3 > /proc/sys/vm/drop_caches'
}

handle_ctrl_c() {
  echo "Ctrl+C pressed. Cleaning up..."
  cleanup
  echo "Exiting..."
  exit 1
}

trap cleanup EXIT
trap handle_ctrl_c SIGINT

for n in "${no_order[@]}"
do
    echo Running $n orders...
    # data_file="./data/real_orders_$n.json"
    ganache -a 200 -m '' -e 99999999999 --chain.chainId 1337 --p 8545 > /dev/null 2>&1 &
    ganache_pid=$!
    sleep 2
    go run . "./data/real_orders_$n"_1".json"
    kill $ganache_pid 2>/dev/null
    wait $ganache_pid 2>/dev/null
    sync
    sudo sh -c 'echo 3 > /proc/sys/vm/drop_caches'
    echo
done
