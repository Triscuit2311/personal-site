*Dante Trisciuzzi 09-05-2023*
 
 

### The Story

A few years ago I had a friend who lost significant mobility in a pretty terrible accident. I'll spare the details, but they were left with essentially a few fingers that function mostly correctly. As a gamer and computer enthusiast; this was pretty tough. We looked into a few different solutions and found some things that worked well, but the flexibility was pretty limited. Why not use the commercial products? Well, we tried; unfortunately some game developers saw certain software and HID drivers as a threat to their game's security, some even rule macro software as a violation. Aside from this, what if we want multiple devices in use? Clashing software issues, closed-source kernel drivers, and other complications pushed me to make something simpler. I started working on this project to provide a universal, hardware-based human interface device (HID) emulator that anyone could adapt to their use.

[SLACCr](https://github.com/Triscuit2311/SLACCr) was born. SLACCr (pronounced like 'slacker') isn't a cool acronym, it was just a little inside joke that kinda stuck around. I designed it to be flexible and easy adaptable.

---

### Why Arduino, Why RP2040?

Simple, it's fast, flexible and easy to setup. I took a user-first approach with SLACCr, all it takes is an RP2040 Board (Raspberry Pi Pico for only a few dollars), and it uses a community-driven Arduino framework. This means drag-and-drop firmware installation; anyone can get started without a hitch. Originally I wrote the firmware on bare metal, but this made it difficult for people who wanted to try to get setup.

The RP2040's dual-core architecture was an obvious choice, it's fast. This architecture guarantees that the emulator runs seamlessly, one thread handles the packet I/O, the other is dedicated to sending the inputs to the host.

---

### How?

I worked hard to develop a comprehensive API, meant to be user-friendly, I currently have native APIs for C++, C#, and Python. Getting up and running as a developer is super simple, you can use the pre-built dynamic library and reference the API, or you can link against the static library. Alternatively, build it from source or integrate it into your own project.

Here is the example gamepad test from the repo, as you can see it's super easy to use even from the most verbose implementation (C++):


```cpp
#include <iostream>
#include <SLACCR_API.h>
#include <thread>

void GamePadTest();

int main()
{
    if(Slaccr::Init(false))
    {
        std::cout << "CONNECTED\n";

        while (true) {
            GamePadTest();
        }
    }
}

void GamePadTest()
{
    // Axises
    {
        auto moveAllAxises = [](auto i)
        {
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::X, i);
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::Y, i);
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::Z, i);
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::Zrotate, i);
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::SliderLeft, i);
            Slaccr::GamepadMoveAxis(Peripherals::JoystickAxis::SliderRight, i);
        };

        int axisMaxV = 127;
        for (auto n = 1; n < 5; n++) {
            auto i = -axisMaxV;
            for (i; i <= axisMaxV; i++) {
                moveAllAxises(i);
            }
            axisMaxV /= n;
            for (i; i >= -axisMaxV; i--) {
                moveAllAxises(i);
            }
        }
        Slaccr::GamepadResetAxises();
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
    }

    // Hat
    {
        for (auto i = 0; i <= 361; i += 15) {
            Slaccr::GamepadSetHat(i);
            std::this_thread::sleep_for(std::chrono::milliseconds(20));
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(250));
        Slaccr::GamepadResetHat();
        std::this_thread::sleep_for(std::chrono::milliseconds(250));
    }

    // Buttons
    {
        for (auto i = 1; i <= 32; i++) {
            Slaccr::GamepadButtonPress(i, 100);
            std::this_thread::sleep_for(std::chrono::milliseconds(150));
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(500));

        for (auto i = 1; i <= 32; i++) {
            Slaccr::GamepadButtonState(i, true);
            std::this_thread::sleep_for(std::chrono::milliseconds(50));
        }
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
        Slaccr::GamepadResetButtons();
    }
}
```

---

### On the Horizon

As with anything, there is always room for improvement. Immediate plans for SLACCr include adding more practical examples, refining the APIs, and expanding documentation.

Future roadmap milestones include obtaining unique USB vendor and product IDs and transitioning the firmware back to a bare-metal approach, but with better documentation and a dedicated firmware update tool. Wireless maybe?

---



SLACCr is my statment on the importance of inclusive technology.
