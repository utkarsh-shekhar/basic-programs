fun main(arg: Array<String>) {
    val array = mutableListOf(1, 4, 6, 8, 9, 23)
    val searchNumber = 9
    System.out.println("Number $searchNumber sate on index : ${binarySearch(array, searchNumber)}")
}

fun binarySearch(array: MutableList<Int>, searchNumber: Int): Int {
    val center = array.size / 2
    if (array[center] == searchNumber) {
        return center
    }
    return if (array[center] < searchNumber) {
        center + binarySearch(array.subList(center, array.size), searchNumber)
    } else {
        binarySearch(array.subList(0, center), searchNumber)
    }
}