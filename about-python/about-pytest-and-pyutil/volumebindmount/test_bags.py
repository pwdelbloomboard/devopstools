import unittest

class Bag:
    def __init__(self):
        self.items = []

    def add_item(self, item):
        self.items.append(item)

    def remove_item(self, item):
        if item in self.items:
            self.items.remove(item)
            return item
        return None
    
class TestBag(unittest.TestCase):
    def test_add_item(self):
        my_bag = Bag()
        my_bag.add_item("apple")
        self.assertIsNotNone(my_bag.items, "my_bag.items should not be None")
        self.assertIn("apple", my_bag.items, "apple should be in my_bag.items")

    def test_remove_item(self):
        my_bag = Bag()
        my_bag.add_item("apple")
        removed_item = my_bag.remove_item("apple")
        self.assertIsNotNone(removed_item, "removed_item should not be None")
        self.assertEqual(removed_item, "apple", "removed_item should be 'apple'")

    def test_remove_nonexistent_item(self):
        my_bag = Bag()
        removed_item = my_bag.remove_item("apple")
        self.assertIsNone(removed_item, "removed_item should be None when item is not in the bag")

if __name__ == "__main__":
    unittest.main()