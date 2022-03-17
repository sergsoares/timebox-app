# tiny-timebox

Timebox app created to notifiy in linux using [Hey Notification](imgs/https://github.com/lukechampine/hey)

## Installation

```
go get -u github.com/sergsoares/tiny-timebox
```

## Screenshots

It create a tiny notification when timebox start and end.
![](imgs/2022-03-17-13-40-59.png)

## Integration with Ulauncher

Project can be integrated as a Ulauncher Shortcut.


![Configuration for Ulauncher](imgs/2022-03-17-13-41-36.png)


#### Snippet to be configured
```
#!/bin/bash

cd <TIME_BOX>
go run main.go -t $@
```

And be used as a normal shortcut.

![Command being used](imgs/2022-03-17-13-44-28.png)


## References

- https://github.com/gen2brain/beeep
- https://github.com/lukechampine/hey
- https://github.com/TheCreeper/go-notify