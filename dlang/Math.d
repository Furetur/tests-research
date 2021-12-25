import std.stdio;

class Sum
{
    int add(int x, int y) { return x + y; }

    unittest
    {
        Sum sum = new Sum;
        assert(sum.add(3,4) == 7);
        assert(sum.add(-2,0) == -2);
    }

    unittest
    {
        assert(1 == 2);
    }
}

void main() {
    writeln("Hello world");
}
