// https://leetcode.com/problems/two-sum/description/
const twoSum = function(nums, target) {
    const map = {};

    for (let i = 0, len = nums.length; i < len; i++) {
        const v = nums[i];
        const n = target - v;

        if (map[n] != undefined) {
            return [map[n], i];
        }
        map[v] = i;
    }
    return [];
};


const assert = require('assert');

describe('twoSum', function() {
    context('with a basic case', function() {
        const given = [2, 7, 11, 15];
        const target = 9;

        it('returns indices of 2 and 7', function() {
            const result = twoSum(given, target);

            assert.equal(String(result), String([0, 1]));
        })
    })
    context('with a custom case', function() {
        const given = [3, 3];
        const target = 6;

        it('returns indices of 3 and 3', function() {
            const result = twoSum(given, target);

            assert.equal(String(result), String([0, 1]));
        })
    })
})
