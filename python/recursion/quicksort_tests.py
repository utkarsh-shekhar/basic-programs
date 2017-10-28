import unittest
from random import randint
from quicksort import quicksort

class QuickSortUnitTest(unittest.TestCase):
    
    def setUp(self):
        pass
    
    def test_empty_list(self):
        a, b = [], []
        quicksort(a)
        self.assertEqual(a, b)
        
    def test_one_item_list(self):
        a, b = [1], [1]
        quicksort(a)
        self.assertEqual(a, b)
    
    def test_sorted_list(self):
        a, b = [1,2,3,4,5], [1,2,3,4,5]
        quicksort(a)
        self.assertEqual(a, b)
    
    def test_partially_sorted(self):
        a, b = [5,12,4,18,0,1,2,3], [5,12,4,18,0,1,2,3]
        quicksort(a)
        #sort a copy using python inbuilt sort method
        b.sort()
        self.assertEqual(a, b)
        
    def test_reverse_sorted_list(self):
        a, b = [5,3,1,0], [5,3,1,0]
        quicksort(a)
        b.sort()
        self.assertEqual(a,b)

    def test_random_list(self):
        random_sz = randint(1, 2000)
        a = [randint(-2000,5000) for i in range(0, random_sz)]
        b = a[:]
        quicksort(a)
        b.sort()
        self.assertEqual(a,b) 

if __name__ == '__main__':
    unittest.main()