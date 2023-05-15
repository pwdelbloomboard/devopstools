import unittest

def multiply(a, b):
    return a * b

def divide(a, b):
    if b == 0:
        raise ValueError("Division by zero is not allowed.")
    return a / b

class TestExample(unittest.TestCase):
    def test_sum(self):
        self.assertEqual(sum([1, 2, 3]), 6, "Should be 6")


class TestMath(unittest.TestCase):
    def test_multiply(self):
        self.assertEqual(multiply(2,3),6, "Should be 6")
    
    def test_divide(self):
        with self.assertRaises(ValueError, msg="Division by zero is not allowed, bro!"):
            divide(2, 0)

    def test_divide_negative(self):
        self.assertEqual(divide(-4, 2),-2, "Should be -2")