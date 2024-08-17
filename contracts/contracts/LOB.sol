pragma solidity ^0.8.0;

import "../openzeppelin-contracts/contracts/utils/Strings.sol";

contract LOB {
    using Strings for uint256;

    struct Order {
        bytes16 id;
        uint256 price;
        uint256 amount;
        bool side;
        bytes16 next;
        bytes16 prev;
    }

    bytes16 public bidHeadId;
    bytes16 public askHeadId;
    mapping(bytes16 => Order) bidOrderMapping;
    mapping(bytes16 => Order) askOrderMapping;
    uint256 noOrder;

    event MatchAnOrder();

    constructor() {
        bidHeadId = bytes16(0);
        askHeadId = bytes16(0);
        bidOrderMapping[bidHeadId] = Order(
            bytes16(0),
            0,
            0,
            false,
            bytes16(0),
            bytes16(0)
        );
        askOrderMapping[askHeadId] = Order(
            bytes16(0),
            0,
            0,
            false,
            bytes16(0),
            bytes16(0)
        );
        noOrder = 0;
    }

    function addOrder(
        bytes16 id,
        uint256 price,
        uint256 amount,
        bool side
    ) public {
        if (side) {
            addBidOrder(id, price, amount);
        } else {
            addAskOrder(id, price, amount);
        }
        matching();
    }

    function matching() internal {
        while (
            bidHeadId != bytes16(0) &&
            askHeadId != bytes16(0) &&
            bidOrderMapping[bidHeadId].price >= askOrderMapping[askHeadId].price
        ) {
            uint256 minAmount = bidOrderMapping[bidHeadId].amount;
            if (minAmount > askOrderMapping[askHeadId].amount) {
                minAmount = askOrderMapping[askHeadId].amount;
            }

            bidOrderMapping[bidHeadId].amount -= minAmount;
            askOrderMapping[askHeadId].amount -= minAmount;

            if (bidOrderMapping[bidHeadId].amount == 0) {
                emit MatchAnOrder();
                if (bidOrderMapping[bidHeadId].next != bytes16(0)) {
                    bidHeadId = bidOrderMapping[bidHeadId].next;
                    bidOrderMapping[bidHeadId].prev = bytes16(0);
                } else {
                    bidHeadId = bytes16(0);
                }
            }

            if (askOrderMapping[askHeadId].amount == 0) {
                emit MatchAnOrder();
                if (askOrderMapping[askHeadId].next != bytes16(0)) {
                    askHeadId = askOrderMapping[askHeadId].next;
                    askOrderMapping[askHeadId].prev = bytes16(0);
                } else {
                    askHeadId = bytes16(0);
                }
            }
        }
    }

    function addBidOrder(bytes16 id, uint256 price, uint256 amount) internal {
        Order memory newOrder = Order(
            id,
            price,
            amount,
            true,
            bytes16(0),
            bytes16(0)
        );

        Order memory bidHead = bidOrderMapping[bidHeadId];
        if (bidHead.id == bytes16(0)) {
            bidHeadId = newOrder.id;
            bidOrderMapping[newOrder.id] = newOrder;
        } else if (bidHead.next == bytes16(0)) {
            if (bidHead.price >= price) {
                newOrder.prev = bidHead.id;
                bidHead.next = newOrder.id;
                bidOrderMapping[bidHead.id] = bidHead;
                bidOrderMapping[newOrder.id] = newOrder;
            } else {
                newOrder.next = bidHead.id;
                bidHead.prev = newOrder.id;
                bidHeadId = newOrder.id;
                bidOrderMapping[bidHead.id] = bidHead;
                bidOrderMapping[newOrder.id] = newOrder;
            }
        } else {
            Order memory curr = bidHead;
            while (curr.id != bytes16(0)) {
                if (curr.price < price) {
                    newOrder.prev = curr.prev;
                    newOrder.next = curr.id;
                    if (curr.prev != bytes16(0)) {
                        bidOrderMapping[curr.prev].next = newOrder.id;
                    } else {
                        bidHeadId = newOrder.id;
                    }
                    curr.prev = newOrder.id;
                    bidOrderMapping[curr.id] = curr;
                    bidOrderMapping[newOrder.id] = newOrder;
                    return;
                }
                if (curr.next == bytes16(0)) {
                    break;
                }
                curr = bidOrderMapping[curr.next];
            }
            curr.next = newOrder.id;
            newOrder.prev = curr.id;
            bidOrderMapping[curr.id] = curr;
            bidOrderMapping[newOrder.id] = newOrder;
        }
    }

    function addAskOrder(bytes16 id, uint256 price, uint256 amount) internal {
        Order memory newOrder = Order(
            id,
            price,
            amount,
            true,
            bytes16(0),
            bytes16(0)
        );
        Order memory askHead = askOrderMapping[askHeadId];
        if (askHead.id == bytes16(0)) {
            askHeadId = newOrder.id;
            askOrderMapping[newOrder.id] = newOrder;
        } else if (askHead.next == bytes16(0)) {
            if (askHead.price <= price) {
                newOrder.prev = askHead.id;
                askHead.next = newOrder.id;
                askOrderMapping[askHead.id] = askHead;
                askOrderMapping[newOrder.id] = newOrder;
            } else {
                newOrder.next = askHead.id;
                askHead.prev = newOrder.id;
                askHeadId = newOrder.id;
                askOrderMapping[askHead.id] = askHead;
                askOrderMapping[newOrder.id] = newOrder;
            }
        } else {
            Order memory curr = askHead;
            while (curr.id != bytes16(0)) {
                if (curr.price > price) {
                    newOrder.prev = curr.prev;
                    newOrder.next = curr.id;
                    if (curr.prev != bytes16(0)) {
                        askOrderMapping[curr.prev].next = newOrder.id;
                    } else {
                        askHeadId = newOrder.id;
                    }
                    curr.prev = newOrder.id;
                    askOrderMapping[curr.id] = curr;
                    askOrderMapping[newOrder.id] = newOrder;
                    return;
                }
                if (curr.next == bytes16(0)) {
                    break;
                }
                curr = askOrderMapping[curr.next];
            }
            curr.next = newOrder.id;
            newOrder.prev = curr.id;
            askOrderMapping[curr.id] = curr;
            askOrderMapping[newOrder.id] = newOrder;
        }
    }

    function uintToString(uint256 value) internal pure returns (string memory) {
        return value.toString();
    }

    function getOrderBook() public view returns (string memory) {
        string memory result = "";
        Order memory curr = bidOrderMapping[bidHeadId];
        while (curr.id != bytes16(0)) {
            result = string(
                abi.encodePacked(
                    result,
                    uintToString(curr.price),
                    ",",
                    uintToString(curr.amount),
                    ","
                )
            );
            if (curr.next == bytes16(0)) {
                break;
            }
            curr = bidOrderMapping[curr.next];
        }
        curr = askOrderMapping[askHeadId];
        while (curr.id != bytes16(0)) {
            result = string(
                abi.encodePacked(
                    result,
                    uintToString(curr.price),
                    ",",
                    uintToString(curr.amount),
                    ","
                )
            );
            if (curr.next == bytes16(0)) {
                break;
            }
            curr = askOrderMapping[curr.next];
        }
        return result;
    }
}
