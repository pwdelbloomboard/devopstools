import unittest

class Person:
    def __init__(self, name):
        self.name = name
        self.friends = []

    def add_friend(self, friend):
        if self is not friend:
            self.friends.append(friend)

class TestPerson(unittest.TestCase):
    def test_same_object(self):
        john = Person("John")
        self.assertIs(john, john, "john should be the same object as itself")

    def test_different_objects(self):
        john = Person("John")
        jane = Person("Jane")
        self.assertIsNot(john, jane, "john and jane should be different objects")

    def test_add_friend(self):
        john = Person("John")
        jane = Person("Jane")
        john.add_friend(jane)
        self.assertIn(jane, john.friends, "jane should be in john's friends list")
        self.assertNotIn(john, jane.friends, "john should not be in jane's friends list")
        self.assertIsNot(john.friends, jane.friends, "john and jane should have different friends lists")

if __name__ == "__main__":
    unittest.main()