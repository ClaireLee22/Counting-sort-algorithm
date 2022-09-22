def simplifiedCountingSort(array):
    maxValue = max(array)
    minValue = min(array)
    # calculate the offset
    offset = 0 - minValue
    counts = [0]*(1 + maxValue + offset)

    for num in array:
        # add offset to to counts index
        counts[num+offset] += 1

    sortedIdx = 0
    for element, frequency in enumerate(counts):
        while frequency > 0:
            # subtract offset to to counts index to get current value
            array[sortedIdx] = element-offset
            sortedIdx += 1
            frequency -= 1

    return array

def countingSort(array):
    maxValue = max(array)
    minValue = min(array)
    # calculate the offset
    offset = 0 - minValue
    counts = [0]*(1 + maxValue + offset)
    sortedArray = [0]*len(array)

    for num in array:
        # add offset to to counts index
        counts[num+offset] += 1

    for idx in range(1, len(counts)):
        counts[idx] += counts[idx-1]
    
    for idx in range(len(array)-1, -1, -1):
        currentElement = array[idx]
        # add offset to to counts index
        position = counts[currentElement+offset]
        sortedIdx = position - 1
        sortedArray[sortedIdx] = array[idx]
        # add offset to to counts index
        counts[currentElement+offset] -= 1

    return sortedArray

if __name__ == "__main__":
    array = [9, 5, -2, 4, 2, 8, 5]
    print("simplified couting sort", simplifiedCountingSort(array))
    print("couting sort",countingSort(array))

"""
output: 
('simplified couting sort', [-2, 2, 4, 5, 5, 8, 9])
('couting sort', [-2, 2, 4, 5, 5, 8, 9])
"""
