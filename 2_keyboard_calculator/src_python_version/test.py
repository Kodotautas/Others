#run automated tests

import unittest
from main import main
import time

class TestMain(unittest.TestCase):
    def test_main(self):
        self.assertEqual(main(), 'Interval must be a number and greater than 0.')
    
    def test_main2(self):
        self.assertEqual(main(), 'Session must be a number and greater than 0.')
    
    def test_main3(self):
        self.assertEqual(main(), 'Session time must be greater than interval time. Try again.')

if __name__ == '__main__':
    unittest.main()