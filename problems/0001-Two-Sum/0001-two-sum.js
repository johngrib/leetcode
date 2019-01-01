// https://leetcode.com/problems/two-sum/description/
var twoSum = function(nums, target) {
    var map = {};
    for(var i = 0, len = nums.length; i < len; i++) {
        var v = nums[i];
        var n = target - v;
        if (map[n] != undefined) {
            return [map[n], i];
        }
        map[v] = i;
    }
    return []
};


const assert = require('assert');

describe('twoSum', function () {
    it('basic', function () {
        var result = twoSum([2,7,11,15], 9);
        assert.equal(String(result), String([0,1]));
    })
    it('custom', function () {
        var result = twoSum([3,3], 6);
        assert.equal(String(result), String([0,1]));
    })
})
