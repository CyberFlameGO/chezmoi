# `promptIntOnce` *map* *key* *prompt* [*default*]

`promptIntOnce` returns *map*.*key* if it exists and is an integer value,
otherwise it prompts the user for a integer value with *prompt* and an optional
*default* using `promptInt`.

!!! example

    ```
    {{ $monitors := promptIntOnce . "monitors" "How many monitors does this machine have" }}
    ```
