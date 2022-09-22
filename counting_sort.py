def simplifiedCountingSort(array):
    # step 1.a: get max value in the array
    maxValue = max(array)
    # step 1.b: initialize one auxiliary array
    counts = [0]*(1+maxValue) # include index 0

    # step2: count the frequency of every element in the array
    for num in array:
        counts[num] += 1

    # step3: rearrange the array in a sorted manner
    sortedIdx = 0
    for element, frequency in enumerate(counts):
        while frequency > 0:
            array[sortedIdx] = element
            sortedIdx += 1
            frequency -= 1

    return array

def countingSort(array):
    # step 1.a: get max value in the array
    maxValue = max(array)
    # step 1.b: initialize two auxiliary arrays
    counts = [0]*(1+maxValue) # include index 0
    sortedArray = [0]*len(array)

    # step2: count the frequency of every element in the array
    for num in array:
        counts[num] += 1

    # step3: compute the cumulative sum of counts
    for idx in range(1, len(counts)):
        counts[idx] += counts[idx-1]
    
    # step 4: loop backward to build up the sorted array
    for idx in range(len(array)-1, -1, -1):
        currentElement = array[idx]
        # find the element's position
        position = counts[currentElement]
        sortedIdx = position - 1
        # place the element to the sorted array
        sortedArray[sortedIdx] = array[idx]
        # decreament the corresponding value in counts by 1
        counts[currentElement] -= 1

    return sortedArray

if __name__ == "__main__":
    array = [9, 5, 2, 4, 2, 8, 5]
    print("simplified couting sort", simplifiedCountingSort(array))
    print("couting sort",countingSort(array))

"""
output: 
('simplified couting sort', [2, 2, 4, 5, 5, 8, 9])
('couting sort', [2, 2, 4, 5, 5, 8, 9])
"""