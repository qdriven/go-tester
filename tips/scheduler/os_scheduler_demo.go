package scheduler

/**
1. Waiting: IO
2. Runnable
3. Execution

If you have a program that is focused on IO-Bound work,
then context switches are going to be an advantage.
Once a Thread moves into a Waiting state,
another Thread in a Runnable state is there to take its place.
This allows the core to always be doing work.
This is one of the most important aspects of scheduling.
Don’t allow a core to go idle if there is work (Threads in a Runnable state) to be done.
*/


