# `/pkg`

Library code that's ok to use by external applications. Other projects will import these libraries expecting them to work, so think twice before you put something here Note that the `internal` directory is a better way to ensure your private packages are not importable because it's enforced by Go.