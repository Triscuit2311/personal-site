*Dante Trisciuzzi 01-19-2024*


## Background

I have been using C and C++ the majority of my life, when I approach programming problems, I think in C. Regardless of the language or project I'm working on, I always consider the systems level or underlying operations, like memory management, allocations, references, etc. I am only mentioning this because I think it is important context for my opinions.

---

## The basics

Let's start with the obvious, Concurrency. I am no stranger to threads, or jthreads as is trendy now in C++, and I don't really find them to be difficult to use. Just like anything else in these languages, its up to the developer to make sure the loose ends are taken care (including errors). Most of the functionality I need can be had with mutexes and atomics. Go's approach is completely different, and it is pretty innovative, or at least it feels that way to me. It challenges me to simplify my solutions, and I like that. Channels in Go are a great concept and probably one of my most used features thus far, maybe I'm just looking for reasons to use them. Nonetheless, concurrency in Go is excellent, but not what makes it stand out to me.

Second, memory management and low-level control. Obviously, since Go has a garbage collector, manual memory management is not required. However, Go does support unsafe C operations like allocation, direct casting (functionally like reinterpret casting), pointer arithmetic, and other features I thought I would miss from C. But truthfully, I do not miss them at all, because once again; I am finding simpler solutions to problems. It is nice to know they are there though.

---

## Why I kind of love it

Go is strict where it needs to be, and forgiving elsewhere. No type promotion, modern type system, and common-sense casting make working with basic types as simple as it should have been in C. Inheritance is nowhere to be found in Go, instead there is interfaces, which are an unambigous solution to provide polymorphism in an extremely simple manner. This means that multiple inheritance and the mess that creates is gone; in Go, if a type implements the interface's specified methods, it satisfies the requirements to be accessed though that interface. In the recent Go language update, we have templates in the stable release now. It seems to be pretty devisive in the community, but I personally don't mind it. It has actually added to the simplicity of the language, even though technically making it more complex; by this I mean that you can be more expressive and do more with less now. With that said, I have really only used them for fairly simple use cases (like my local version control project).


It feels familiar. Writing Go feels like writing C, there are plenty of differences in syntax, but overall it keeps a lot of that simplistic feel. There is however, lots of welcome changes like the simplified conditional system, the `for` loop syntax, variable initialization and assignment, and more. I want to compare this to languages like C#, but C# actually feels much more complex to write.

Errors, while I have seen error handling in Go catching flak, I love it. Errors as values just makes sense, and forcing the responsibility to handle them on the developer makes safety more inherent. 

Lastly, the ecosystem. Go comes ready to rip, I hate saying this, but "It just works". The standard library is full of powerful features and the module system makes it easy to get what you need without a hassle. This also applies to 3rd party libraries, most of the time Go does the work for you in a transparent manner. Compare this to C#, which aims to abstract so much away from the developer. I find this this to be one of the biggest selling points for me, as it really cuts the idea-to-code time down.

---

## Final thoughts

Go has found a place in my toolbox almost instantly. While I (perosnally) can't squeeze every bit of performance out of my code like I can in C, not everything needs to shave nanoseconds. Need to build a CLI tool, web application (like this website), or a GUI desktop application that runs cross-platform? Go has me covered. Unlike Zig (which I love), I was able to get up on my feet and productive in pretty much no time.
