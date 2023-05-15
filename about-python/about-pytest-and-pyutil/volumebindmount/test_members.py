import unittest

class Group:
    def __init__(self):
        self.members = []

    def add_member(self, member):
        if member not in self.members:
            self.members.append(member)

    def remove_member(self, member):
        if member in self.members:
            self.members.remove(member)

class TestGroup(unittest.TestCase):
    def test_add_member(self):
        my_group = Group()
        my_group.add_member("Alice")
        self.assertIn("Alice", my_group.members, "Alice should be in my_group.members")

    def test_remove_member(self):
        my_group = Group()
        my_group.add_member("Alice")
        my_group.remove_member("Alice")
        self.assertNotIn("Alice", my_group.members, "Alice should not be in my_group.members after removal")

if __name__ == "__main__":
    unittest.main()
