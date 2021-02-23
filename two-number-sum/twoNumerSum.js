/**
 * Problem Statement:
 * 
 *  Given an array of distinct integers, find the pair of integers whose sum is equal to the target number, if any
 * 
 * Arguments:
 * 
 *  1. Array of distinct integers
 *  2. Target sum number
 * 
 */

/**
 * input: node twoNumerSum.js test
 * output: ['/home/brijesh/.nvm/versions/node/v14.16.0/bin/node', '/home/brijesh/Desktop/projects/algos/two-number-sum/twoNumerSum.js', 'test']
 */
const arguments = process.argv;


/**
 * O(n^2) time | O(1) space
 * @param {Array} numbersArray 
 * @param {Number} targetSum 
 */
// const twoNumerSum = (numbersArray, targetSum) => {

//     for (let i = 0; i < numbersArray.length; i++) {

//         const firstNumber = numbersArray[i];

//         for (let j = 0; j < numbersArray.length; j++) {

//             const secondNumber = numbersArray[j];

//             if (firstNumber + secondNumber === targetSum) {

//                 return [firstNumber, secondNumber];

//             }

//         }

//     }

//     return [];

// };


/**
 * O(n) time | O(n) space
 * @param {Array} numbersArray 
 * @param {Number} targetSum 
 */
// const twoNumerSum = (numbersArray, targetSum) => {

//     const numbersMap = {};

//     for (const number of numbersArray) {

//         const match = targetSum - number;

//         if (numbersMap[match]) {

//             return [match, number];

//         } else {

//             numbersMap[number] = 'true';

//         }

//     }

//     return [];

// };


/**
 * O(nlogn) time | O(1) space
 * @param {Array} numbersArray 
 * @param {Number} targetSum 
 */
const twoNumerSum = (numbersArray, targetSum) => {

    numbersArray.sort();

    let left = 0, right = numbersArray.length - 1;

    while (left < right) {

        const currentSum = numbersArray[left] + numbersArray[right];

        if (currentSum === targetSum) {

            return [numbersArray[left], numbersArray[right]];

        } else if (currentSum < targetSum) {

            left += 1;

        } else {

            right -= 1;

        }

    }

    return [];

};

/**
 * input: node twoNumerSum.js [3,2,1,7] 10
 * output: [ 3, 7 ]
 */
if (require.main === module) {

    const [numbersArray, targetSum] = [JSON.parse(arguments[2]), Number(arguments[3])];

    const combinations = twoNumerSum(numbersArray, targetSum);

    console.log(combinations);

}