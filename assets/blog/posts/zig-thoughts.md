*Dante Trisciuzzi 08-21-2023*



## The Appeal

I love working in C, nearly all of the embedded work I've done has been in C, and it's the first language I learned... nearly 20 years ago. Zig presents itself as a modern alternative to C, and even goes so far to say "faster then C" (Spoiler; it is). I have only been fiddling around with Zig for a couple of months, but I've found a few significant improvements over C; `comptime`, definitive type system, and the approach to memory allocation. Below are a few examples from my experience, they are far from exhaustive and I won't talk about error handling, C compatibility, builtins, or even ASM (maybe in a later post).

---

## Comptime

Zig's `comptime` enables compile-time code execution. Preprocessor macros are notoriously horrible, and Zig fixes this with a solution that gives us flexibility and clarity. This is a huge benefit and the primary reason I like Zig, it just makes sense.

Here's an example of how this could be leveraged to switch build modes in some application.

```zig
const Mode = enum {
    Development,
    Staging,
    Production
};

// Set this according to the desired build mode
comptime const CurrentMode = Mode.Development;

const Config = struct {
    const Endpoint = switch (CurrentMode) {
        Mode.Development => "http://localhost:8000/api/",
        Mode.Staging => "https://staging.myapp.com/api/",
        Mode.Production => "https://production.myapp.com/api/",
    };

    const MaxConnections = switch (CurrentMode) {
        Mode.Development => 10,
        Mode.Staging => 50,
        Mode.Production => 100,
    };
    
    // ... other config parameters
};

pub fn main() !void {
    // We treat the config as an object, and we can use it as such
    // it fits into the language
    std.debug.print("API Endpoint: {}\n", .{Config.Endpoint});
    std.debug.print("Max Connections: {}\n", .{Config.MaxConnections});
}
```
and in C...
```c
#include <stdio.h>

// Set the mode by uncommenting one of these
#define Development
// #define Staging
// #define Production

#ifdef Development
    #define ENDPOINT "http://localhost:8000/api/"
    #define MAX_CONNECTIONS 10
#elif defined(Staging)
    #define ENDPOINT "https://staging.myapp.com/api/"
    #define MAX_CONNECTIONS 50
#elif defined(Production)
    #define ENDPOINT "https://production.myapp.com/api/"
    #define MAX_CONNECTIONS 100
#endif

int main() {

    // Now we get to use ENDPOINT and MAX_CONNECTIONS throughout the entire application
    // We do not have a configuration object...

    printf("API Endpoint: %s\n", ENDPOINT); // What type is this?
    printf("Max Connections: %d\n", MAX_CONNECTIONS); // what type is this?
    return 0;
}
```
`comptime` empowers developers to handle situations like these easily. Why would I want to write functional parts of my code in any language in the language I'm writing? `comptime` is Zig, macros are not C.

---
## Type System

Zig ensures an unambiguous type system, no type promotions. Additionally, Zig has type-specific compile time checks, like bounds-checking for arrays. This is better for everyone.

```zig
pub fn main() !void {
    const array = [_]i32{1, 2, 3, 4, 5};
    std.debug.print("{}\n", .{array[10]}); // comptime error: out-of-bounds access
}
```
---
## Memory Allocation

Zig provides a controlled approach to memory management while retaining manual memory allocation. Instead of relying solely on `malloc` and `free`, Zig introduces an allocator object, making memory allocation more consistent. 

```zig
const std = @import("std");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();

    const allocator = arena.allocator();

    const v = try allocator.create(i32);
    _ = v;
}
```

This approach offers the flexibility to switch between different allocation methods, which is cool.

---
## Conclusion

So... where does Zig fit into my toolbox right now? It's fun and provides a lot of good features and niceties (like defer, used in the memory allocation example). I am starting to use it to suppliment my work in C, since it integrates seamlessly. I'd like to find some projects to write in pure Zig, but I'm not working on anything that doesnt already have a fairly established code base currently.
