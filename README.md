# Go Context Basics
This Go application demonstrates the usage of context in managing asynchronous operations and fetching data from a third-party source.

The main functionality of the application is showcased in the fetchUserData function, where user data is retrieved asynchronously from a third-party source using goroutines and channels. Contexts are employed to manage timeouts and prevent context leakage.