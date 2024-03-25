*Dante Trisciuzzi 07-24-2023*



## It's Trendy

I'm not much of a blogger despite the cool catchy title, just a fourth-year computer science student. I am interested in AI and have been experimenting with reinforcement training myself; so I see the potential AI holds for our field. However, I've observed a trend; more and more students are using AI tools like ChatGPT not **as supplements, but as substitutes** to their learning. I'll provide a few small examples in this post, these are just my experiences and opinions. *I might rant a little at times, and the structure is pretty rough.*

***TLDR: This misuse of AI in academic work (employing it as a crutch rather than a tool) undermines your education.***

---

## ChatGPT in the discussion board

It's pretty sad to see how some students put absolutely no effort into their work, where is the pride? Every class I've taken over the past few years has used discussion boards to "engage" us students, and recently I've seen many students turning to ChatGPT as their *magical solution*. They copy the entire discussion prompt in, and paste whatever it spits out into their discussion board. Oftentimes, their responses don't even really address the prompt. Here is an example from a recent class;  

Regardless of the wording, the idea behind the prompt is: **"Pick a household item that uses embedded programming, do some research and explain how it might use hardware interrupts"**.  

The reponse they post is something like the following:
> Firstly, it's important to understand what an interrupt is. In the context of embedded programming, an interrupt is a signal that alters the normal flow of a microcontroller's operation. It is a mechanism that allows external devices or internal processes to request immediate attention from the microcontroller, diverting its current task to handle the urgent request. My item that utilizes embedded programming to control settings and create an optimal environment for the household.

It's worth noting that this is a 4th year CS course on system architechtures, where we are doing embedded C programming. We have already been using interrupts, so we (should) know what they are, and this does not even come close to addressing the prompt.

I can't imagine that a student who posts this understands what they're doing when they actually sit down to start writing code, or maybe this type of work is just too far below them. But, they post this in the board for us (other students) to respond to and **engage** with. How are any of us supposed to find value here, what can we possibly discuss that benefits anyone?

I'm not sure what you're thinking (nothing?) when you post like this, but you affect everyone. I know that these posts still get points somehow (seen the same people class after class for a few years now), so you're just spreading the laziness to others.

---

## AI Code Teaches You Almost Nothing

AI-generated code, while sometimes technically correct, often lacks proper structuring or context, causing safety, maintenance, and scalability challenges.  
  

#### A tiny example:
In one of my courses, we have a Q&A board, just for helping peers. A 4th year student (who has little experience with low-level programming) posted some code, openly admitting that ChatGPT wrote some of it, and was having an issue where his program "just stopped". There was nothing inherently wrong with what he was doing (making a basic adlib-type game) but he failed to post the 'helpers.h' header file along with his main code. Long story short, he was using this ChatGPT generated function that was tucked away in the header file:
```c
void copyString(const char* source, char* destination) {
    strcpy(destination, source);
}
```
The problem here is not that GPT generated some junky code, but that this student has almost no experience debugging after years of CS courses, and did not even know how to start diagnosing the buffer overflow they were experiencing. GPT code gets many projects up-and-running in a matter of minutes by producing code that compiles, but leads to people just pasting away with no intention of understanding what they're doing.

Additionally, due to the lack of context and creativity, AI-generated code may fail to utilize appropriate design patterns, or even recognize the complexities of architecture. It often overlooks fundamental software engineering principles, resulting in code that is hard to extend, test, maintain, or adapt to changing requirements. These issues can impact safety, scalability, and performance.

---

## The Educational System's Failure to Educate

At my school, many instructors rely on automated grading systems. These systems have varying degrees of independence, but they frequently fail to distinguish well-structured code from technically correct yet poorly written code. This is pretty expected; but what is not expected is that often, instructors do not look at the work themselves. 

I found this out because I lost points on an assignement and requested to know why. I was given the explaination that there was not enough comments in [my C++ code](https://github.com/Triscuit2311/SNHU-Portfolios/blob/main/CS300/CS300-Proj2/CS300-Proj2/BinarySearchTree.cpp) . There were tons of comments in my code, all following the format specified in our technical requirements document. The tool missed these for one reason or another, and after the instructor reviewed my code (for the first time, after I'd already been graded) I got full points for the assignment. So this begs the question, what is the incentive to write clean or generally nice code, if it's not even being looked at? Well, I'm in this for myself, but this experience opened my eyes a bit. 

This is an oversight that disregards problem-solving abilities, conceptual understanding, and the creativity to come up with a clever solution; all vital attributes in my opinion. This is a flaw in our educational evaluation system, and it devalues diligent effort and waters down academic achievements related to coding. The education we're pursuing isn't just about learning to code; it's about learning to think, solve problems, and to be creative. 

When AI can produce code that barely meets the minimum requirements and achieves the same recognition as thoughtfully crafted code, it's a disservice to every student who is genuinely investing time and effort to learn. This failure points to a systemic issue that needs addressing.
   

---

## Take Pride In Your Work

It's alarming to me that these students will join the workforce as my peers. Let's try remember that education isn't just about obtaining a degree; it's about learning and preparing for real-world challenges. I encourage all students to use AI responsibly, use it to **aid your education** rather than **diminish it**.
