## About pytest and pyutil

* Unit testing in general is a way to accelerate code development by creating covered ways to test code as it is being written without having to run the code on a particular environment.
* While usability testing is recommended upon the first iterations of creating software, or at whatever lifecycles are necessary to make sure the software is useful from a human perspective, unit testing is a way of accellerating development and ensuring functionality, similar to running on local prior to running on cloud, and can be done in parallel, taking away the necessity of doing usability testing which incorporates unit testing, and increasing reliability over pure human testing, which might miss something.
* While pytest is a custom library that allows for lighter weight syntax, unitttest is built into python.

## About pytest

* The basics of how pytest works is documented in test_example.py

* @pytest.mark.skip is a decorator provided by the pytest testing framework, which allows you to skip a specific test function when running your test suite. It can be useful when you have a test that is not yet implemented or temporarily broken, and you don't want it to be executed along with the other tests.

## About unittest

* unittest is a part of the fundamental python library, which is why it is not designated as a specific package on PyPi, or rather, PyPi just points to version 0.0.

* examples are given at unit_example.py

* to run a python file that includes unittests, you use the following type of line:

```
python -m unittest unit_example.py
```

### Pytest Expanding Upon self.assertFalse() or self.assertTrue() for Generalizeable Purposes

* The self.assertFalse() method is part of the unittest framework and is used to verify if a given condition is False. In the provided code snippet, it checks if the result of the command.run.called_with() method is False. If the condition is False, the test will pass, and if it's True, the test will fail.

So if you had a command like the following within your test suite, which could be from a custom utility called_with() that checks whether the run function was called with the arguments f'git clone {name} {dest}'.:

```
self.assertFalse(command.run.called_with(f'git clone {name} {dest}'))  # type: ignore
```

* So this could allow a generalizeable way of checking a string with variables {name} and {dest} dynamically rendered with f' string notation:

```
 f'git clone {name} {dest}'.
```


### Other Fundamental Methods Within unittest

This exercise has gone through some of the simple ways of running unittests, the fundamental methods. Other methods include:

* self.assertIs(a, b): Checks if a and b are the same object.
* self.assertIsNot(a, b): Checks if a and b are not the same object.
* self.assertIsNone(x): Checks if x is None.
* self.assertIsNotNone(x): Checks if x is not None.
* self.assertIn(a, b): Checks if a is a member of the container b.
* self.assertNotIn(a, b): Checks if a is not a member of the container b.
* self.assertIsInstance(a, b): Checks if a is an instance of class b or a subclass thereof.
* self.assertNotIsInstance(a, b): Checks if a is not an instance of class b or a subclass thereof.
* self.assertGreater(a, b): Checks if a is greater than b.
* self.assertGreaterEqual(a, b): Checks if a is greater than or equal to b.
* self.assertLess(a, b): Checks if a is less than b.
* self.assertLessEqual(a, b): Checks if a is less than or equal to b.
* self.assertListEqual(a, b): Checks if lists a and b are equal.
* self.assertTupleEqual(a, b): Checks if tuples a and b are equal.
* self.assertDictEqual(a, b): Checks if dictionaries a and b are equal.
* self.assertSetEqual(a, b): Checks if sets a and b are equal.
* self.assertSequenceEqual(a, b): Checks if sequences a and b are equal.

These methods can be combined to create a wide range of test cases. We will give examples below for the unittest library.

### Equivalent Methods in PyTest

* All of the methods in unittest should be available in pytest as well.

self.assertIs(a, b): assert a is b
self.assertIsNot(a, b): assert a is not b
self.assertIsNone(x): assert x is None
self.assertIsNotNone(x): assert x is not None
self.assertIn(a, b): assert a in b
self.assertNotIn(a, b): assert a not in b
self.assertIsInstance(a, b): assert isinstance(a, b)
self.assertNotIsInstance(a, b): assert not isinstance(a, b)
self.assertGreater(a, b): assert a > b
self.assertGreaterEqual(a, b): assert a >= b
self.assertLess(a, b): assert a < b
self.assertLessEqual(a, b): assert a <= b
self.assertListEqual(a, b): assert a == b
self.assertTupleEqual(a, b): assert a == b
self.assertDictEqual(a, b): assert a == b
self.assertSetEqual(a, b): assert a == b
self.assertSequenceEqual(a, b): assert a == b


### Object Testing

* This is demonstated within test_people.py. We can set up a scenario in which a class person has two properties, a name and friends.
* We can then, given these properties, check for how an object fits within a venn diagram, using logical properties of inclusivity and exclusivity.

```
self.name = name
self.friends = []
```

* So for example, after we do, "john.add_friend(jane)", john is not necessarily in jane's friend list, but jane is on john's friend list. Two friend lists of any two given people would ever be the same because one's own friends list does not contain oneself.

### Null Testing

* This is demonstrated within test_bags.py
* Thinking about objects and objects that hold other objects in physical reality, one could have a bag and an apple. If an apple is put into a bag, then the status of that bag is not null. If the apple is then taken out of the bag, then the status of the bag could be accepted as null.

So fundamentally, asserting none would be accomplished by removing an apple.

```
    def test_remove_nonexistent_item(self):
        my_bag = Bag()
        removed_item = my_bag.remove_item("apple")
        self.assertIsNone(removed_item, "removed_item should be None when item is not in the bag")
```

### Group Member Testing

* This is demonstated in test_members.py
* The concept behind membership is different than null testing, because it has to do with a specific object, not merely a type of object. E.g. an apple can be thought of as a type of object and all apples could be considered annonymous, the same, effectively equivalent to one another, whereas people can be considered individuals.

* self.assertIn(a, b): Checks if a is a member of the container b.
* self.assertNotIn(a, b): Checks if a is not a member of the container b.

So while Null Testing and Group Member Testing both involve containers or placing objects into other objects, Group Member testing has to do with the individuality of the objects in question.

### Instance Testing

* A demonstration of Instance Testing is given in test_vehicles.py and test_school.py.
* In general Instances can be thought of as a way of structuring data conceptually.
* Instances assume that you use Classes. A class of something could be a Person, which has a name. However a class could also be an Employee or a Student, which has an associated Company or School respectively, and could be dependant upon or, "inherent from," a Person class, meaning that it would be necessary to establish the concept of a Person to have an Employee, e.g.:

```
class Person:
    def __init__(self, name):
        self.name = name

class Employee(Person):
    def __init__(self, name, company):
        super().__init__(name)
        self.company = company
```
* So basically, there are hierarchical layers of concepts that inherent from other concepts.
* From a testing perspective, you could test that a particular object is or is not an *Instance* of various levels of Classes, even if a Class inherits from another Class, e.g.

```
    def test_employee_instance(self):
        john = Employee("John", "Acme Inc")
        self.assertIsInstance(john, Employee, "john should be an instance of Employee")
        self.assertIsInstance(john, Person, "john should also be an instance of Person")
```

### Directory Organization

A common convention in Python projects is to separate the test files from the application code. This structure helps maintain a clean and organized codebase, making it easier to locate and manage test files.

Here's a brief explanation of the organization:

* The main_application folder contains the entire application, including tests and application code.
* The tests folder holds all the test-related files and is separate from the application code.
* The mock_data folder within the tests folder is used to store mock data, like JSON files or any other data formats used in testing.
* Test files are named with the test_ prefix, like test_commands.py, test_utils.py, and test_models.py. This naming convention helps in identifying test files quickly and is required by the test discovery mechanism of testing frameworks like unittest and pytest.
* The app folder contains the actual application code, organized into different services or modules (e.g., service1 and service2).
* This organization ensures a clean separation between test and application code, making it easier to manage and maintain the codebase. It also complies with the naming conventions and test discovery requirements of popular testing frameworks in Python, such as unittest and pytest.

## About Tox

* tox is a powerful Python testing tool that makes it easy to automate testing across multiple Python environments, including different Python versions and interpreter implementations. It is often used in open-source projects to ensure compatibility and consistent behavior across different platforms.
* Tox can be run as a command-line tool, with the simplest implementation of tox being:

```
[tox]
envlist = py39

[testenv]
deps = pytest
commands = pytest test_example.py
```

* and then running tox results in:

```
# tox
py39: install_deps> python -I -m pip install pytest
py39: commands[0]> pytest test_example.py
========= test session starts =============================
platform linux -- Python 3.9.16, pytest-7.3.1, pluggy-1.0.0
cachedir: .tox/py39/.pytest_cache
rootdir: /home/volumebindmount
collected 9 items
test_example.py .........
=========== 9 passed in 0.05s ===========
  py39: OK (12.70=setup[12.14]+cmd[0.56] seconds)
  congratulations :) (12.80 seconds)
```

### Tox Code Coverage

* We can add `pytest-cov` as wellas the --cov=. flag into the tox file to get code coverage.

```
[tox]
envlist = py39

[testenv]
deps =
    pytest
    pytest-cov
commands = pytest --cov=. test_example.py
```

So when this is added into the tox file, we get results on the code coverage for every file in the directory:

```
---------- coverage: platform linux, python 3.9.16-final-0 -----------
Name               Stmts   Miss  Cover
--------------------------------------
test_bags.py          29     29     0%
test_example.py       20      0   100%
test_members.py       22     22     0%
test_people.py        25     25     0%
test_school.py        29     29     0%
test_vehicles.py      24     24     0%
unit_example.py       18     18     0%
--------------------------------------
TOTAL                167    147    12%
```

* Above we have test_example.py which shows 100% code coverage because we have written tests, at least in pytest for only that file.
* To interpret the above, the following definitions are used:
    * Stmts: The total number of statements in the file.
    * Miss: The number of statements that were not executed during the tests.
    * Cover: The percentage of statements executed during the tests, calculated as (Stmts - Miss) / Stmts * 100.

#### Testing More Files

* Note that you have to pass the files in as arguments that you want to test within tox.  So if you wanted to test test_boto3.py as well, you would add it into the tox as follows:

```
[tox]
envlist = py39

[testenv]
deps =
    -rrequirements.txt

commands = \
    pytest --cov=. test_example.py test_boto3.py
    python -m unittest discover  # This command will run unittest tests
```

...and if you just wanted to test everything within the same tox directory, you would just run `pytest --cov=. ./`

### unittests within tox.ini

* To run unittest tests using tox, you need to specify the test command in the tox.ini file. You can use the commands key under a specific test environment section to define the command that will be executed to run the tests.

* Here's an example of how to configure a tox.ini file to run unittest tests:

```
[tox]
envlist = py39  # You can list more Python versions here

[testenv]
deps =
    # Add any dependencies required for testing here

commands =
    python -m unittest discover  # This command will run unittest tests
```

* In this example, the commands key is set to python -m unittest discover. This command runs the unittest test discovery, which searches for and runs all test cases in your project.

### Breaking Down Pytest Flags within Tox

* pytest is the command to run your tests using the pytest testing framework. So let's say there were a number of arguments being fed in, as follows:

```
--hypothesis-show-statistics is an option for Hypothesis, a property-based testing framework. If you're using Hypothesis in your tests, this will print a statistics report at the end of the test run.
--verbose makes pytest's output more verbose, meaning it will provide more details about the tests that are being run.
--capture=sys changes how pytest captures output. With sys, writes to sys.stdout and sys.stderr will be captured.
--cov app runs coverage analysis on your code with the pytest-cov plugin, specifically for the "app" module or package.
--cov-branch adds branch coverage to the coverage analysis. This means it will check if both the True and False branches of if/else statements are covered by your tests.
--cov-report html generates an HTML report of the coverage analysis. This will be a detailed report that you can view in a web browser.
--cov-report term outputs a summary of the coverage analysis to the terminal when the tests are done.
--junitxml=report.xml generates a junit xml report. This is a standard format that can be used by other tools to understand the test results.
```

### Tox Sylte Constraints Settings Related to Flake8 or other Linting Tools

* This is related to code complexity and code style constraints. They are used to define limits on various aspects of your code to maintain readability and maintainability. 

* max-local-variables = X: Sets the maximum number of local variables allowed in a function. Having too many local variables can make a function difficult to understand and maintain.

* max-line-complexity = X: Sets the maximum line complexity allowed. This is usually measured using metrics like McCabe complexity or cyclomatic complexity, which quantify the number of independent code paths in a line. A high complexity value indicates that the line may be hard to understand and maintain.

* max-expressions = X: Sets the maximum number of expressions allowed in a single function. Functions with too many expressions can become difficult to understand and maintain.

* max-arguments = X: Sets the maximum number of arguments allowed for functions or methods. Functions with too many arguments can be challenging to understand, maintain, and use.

* max-module-members = X: Sets the maximum number of methods or classes allowed per module. This helps to limit the size and complexity of individual modules, promoting modularity and maintainability.

* max-cognitive-score = X: Sets the maximum cognitive complexity score allowed for functions. Cognitive complexity is a measure of how difficult it is for humans to understand a piece of code. It takes into account factors such as nesting, branching, and the number of conditions in a function. A high cognitive complexity score indicates that the function may be hard to understand and maintain.


### Other Things To Possibly Add Within tox.ini:

* exclude = .git/*, .tox/*, __pycache__/*, .eggs/*, sphinx-docs/*, .direnv/*, .history

### API Testing, Fake Implementations of APIs

#### moto

* moto is a Python library that allows you to mock AWS services in your unit tests. It works by providing a fake implementation of the AWS APIs, so that you can write tests that interact with these APIs without actually connecting to AWS. boto3 is a Python library that provides a low-level interface to AWS services, allowing you to interact with them programmatically. You can use boto3 with moto to write tests for your AWS code.

* So if we use the following code, including importing mock_s3 from moto:

```
import pytest
from moto import mock_s3
import boto3

def create_bucket(bucket_name):
    s3 = boto3.client('s3')
    s3.create_bucket(Bucket=bucket_name)
    print("Bucket name created.")

def test_create_bucket():
    with mock_s3():
        bucket_name = 'my-test-bucket'
        s3 = boto3.client('s3')
        s3.create_bucket(Bucket=bucket_name)
        
        buckets = s3.list_buckets()
        assert bucket_name in [bucket['Name'] for bucket in buckets['Buckets']]

```

* The test_create_bucket() function is using the mock_s3 decorator from the moto library to mock the S3 service. This means that instead of making real API calls to the S3 service, the mock_s3 decorator creates a mock S3 service that can be used for testing. 

* A mock_s3 service is a mock implementation of Amazon S3, which simulates the behavior of the real S3 service, allowing developers to test their code that interacts with S3 without actually making any requests to the real S3 service. mock_s3 creates a local in-memory S3 service that can be used to test S3-related functionality in isolation, making it ideal for unit testing.

* Within the test function, the create_bucket() function is called using boto3.client('s3') to create a bucket with the name 'my-test-bucket' and the s3.list_buckets() method is used to list all the available buckets in the mocked S3 service. The test then checks if the bucket with name 'my-test-bucket' is in the list of buckets returned by the mocked S3 service. Since the mock_s3 decorator has created a mock S3 service and the create_bucket() function is using this mocked service to create a bucket, the test passes without actually creating a real bucket in the AWS S3 service.

### MagicMock, patch

* We define MagicMock as a way to test a fake AWS service further.

#### MagicMock

* MagicMock class to create mock objects that simulate the behavior of boto3 and moto objects.
* These mock objects simulate the behavior of the objects they replace, allowing you to test your code in isolation without relying on external dependencies.


```
def test_upload_file_to_s3(mock_boto3_client):
    mock_s3_client = MagicMock()
    mock_boto3_client.return_value = mock_s3_client
```

* mock_s3_client = MagicMock() creates a new MagicMock object, which simulates the behavior of boto3 and moto objects, and assigns it to the variable mock_s3_client.
* So once you have mock_s3_client, then you can run functions on that, such as mock_boto3_client.return_value = mock_s3_client sets the return value of the faked boto3.client function to mock_s3_client when it is called within the scope of the decorated test function.
* The MagicMock object can then be used to set up behavior for the client and assert that certain functions were called with specific parameters.

```
...
    upload_file_to_s3('test.txt', 'mybucket', 'test.txt')

    mock_boto3_client.assert_called_once_with('s3')
    mock_s3_client.upload_file.assert_called_once_with('test.txt', 'mybucket', 'test.txt')
```

* So basically, we do an upload_file_to_s3() function above, which is our custom function.
* mock_boto3_client.assert_called_once_with('s3') is an assertion that checks whether the mock_boto3_client mock object has been called exactly once with the argument 's3'.
* In the context of the given code, these statements ensure that the upload_file_to_s3 function is correctly called and that the boto3.client function is called with the expected argument 's3'.
* mock_s3_client.upload_file.assert_called_once_with('test.txt', 'mybucket', 'test.txt') is an assertion that checks whether the upload_file method of the mock_s3_client mock object has been called exactly once with the specified arguments 'test.txt', 'mybucket', and 'test.txt'.
* In the context of the given code, this statement ensures that the upload_file method of the mocked S3 client is called with the expected arguments when the upload_file_to_s3 function is invoked.
* So basically, there are two tests, 1. To ensure that the s3 argument was called with upload_file_to_s3() and then 2. To ensure that the mock_s3_client.upload_file works in general.

#### Patch

* In the test code, we have decorated the test function with @patch('boto3.client'). This replaces the boto3.client function with a mock object. When the test function is called and it tries to call boto3.client, it will instead call the mock object created by patch.

* Then, mock_boto3_client.return_value = mock_s3_client sets the return value of the mock object to mock_s3_client. This means that when boto3.client is called within the test function, it will return mock_s3_client.

* This is useful because we can set up mock_s3_client to return mock objects for S3-related functions, allowing us to simulate S3 behavior without actually interacting with a real S3 service.