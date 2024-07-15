// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract LOB {
    struct Order {
        bytes16 orderID;
        uint256 price;
        uint256 amount;
        bool side;
    }

    event MyMatchEvent(bytes16, uint256);
    
    /**
     * MATCHING TIME
     */
    event MatchTimestamp(uint256);
    mapping(bytes16 => uint256) OrderTimestamp;
    uint256 noOrders;
    //

    Order[] bidOrders;
    Order[] askOrders;

    constructor() {
        noOrders = 0;
    }

    function addOrder(Order memory newOrder) public {
        /**
        * MATCHING TIME
        */
        noOrders += 1;
        OrderTimestamp[newOrder.orderID] = noOrders;
        // 

        if (newOrder.orderID == bytes16(0)) {
            emit MyMatchEvent(newOrder.orderID, newOrder.amount);
        }

        if (newOrder.side) {
            addBidOrder(newOrder);
        } else {
            addAskOrder(newOrder);
        }
        matching();
    }

    function matching() internal {
        while (canMatch()) {
            uint256 minAmount = bidOrders[0].amount;
            if (minAmount > askOrders[0].amount) {
                minAmount = askOrders[0].amount;
            }

            emit MyMatchEvent(bidOrders[0].orderID, minAmount);
            // emit MyMatchEvent(askOrders[0].orderID, minAmount);
            
            /**
            * MATCHING TIME
            */
            emit MatchTimestamp(
                // bidOrders[0].orderID,
                noOrders - OrderTimestamp[bidOrders[0].orderID]
            );
            emit MatchTimestamp(
                // askOrders[0].orderID,
                noOrders - OrderTimestamp[askOrders[0].orderID]
            );
            //

            bidOrders[0].amount -= minAmount;
            askOrders[0].amount -= minAmount;

            if (bidOrders[0].amount == 0) {
                for (uint i = 0; i < bidOrders.length - 1; i++) {
                    bidOrders[i] = bidOrders[i + 1];
                }
                bidOrders.pop();
            }

            if (askOrders[0].amount == 0) {
                for (uint i = 0; i < askOrders.length - 1; i++) {
                    askOrders[i] = askOrders[i + 1];
                }
                askOrders.pop();
            }
        }
    }

    function canMatch() internal view returns (bool) {
        if (bidOrders.length == 0 || askOrders.length == 0) {
            return false;
        }

        return bidOrders[0].price >= askOrders[0].price;
    }

    function addBidOrder(Order memory newOrder) internal {
        if (bidOrders.length == 0) {
            bidOrders.push(newOrder);
        } else if (bidOrders.length == 1) {
            if (bidOrders[0].price < newOrder.price) {
                bidOrders.push(bidOrders[0]);
                bidOrders[0] = newOrder;
            } else {
                bidOrders.push(newOrder);
            }
        } else {
            for (uint8 i = 0; i < bidOrders.length; i++) {
                if (bidOrders[i].price < newOrder.price) {
                    bidOrders.push(newOrder);
                    for (uint8 j = uint8(bidOrders.length) - 1; j > i; j--) {
                        bidOrders[j] = bidOrders[j - 1];
                    }
                    bidOrders[i] = newOrder;
                    return;
                }
            }
            bidOrders.push(newOrder);
        }
    }

    function addAskOrder(Order memory newOrder) internal {
        if (askOrders.length == 0) {
            askOrders.push(newOrder);
        } else if (askOrders.length == 1) {
            if (askOrders[0].price > newOrder.price) {
                askOrders.push(askOrders[0]);
                askOrders[0] = newOrder;
            } else {
                askOrders.push(newOrder);
            }
        } else {
            for (uint8 i = 0; i < askOrders.length; i++) {
                if (askOrders[i].price > newOrder.price) {
                    askOrders.push(newOrder);
                    for (uint8 j = uint8(askOrders.length) - 1; j > i; j--) {
                        askOrders[j] = askOrders[j - 1];
                    }
                    bidOrders[i] = newOrder;
                    return;
                }
            }
            bidOrders.push(newOrder);
        }
    }
}
