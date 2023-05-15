import unittest

class Person:
    def __init__(self, name):
        self.name = name

class Employee(Person):
    def __init__(self, name, company):
        super().__init__(name)
        self.company = company

class Student(Person):
    def __init__(self, name, school):
        super().__init__(name)
        self.school = school

class TestPerson(unittest.TestCase):
    def test_employee_instance(self):
        john = Employee("John", "Acme Inc")
        self.assertIsInstance(john, Employee, "john should be an instance of Employee")
        self.assertIsInstance(john, Person, "john should also be an instance of Person")

    def test_student_instance(self):
        alice = Student("Alice", "XYZ University")
        self.assertIsInstance(alice, Student, "alice should be an instance of Student")
        self.assertIsInstance(alice, Person, "alice should also be an instance of Person")

    def test_not_employee_instance(self):
        alice = Student("Alice", "XYZ University")
        self.assertNotIsInstance(alice, Employee, "alice should not be an instance of Employee")

    def test_not_student_instance(self):
        john = Employee("John", "Acme Inc")
        self.assertNotIsInstance(john, Student, "john should not be an instance of Student")

if __name__ == "__main__":
    unittest.main()