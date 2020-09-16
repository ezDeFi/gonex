// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "../interfaces/IToken.sol";

/**
 * Library for token pair exchange.
 *
 * Has no knownledge about what the token does. Any logic deal with stable or
 * volatile nature of the token must put in the contract level.
 */
library dex {
    bytes32 constant ZERO_ID = bytes32(0x0);
    bytes32 constant LAST_ID = bytes32(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF);
    address constant ZERO_ADDRESS = address(0x0);
    uint constant INPUTS_MAX = 2 ** 128;

    // order id works much like HDKey addresses, allows client to recover and identify
    // historical orders using a deterministic chain of reusable indexes.
    function _calcID(
        address maker,
        bytes32 index
    )
        internal
        pure
        returns (bytes32)
    {
        return sha256(abi.encodePacked(maker, index));
    }

    struct Order {
        address maker;
        uint haveAmount;
        uint wantAmount;

        // linked list
        bytes32 prev;
        bytes32 next;
    }
    using dex for Order;

    function exists(Order storage order) internal view returns(bool)
    {
        // including meta orders [0] and [FF..FF]
        return order.maker != ZERO_ADDRESS;
    }

    function isEmpty(Order memory order) internal pure returns(bool)
    {
        return order.haveAmount == 0 || order.wantAmount == 0;
    }

    function betterThan(Order storage o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.wantAmount > p.haveAmount * o.wantAmount;
    }

    // memory version of betterThan
    function m_betterThan(Order memory o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.wantAmount > p.haveAmount * o.wantAmount;
    }

    function fillableBy(Order memory o, Order storage p)
        internal
        view
        returns (bool)
    {
        return o.haveAmount * p.haveAmount >= p.wantAmount * o.wantAmount;
    }


    struct Book {
        IToken haveToken;
        IToken wantToken;
        mapping (bytes32 => Order) orders;
        // bytes32 top;	// the highest priority (lowest sell or highest buy)
        // bytes32 bottom;	// the lowest priority (highest sell or lowest buy)
    }
    using dex for Book;

    function init(
        Book storage book,
        IToken haveToken,
        IToken wantToken
    )
        internal
    {
        book.haveToken = haveToken;
        book.wantToken = wantToken;
        book.orders[ZERO_ID] = Order(address(this), 0, 0, ZERO_ID, LAST_ID); // [0] meta order
        book.orders[LAST_ID] = Order(address(this), 0, 1, ZERO_ID, LAST_ID); // worst order meta
    }

    // read functions
    function topID(
        Book storage book
    )
        internal
        view
        returns (bytes32)
    {
        return book.orders[ZERO_ID].next;
    }

    function bottomID(
        Book storage book
    )
        internal
        view
        returns (bytes32)
    {
        return book.orders[LAST_ID].prev;
    }

    function createOrder(
        address maker,
        bytes32 index,
        uint haveAmount,
        uint wantAmount
    )
        internal
        pure
        returns (bytes32 id, Order memory order)
    {
        // TODO move require check to API
        require(haveAmount > 0 && wantAmount > 0, "zero input");
        require(haveAmount < INPUTS_MAX && wantAmount < INPUTS_MAX, "input over limit");
        id = _calcID(maker, index);
        // Order storage order = book.orders[id];
        // require(!order.exists(), "order index exists");
        // create new order
        order = Order(maker, haveAmount, wantAmount, 0, 0);
        return (id, order);
    }

    // insert [id] as [prev].next
    function insertAfter(
        Book storage book,
        bytes32 id,
        bytes32 prev
    )
        internal
    {
        // prev => [id] => next
        bytes32 next = book.orders[prev].next;
        book.orders[id].prev = prev;
        book.orders[id].next = next;
        book.orders[next].prev = id;
        book.orders[prev].next = id;
    }

    // find the next id (position) to insertAfter
    function find(
        Book storage book,
        Order storage newOrder,
        bytes32 id // [id] => newOrder
    )
        internal
        view
 	    returns (bytes32)
    {
        // [junk] => [0] => order => [FF]
        Order storage order = book.orders[id];
        do {
            order = book.orders[id = order.next];
        } while(!newOrder.betterThan(order));

        // [0] <= order <= [FF]
        do {
            order = book.orders[id = order.prev];
        } while(newOrder.betterThan(order));

        return id;
    }

    // memory version of find
    function m_find(
        Book storage book,
        Order memory newOrder,
        bytes32 id // [id] => newOrder
    )
        internal
        view
 	    returns (bytes32)
    {
        // [junk] => [0] => order => [FF]
        Order storage order = book.orders[id];
        do {
            order = book.orders[id = order.next];
        } while(!newOrder.m_betterThan(order));

        // [0] <= order <= [FF]
        do {
            order = book.orders[id = order.prev];
        } while(newOrder.m_betterThan(order));

        return id;
    }

    // place the new order into its correct position
    function place(
        Book storage book,
        bytes32 id,
        Order memory order,
        bytes32 assistingID
    )
        internal
    {
        require(!book.orders[id].exists(), "order index exists");
        book.orders[id] = order;
        bytes32 prev = book.m_find(order, assistingID);
        book.insertAfter(id, prev);
    }

    // NOTE: this function does not payout nor refund
    // Use payout/refund/fill instead
    function _remove(
        Book storage book,
        bytes32 id
    )
        internal
    {
        // TODO: handle order outside of the book, where next or prev is nil
        Order storage order = book.orders[id];
        // before: prev => order => next
        // after:  prev ==========> next
        book.orders[order.prev].next = order.next;
        book.orders[order.next].prev = order.prev;
        delete book.orders[id];
    }

    function payout(
        Book storage book,
        Order memory order
    )
        internal
    {
        if (order.wantAmount > 0) {
            book.wantToken.transfer(order.maker, order.wantAmount);
        }
        order.wantAmount = 0;
        order.haveAmount = 0;
    }

    function payout(
        Book storage book,
        bytes32 id
    )
        internal
    {
        Order memory order = book.orders[id];
        if (order.wantAmount > 0) {
            book.wantToken.transfer(order.maker, order.wantAmount);
        }
        book._remove(id);
    }

    function refund(
        Book storage book,
        Order memory order
    )
        internal
    {
        if (order.haveAmount > 0) {
            book.haveToken.transfer(order.maker, order.haveAmount);
        }
        order.haveAmount = 0;
        order.wantAmount = 0;
    }

    function refund(
        Book storage book,
        bytes32 id
    )
        internal
    {
        Order memory order = book.orders[id];
        if (order.haveAmount > 0) {
            book.haveToken.transfer(order.maker, order.haveAmount);
        }
        book._remove(id);
    }

    function payoutPartial(
        Book storage book,
        Order memory order,
        uint fillableHave,
        uint fillableWant
    )
        internal
    {
        require (fillableHave <= order.haveAmount, "PP: fillable > have");
        require (fillableWant <= order.wantAmount, "PP: fillable > want");
        order.haveAmount -= fillableHave; // safe
        order.wantAmount -= fillableWant; // safe
        book.wantToken.transfer(order.maker, fillableWant);
    }

    function payoutPartial(
        Book storage book,
        bytes32 id,
        uint fillableHave,
        uint fillableWant
    )
        internal
    {
        Order storage order = book.orders[id];
        require (fillableHave <= order.haveAmount, "PP: fillable > have");
        require (fillableWant <= order.wantAmount, "PP: fillable > want");
        order.haveAmount -= fillableHave; // safe
        order.wantAmount -= fillableWant; // safe
        book.wantToken.transfer(order.maker, fillableWant);
        if (order.isEmpty()) {
            book.refund(id);
        }
    }

    function fill(
        Book storage orderBook,
        Order memory order,
        Book storage redroBook
    )
        internal
    {
        bytes32 redroID = redroBook.topID();

        while (redroID != LAST_ID) {
            if (order.isEmpty()) {
                break;
            }
            Order storage redro = redroBook.orders[redroID];
            if (!order.fillableBy(redro)) {
                break;
            }
            if (order.wantAmount < redro.haveAmount) {
                uint fillable = order.wantAmount * redro.wantAmount / redro.haveAmount;
                // partially payout the redro
                redroBook.payoutPartial(redroID, order.wantAmount, fillable);
                // fully spent the order and stop
                orderBook.payoutPartial(order, fillable, order.wantAmount);
                break;
            }
            // partially payout order
            orderBook.payoutPartial(order, redro.wantAmount, redro.haveAmount);
            // fully payout redro
            redroBook.payout(redroID);
            // next order
            redroID = redroBook.topID();
        }
    }

    function absorb(
        Book storage book,
        bool useHaveAmount,
        uint target
    )
        internal
        returns(uint totalBMT, uint totalAMT)
    {
        bytes32 id = book.topID();
        while(id != LAST_ID && totalAMT < target) {
            Order storage order = book.orders[id];
            uint amt = useHaveAmount ? order.haveAmount : order.wantAmount;
            uint bmt = useHaveAmount ? order.wantAmount : order.haveAmount;
            uint fillableAMT = target - totalAMT; // safe
            if (amt <= fillableAMT) {
                // fill the order
                book.haveToken.dexBurn(order.haveAmount);
                book.wantToken.dexMint(order.wantAmount);
                // emit FullFill(id, order.maker);
                bytes32 next = order.next;
                book.payout(id);
                id = next;
            } else {
                // partial order fill
                bmt = bmt * fillableAMT / amt;
                amt = fillableAMT;
                if (totalBMT + bmt < totalBMT) {
                    // overflow: stop the absorption prematurely
                    return (totalBMT, totalAMT);
                }
                uint fillableHave = useHaveAmount ? amt : bmt;
                uint fillableWant = useHaveAmount ? bmt : amt;
                // fill the partial order
                book.haveToken.dexBurn(fillableHave);
                book.wantToken.dexMint(fillableWant);
                // emit PartialFill(id, order.maker);
                book.payoutPartial(id, fillableHave, fillableWant);
                // extra step to make sure the loop will stop after this
                id = LAST_ID;
            }
            totalBMT += bmt; // safe
            totalAMT += amt; // safe
        }
        // not enough order, return all we have
        return (totalBMT, totalAMT);
    }
}
