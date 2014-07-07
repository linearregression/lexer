package lexer_test

import (
  "github.com/Southern/lexer"
  "testing"
)

var l = lexer.New()

func TestParse(t *testing.T) {

  scan, err := l.Parse("Javascript", `/*

  Animal can be used as a base for different types of animals.

*/
function Animal(name) {
  this.name = name;
}

// Make our animal say hello!
Animal.prototype.sayHello = function() {
  return 'Hello from ' + this.name;
};

// Make our animal make some noise!
Animal.prototype.makeNoise = function() {
  return this.noise || '<chirp>';
};

/*

  Aw, look. It's a cute little dog.

*/
function Dog(name, breed) {
  this.name = name;
  this.breed = breed;
  this.noise = 'Woof!';
}

// Inherit Animal
Dog.prototype = new Animal();

// Our dog is smart. He can say hello AND his breed.
Dog.prototype.sayExtendedHello = function() {
  return this.sayHello() + ', ' + this.breed;
};

// Our dog can also bark. Not as impressive.
Dog.prototype.bark = function() {
  return this.noise;
};

// Expose our Animal and Dog to the outside world.
module.exports = {
  Animal: Animal,
  Dog: Dog,
};
`)

  if err != nil {
    t.Errorf("Unexpected error: %s", err)
    return
  }

  Status("Scan: %+v", scan)
}

func TestParseNoDataError(t *testing.T) {
  _, err := l.Parse()

  if err == nil {
    t.Errorf("Expected error.")
    return
  }

  Status("Got error: %+v", err)
}

func TestParseStringFirstError(t *testing.T) {
  _, err := l.Parse([]int{1, 2, 3, 4, 5}, "Test")

  if err == nil {
    t.Errorf("Expected error.")
    return
  }

  Status("Got error: %+v", err)
}

func TestParseScannerError(t *testing.T) {
  _, err := l.Parse([]int{1, 2, 3, 4, 5})

  if err == nil {
    t.Errorf("Expected error.")
    return
  }

  Status("Got error: %+v", err)
}

func TestReadFile(t *testing.T) {
  scan, err := l.ReadFile("testdata/animal.js")

  if err != nil {
    t.Errorf("Unexpected error: %s", err)
    return
  }

  Status("Scan: %+v", scan)
}

func TestReadFileWithNoLanguageMatch(t *testing.T) {
  scan, err := l.ReadFile("testdata/plain.txt")

  if err != nil {
    t.Errorf("Unexpected error: %s", err)
    return
  }

  Status("Scan: %+v", scan)
}

func TestReadFileInvalidFileError(t *testing.T) {
  _, err := l.ReadFile("testdata/idontexist")

  if err == nil {
    t.Errorf("Expected error.")
    return
  }

  Status("Got error: %+v", err)
}
