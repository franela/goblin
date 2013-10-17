Goblin
======

A [Mocha](http://visionmedia.github.io/mocha/) like BDD testing framework for Go

No extensive documentation nor complicated steps to get it running

Run tests as usual with `go test`

Colorful reports and beautiful syntax


Why Goblin?
-----------

Inspired by the flexibility and simplicity of Node BDD and frustrated by the
rigorousness of Go way of testing, we wanted to bring a new tool to 
write self-describing and comprehensive code.



What do I get with it?
----------------------

- Preserve the exact same syntax and behaviour as Node's Mocha
- Nest as many `Describe` and `It` blocks as you want
- Use `before`, `beforeEach`, `after` and `afterEach` for setup and teardown your tests
- No need to remember confusing parameters in `Describe` and `It` blocks
- Use a declarative and expressive language to write your tests
- Plug different assertion libraries (Gomega supported so far)
- Skip your tests the same way as you would do in Mocha
- Two line setup is all you need to get up running



How do I use it?
----------------

### Mocha's: 

```javascript
describe("Numbers", function() {
  it("Should add two numbers", function() {
    (1+1).should.equal(2);
  });
  it("Should match equal numbers", function() {
    (2).should.equal(4);
  });
  it("Should substract two numbers");
});
```


### Goblin's: 

```go
g := Goblin(t)
g.Describe("Numbers", func() {
    g.It("Should resolve the caller filename ", func() {
        g.Assert(1+1).Equal(2)
    })
    g.It("Should match equal numbers", func() {
        g.Assert(2).Equal(4)
    })
    g.It("Should substract two numbers")

})

```


Ouput will be something like:



Nice and easy, do you think?.


TODO:
-----

We do have a couple of [issues](https://github.com/franela/goblin/issues) pending we'd like to fix someday. Feel free to
contribute and send us some PR's (with tests please :smile:)
