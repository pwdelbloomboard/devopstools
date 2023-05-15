import unittest

class Vehicle:
    pass

class Car(Vehicle):
    pass

class Bike(Vehicle):
    pass

class TestVehicles(unittest.TestCase):
    def test_car_is_instance(self):
        my_car = Car()
        self.assertIsInstance(my_car, Car, "my_car should be an instance of Car")
        self.assertIsInstance(my_car, Vehicle, "my_car should be an instance of Vehicle (or a subclass)")

    def test_bike_is_instance(self):
        my_bike = Bike()
        self.assertIsInstance(my_bike, Bike, "my_bike should be an instance of Bike")
        self.assertIsInstance(my_bike, Vehicle, "my_bike should be an instance of Vehicle (or a subclass)")

    def test_car_not_instance(self):
        my_car = Car()
        self.assertNotIsInstance(my_car, Bike, "my_car should not be an instance of Bike")

    def test_bike_not_instance(self):
        my_bike = Bike()
        self.assertNotIsInstance(my_bike, Car, "my_bike should not be an instance of Car")

if __name__ == "__main__":
    unittest.main()