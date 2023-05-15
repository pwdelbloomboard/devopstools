import pytest

def multiply(a, b):
    return a * b

def divide(a, b):
    if b == 0:
        raise ValueError("Division by zero is not allowed.")
    return a/b

def test_sum():
    assert sum([1, 2, 3]) == 6, "Should be 6"

@pytest.mark.parametrize("a, b, expected", [
    (2, 3, 6),
    (4, 5, 20),
    (6, 7, 42),
])
def test_multiply(a, b, expected):
    assert multiply(a, b) == expected, f"Should be {expected}"

@pytest.mark.parametrize("a, b, expected", [
    (9, 3, 3),
    (15, 5, 3),
    (16, 4, 4),
])
def test_divide(a, b, expected):
    assert divide(a, b) == expected, f"Should be {expected}"

def test_divide_zero():
    with pytest.raises(ValueError, match="Division by zero is not allowed."):
        divide(2, 0)

def test_divide_negative():
    assert divide(-4, 2) == -2, "Should be -2"


