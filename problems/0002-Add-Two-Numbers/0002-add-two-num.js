// https://leetcode.com/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 */
function ListNode(val) {
    this.val = val;
    this.next = null;
}

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {

    var head = new ListNode(0)
    var pointer = head
    var next = 0

    while (true) {

        pointer.val = l1.val + l2.val + next
        next = parseInt(pointer.val / 10, 10)
        pointer.val %= 10

        l1 = l1.next;
        l2 = l2.next;

        if (l1 == undefined && l2 == undefined) {
            break;
        }
        if (l1 == undefined) {
            l1 = new ListNode(0)
        }
        if (l2 == undefined) {
            l2 = new ListNode(0)
        }

        pointer.next = new ListNode(0)
        pointer = pointer.next
    }

    if (next != 0) {
        pointer.next = new ListNode(next)
    }

    return head;
};

const assert = require('assert');

describe('addTwoNumbers', function() {
    context('with a basic case', function() {
        it('returns the correct result', function() {

            var l1 = new ListNode(2);
            var l2 = new ListNode(4);
            var l3 = new ListNode(3);
            l1.next = l2;
            l2.next = l3;

            var r1 = new ListNode(5);
            var r2 = new ListNode(6);
            var r3 = new ListNode(4);
            r1.next = r2;
            r2.next = r3;

            var result = addTwoNumbers(l1, r1);
            assert.equal(result.val, 7);
            assert.equal(result.next.val, 0);
            assert.equal(result.next.next.val, 8);
        })
    })

    context('with a custom case', function() {
        it('returns the correct result', function() {

            var l1 = new ListNode(5);
            var r1 = new ListNode(5);

            var result = addTwoNumbers(l1, r1);
            assert.equal(result.val, 0);
            assert.equal(result.next.val, 1);
        })
    })
})
